// Code generated by protoc-gen-router. DO NOT EDIT.

package lockunlock

import (
	context "context"
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
	io "io"
)

// ApiRouter is a traits.LockUnlockApiServer that allows routing named requests to specific traits.LockUnlockApiClient
type ApiRouter struct {
	traits.UnimplementedLockUnlockApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.LockUnlockApiServer = (*ApiRouter)(nil)

func NewApiRouter(opts ...router.Option) *ApiRouter {
	return &ApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithLockUnlockApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithLockUnlockApiClientFactory(f func(name string) (traits.LockUnlockApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *ApiRouter) Register(server *grpc.Server) {
	traits.RegisterLockUnlockApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.LockUnlockApiClient.
func (r *ApiRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.LockUnlockApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *ApiRouter) HoldsType(client any) bool {
	_, ok := client.(traits.LockUnlockApiClient)
	return ok
}

func (r *ApiRouter) AddLockUnlockApiClient(name string, client traits.LockUnlockApiClient) traits.LockUnlockApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.LockUnlockApiClient)
}

func (r *ApiRouter) RemoveLockUnlockApiClient(name string) traits.LockUnlockApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.LockUnlockApiClient)
}

func (r *ApiRouter) GetLockUnlockApiClient(name string) (traits.LockUnlockApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.LockUnlockApiClient), nil
}

func (r *ApiRouter) GetLockUnlock(ctx context.Context, request *traits.GetLockUnlockRequest) (*traits.LockUnlock, error) {
	child, err := r.GetLockUnlockApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetLockUnlock(ctx, request)
}

func (r *ApiRouter) UpdateLockUnlock(ctx context.Context, request *traits.UpdateLockUnlockRequest) (*traits.LockUnlock, error) {
	child, err := r.GetLockUnlockApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.UpdateLockUnlock(ctx, request)
}

func (r *ApiRouter) PullLockUnlock(request *traits.PullLockUnlockRequest, server traits.LockUnlockApi_PullLockUnlockServer) error {
	child, err := r.GetLockUnlockApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullLockUnlock(reqCtx, request)
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

		var msg *traits.PullLockUnlockResponse
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
