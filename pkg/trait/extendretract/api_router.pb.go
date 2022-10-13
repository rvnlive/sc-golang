// Code generated by protoc-gen-router. DO NOT EDIT.

package extendretract

import (
	context "context"
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
	io "io"
)

// ApiRouter is a traits.ExtendRetractApiServer that allows routing named requests to specific traits.ExtendRetractApiClient
type ApiRouter struct {
	traits.UnimplementedExtendRetractApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.ExtendRetractApiServer = (*ApiRouter)(nil)

func NewApiRouter(opts ...router.Option) *ApiRouter {
	return &ApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithExtendRetractApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithExtendRetractApiClientFactory(f func(name string) (traits.ExtendRetractApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *ApiRouter) Register(server *grpc.Server) {
	traits.RegisterExtendRetractApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.ExtendRetractApiClient.
func (r *ApiRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.ExtendRetractApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *ApiRouter) HoldsType(client any) bool {
	_, ok := client.(traits.ExtendRetractApiClient)
	return ok
}

func (r *ApiRouter) AddExtendRetractApiClient(name string, client traits.ExtendRetractApiClient) traits.ExtendRetractApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.ExtendRetractApiClient)
}

func (r *ApiRouter) RemoveExtendRetractApiClient(name string) traits.ExtendRetractApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.ExtendRetractApiClient)
}

func (r *ApiRouter) GetExtendRetractApiClient(name string) (traits.ExtendRetractApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.ExtendRetractApiClient), nil
}

func (r *ApiRouter) GetExtension(ctx context.Context, request *traits.GetExtensionRequest) (*traits.Extension, error) {
	child, err := r.GetExtendRetractApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetExtension(ctx, request)
}

func (r *ApiRouter) UpdateExtension(ctx context.Context, request *traits.UpdateExtensionRequest) (*traits.Extension, error) {
	child, err := r.GetExtendRetractApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.UpdateExtension(ctx, request)
}

func (r *ApiRouter) Stop(ctx context.Context, request *traits.ExtendRetractStopRequest) (*traits.Extension, error) {
	child, err := r.GetExtendRetractApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.Stop(ctx, request)
}

func (r *ApiRouter) CreateExtensionPreset(ctx context.Context, request *traits.CreateExtensionPresetRequest) (*traits.ExtensionPreset, error) {
	child, err := r.GetExtendRetractApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.CreateExtensionPreset(ctx, request)
}

func (r *ApiRouter) PullExtensions(request *traits.PullExtensionsRequest, server traits.ExtendRetractApi_PullExtensionsServer) error {
	child, err := r.GetExtendRetractApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullExtensions(reqCtx, request)
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

		var msg *traits.PullExtensionsResponse
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
