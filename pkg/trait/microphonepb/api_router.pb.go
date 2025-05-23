// Code generated by protoc-gen-router. DO NOT EDIT.

package microphonepb

import (
	context "context"
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	types "github.com/smart-core-os/sc-api/go/types"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
	io "io"
)

// ApiRouter is a traits.MicrophoneApiServer that allows routing named requests to specific traits.MicrophoneApiClient
type ApiRouter struct {
	traits.UnimplementedMicrophoneApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.MicrophoneApiServer = (*ApiRouter)(nil)

func NewApiRouter(opts ...router.Option) *ApiRouter {
	return &ApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithMicrophoneApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithMicrophoneApiClientFactory(f func(name string) (traits.MicrophoneApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *ApiRouter) Register(server grpc.ServiceRegistrar) {
	traits.RegisterMicrophoneApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.MicrophoneApiClient.
func (r *ApiRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.MicrophoneApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *ApiRouter) HoldsType(client any) bool {
	_, ok := client.(traits.MicrophoneApiClient)
	return ok
}

func (r *ApiRouter) AddMicrophoneApiClient(name string, client traits.MicrophoneApiClient) traits.MicrophoneApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.MicrophoneApiClient)
}

func (r *ApiRouter) RemoveMicrophoneApiClient(name string) traits.MicrophoneApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.MicrophoneApiClient)
}

func (r *ApiRouter) GetMicrophoneApiClient(name string) (traits.MicrophoneApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.MicrophoneApiClient), nil
}

func (r *ApiRouter) GetGain(ctx context.Context, request *traits.GetMicrophoneGainRequest) (*types.AudioLevel, error) {
	child, err := r.GetMicrophoneApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetGain(ctx, request)
}

func (r *ApiRouter) UpdateGain(ctx context.Context, request *traits.UpdateMicrophoneGainRequest) (*types.AudioLevel, error) {
	child, err := r.GetMicrophoneApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.UpdateGain(ctx, request)
}

func (r *ApiRouter) PullGain(request *traits.PullMicrophoneGainRequest, server traits.MicrophoneApi_PullGainServer) error {
	child, err := r.GetMicrophoneApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullGain(reqCtx, request)
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

		var msg *traits.PullMicrophoneGainResponse
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
