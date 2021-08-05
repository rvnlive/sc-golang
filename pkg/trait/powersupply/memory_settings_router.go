package powersupply

import (
	"context"
	"io"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MemorySettingsRouter is a MemorySettingsApiServer that allows routing named requests to specific MemorySettingsApiClients
type MemorySettingsRouter struct {
	UnimplementedMemorySettingsApiServer

	mu       sync.Mutex
	registry map[string]MemorySettingsApiClient
	// Factory can be used to dynamically create api clients if requests come in for devices we haven't seen.
	Factory func(string) (MemorySettingsApiClient, error)
}

var _ MemorySettingsApiServer = (*MemorySettingsRouter)(nil) // compiler check we implement the interface

func NewMemorySettingsRouter() *MemorySettingsRouter {
	return &MemorySettingsRouter{
		registry: make(map[string]MemorySettingsApiClient),
	}
}

func (r *MemorySettingsRouter) Register(server *grpc.Server) {
	RegisterMemorySettingsApiServer(server, r)
}

func (r *MemorySettingsRouter) Add(name string, client MemorySettingsApiClient) MemorySettingsApiClient {
	r.mu.Lock()
	defer r.mu.Unlock()
	old := r.registry[name]
	r.registry[name] = client
	return old
}

func (r *MemorySettingsRouter) Remove(name string) MemorySettingsApiClient {
	r.mu.Lock()
	defer r.mu.Unlock()
	old := r.registry[name]
	delete(r.registry, name)
	return old
}

func (r *MemorySettingsRouter) Has(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.registry[name]
	return exists
}

func (r *MemorySettingsRouter) Get(name string) (MemorySettingsApiClient, error) {
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

func (r *MemorySettingsRouter) GetSettings(ctx context.Context, req *GetMemorySettingsReq) (*MemorySettings, error) {
	child, err := r.Get(req.Name)
	if err != nil {
		return nil, err
	}
	return child.GetSettings(ctx, req)
}

func (r *MemorySettingsRouter) UpdateSettings(ctx context.Context, req *UpdateMemorySettingsReq) (*MemorySettings, error) {
	child, err := r.Get(req.Name)
	if err != nil {
		return nil, err
	}
	return child.UpdateSettings(ctx, req)
}

func (r *MemorySettingsRouter) PullSettings(request *PullMemorySettingsReq, server MemorySettingsApi_PullSettingsServer) error {
	child, err := r.Get(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullSettings(reqCtx, request)
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

		var msg *PullMemorySettingsRes
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
