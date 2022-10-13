// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package inputselect

import (
	context "context"
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapApi	adapts a traits.InputSelectApiServer	and presents it as a traits.InputSelectApiClient
func WrapApi(server traits.InputSelectApiServer) traits.InputSelectApiClient {
	return &apiWrapper{server}
}

type apiWrapper struct {
	server traits.InputSelectApiServer
}

// compile time check that we implement the interface we need
var _ traits.InputSelectApiClient = (*apiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *apiWrapper) UnwrapServer() traits.InputSelectApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *apiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *apiWrapper) UpdateInput(ctx context.Context, req *traits.UpdateInputRequest, _ ...grpc.CallOption) (*traits.Input, error) {
	return w.server.UpdateInput(ctx, req)
}

func (w *apiWrapper) GetInput(ctx context.Context, req *traits.GetInputRequest, _ ...grpc.CallOption) (*traits.Input, error) {
	return w.server.GetInput(ctx, req)
}

func (w *apiWrapper) PullInput(ctx context.Context, in *traits.PullInputRequest, opts ...grpc.CallOption) (traits.InputSelectApi_PullInputClient, error) {
	stream := wrap.NewClientServerStream(ctx)
	server := &pullInputApiServerWrapper{stream.Server()}
	client := &pullInputApiClientWrapper{stream.Client()}
	go func() {
		err := w.server.PullInput(in, server)
		stream.Close(err)
	}()
	return client, nil
}

type pullInputApiClientWrapper struct {
	grpc.ClientStream
}

func (w *pullInputApiClientWrapper) Recv() (*traits.PullInputResponse, error) {
	m := new(traits.PullInputResponse)
	if err := w.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type pullInputApiServerWrapper struct {
	grpc.ServerStream
}

func (s *pullInputApiServerWrapper) Send(response *traits.PullInputResponse) error {
	return s.ServerStream.SendMsg(response)
}
