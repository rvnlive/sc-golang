// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package electric

import (
	context "context"
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// Wrap Api	adapts a traits.ElectricApiServer	and presents it as a traits.ElectricApiClient
func WrapApi(server traits.ElectricApiServer) traits.ElectricApiClient {
	return &apiWrapper{server}
}

type apiWrapper struct {
	server traits.ElectricApiServer
}

// compile time check that we implement the interface we need
var _ traits.ElectricApiClient = (*apiWrapper)(nil)

func (w *apiWrapper) GetDemand(ctx context.Context, req *traits.GetDemandRequest, _ ...grpc.CallOption) (*traits.ElectricDemand, error) {
	return w.server.GetDemand(ctx, req)
}

func (w *apiWrapper) PullDemand(ctx context.Context, in *traits.PullDemandRequest, opts ...grpc.CallOption) (traits.ElectricApi_PullDemandClient, error) {
	stream := wrap.NewClientServerStream(ctx)
	server := &pullDemandApiServerWrapper{stream.Server()}
	client := &pullDemandApiClientWrapper{stream.Client()}
	go func() {
		err := w.server.PullDemand(in, server)
		stream.Close(err)
	}()
	return client, nil
}

type pullDemandApiClientWrapper struct {
	grpc.ClientStream
}

func (w *pullDemandApiClientWrapper) Recv() (*traits.PullDemandResponse, error) {
	m := new(traits.PullDemandResponse)
	if err := w.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type pullDemandApiServerWrapper struct {
	grpc.ServerStream
}

func (s *pullDemandApiServerWrapper) Send(response *traits.PullDemandResponse) error {
	return s.ServerStream.SendMsg(response)
}

func (w *apiWrapper) GetActiveMode(ctx context.Context, req *traits.GetActiveModeRequest, _ ...grpc.CallOption) (*traits.ElectricMode, error) {
	return w.server.GetActiveMode(ctx, req)
}

func (w *apiWrapper) UpdateActiveMode(ctx context.Context, req *traits.UpdateActiveModeRequest, _ ...grpc.CallOption) (*traits.ElectricMode, error) {
	return w.server.UpdateActiveMode(ctx, req)
}

func (w *apiWrapper) ClearActiveMode(ctx context.Context, req *traits.ClearActiveModeRequest, _ ...grpc.CallOption) (*traits.ElectricMode, error) {
	return w.server.ClearActiveMode(ctx, req)
}

func (w *apiWrapper) PullActiveMode(ctx context.Context, in *traits.PullActiveModeRequest, opts ...grpc.CallOption) (traits.ElectricApi_PullActiveModeClient, error) {
	stream := wrap.NewClientServerStream(ctx)
	server := &pullActiveModeApiServerWrapper{stream.Server()}
	client := &pullActiveModeApiClientWrapper{stream.Client()}
	go func() {
		err := w.server.PullActiveMode(in, server)
		stream.Close(err)
	}()
	return client, nil
}

type pullActiveModeApiClientWrapper struct {
	grpc.ClientStream
}

func (w *pullActiveModeApiClientWrapper) Recv() (*traits.PullActiveModeResponse, error) {
	m := new(traits.PullActiveModeResponse)
	if err := w.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type pullActiveModeApiServerWrapper struct {
	grpc.ServerStream
}

func (s *pullActiveModeApiServerWrapper) Send(response *traits.PullActiveModeResponse) error {
	return s.ServerStream.SendMsg(response)
}

func (w *apiWrapper) ListModes(ctx context.Context, req *traits.ListModesRequest, _ ...grpc.CallOption) (*traits.ListModesResponse, error) {
	return w.server.ListModes(ctx, req)
}

func (w *apiWrapper) PullModes(ctx context.Context, in *traits.PullModesRequest, opts ...grpc.CallOption) (traits.ElectricApi_PullModesClient, error) {
	stream := wrap.NewClientServerStream(ctx)
	server := &pullModesApiServerWrapper{stream.Server()}
	client := &pullModesApiClientWrapper{stream.Client()}
	go func() {
		err := w.server.PullModes(in, server)
		stream.Close(err)
	}()
	return client, nil
}

type pullModesApiClientWrapper struct {
	grpc.ClientStream
}

func (w *pullModesApiClientWrapper) Recv() (*traits.PullModesResponse, error) {
	m := new(traits.PullModesResponse)
	if err := w.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type pullModesApiServerWrapper struct {
	grpc.ServerStream
}

func (s *pullModesApiServerWrapper) Send(response *traits.PullModesResponse) error {
	return s.ServerStream.SendMsg(response)
}
