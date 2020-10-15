package router

import (
	"context"
	"io"
	"sync"

	"git.vanti.co.uk/smartcore/sc-api/go/device/traits"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CountApiRouter is a CountApiServer that allows routing named requests to specific CountApiClients
type CountApiRouter struct {
	traits.UnimplementedCountApiServer

	mu       sync.Mutex
	registry map[string]traits.CountApiClient
	// Factory can be used to dynamically create api clients if requests come in for devices we haven't seen.
	Factory func(string) (traits.CountApiClient, error)
}

// compile time check that we implement the interface we need
var _ traits.CountApiServer = &CountApiRouter{}

func NewCountApiRouter() *CountApiRouter {
	return &CountApiRouter{
		registry: make(map[string]traits.CountApiClient),
	}
}

func (r *CountApiRouter) Register(server *grpc.Server) {
	traits.RegisterCountApiServer(server, r)
}

func (r *CountApiRouter) Add(name string, client traits.CountApiClient) traits.CountApiClient {
	r.mu.Lock()
	defer r.mu.Unlock()
	old := r.registry[name]
	r.registry[name] = client
	return old
}

func (r *CountApiRouter) Remove(name string) traits.CountApiClient {
	r.mu.Lock()
	defer r.mu.Unlock()
	old := r.registry[name]
	delete(r.registry, name)
	return old
}

func (r *CountApiRouter) Has(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.registry[name]
	return exists
}

func (r *CountApiRouter) Get(name string) (traits.CountApiClient, error) {
	r.mu.Lock()
	child, exists := r.registry[name]
	defer r.mu.Unlock()
	if !exists {
		if r.Factory != nil {
			child, err := r.Factory(name)
			if err != nil {
				return nil, err
			}
			r.registry[name] = child
			return child, nil
		}
		return nil, status.Error(codes.NotFound, name)
	}
	return child, nil
}

func (r *CountApiRouter) GetCount(ctx context.Context, request *traits.GetCountRequest) (*traits.Count, error) {
	child, err := r.Get(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetCount(ctx, request)
}

func (r *CountApiRouter) ResetCount(ctx context.Context, request *traits.ResetCountRequest) (*traits.Count, error) {
	child, err := r.Get(request.Name)
	if err != nil {
		return nil, err
	}

	return child.ResetCount(ctx, request)
}

func (r *CountApiRouter) UpdateCount(ctx context.Context, request *traits.UpdateCountRequest) (*traits.Count, error) {
	child, err := r.Get(request.Name)
	if err != nil {
		return nil, err
	}

	return child.UpdateCount(ctx, request)
}

func (r *CountApiRouter) PullCounts(request *traits.PullCountsRequest, server traits.CountApi_PullCountsServer) error {
	child, err := r.Get(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullCounts(reqCtx, request)
	if err != nil {
		return err
	}

	// send the stream header
	header, err := stream.Header()
	if err != nil {
		return err
	}
	if err = server.SendHeader(header); err != nil {
		return err
	}

	// send all the messages
	// false means the error is from the child, true means the error is from the caller
	var callerError bool
	for {
		// Impl note: we could improve throughput here by issuing the Recv and Send in different goroutines, but we're doing
		// it synchronously until we have a need to change the behaviour

		var msg *traits.PullCountsResponse
		msg, err = stream.Recv()
		if err != nil {
			break
		}

		err = server.Send(msg)
		if err != nil {
			callerError = true
			break
		}
	}

	// err is guaranteed to be non-nil as it's the only way to exit the loop
	if callerError {
		// cancel the request
		reqDone()
		return err
	} else {
		if trailer := stream.Trailer(); trailer != nil {
			server.SetTrailer(trailer)
		}
		if err == io.EOF {
			return nil
		}
		return err
	}
}