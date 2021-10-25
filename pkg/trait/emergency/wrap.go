package emergency

import (
	"context"

	"github.com/smart-core-os/sc-api/go/traits"
	"github.com/smart-core-os/sc-golang/pkg/wrap"
	"google.golang.org/grpc"
)

// Wrap adapts a traits.EmergencyApiServer and presents it as a traits.EmergencyApiClient
func Wrap(server traits.EmergencyApiServer) traits.EmergencyApiClient {
	return &wrapper{server}
}

type wrapper struct {
	server traits.EmergencyApiServer
}

// compile time check that we implement the interface we need
var _ traits.EmergencyApiClient = (*wrapper)(nil)

func (b *wrapper) GetEmergency(ctx context.Context, in *traits.GetEmergencyRequest, opts ...grpc.CallOption) (*traits.Emergency, error) {
	return b.server.GetEmergency(ctx, in)
}

func (b *wrapper) UpdateEmergency(ctx context.Context, in *traits.UpdateEmergencyRequest, opts ...grpc.CallOption) (*traits.Emergency, error) {
	return b.server.UpdateEmergency(ctx, in)
}

func (b *wrapper) PullEmergency(ctx context.Context, in *traits.PullEmergencyRequest, opts ...grpc.CallOption) (traits.EmergencyApi_PullEmergencyClient, error) {
	stream := wrap.NewClientServerStream(ctx)
	server := &pullEmergencyServerWrapper{stream.Server()}
	client := &pullEmergencyClientWrapper{stream.Client()}
	go func() {
		err := b.server.PullEmergency(in, server)
		stream.Close(err)
	}()
	return client, nil
}

type pullEmergencyClientWrapper struct {
	grpc.ClientStream
}

func (c *pullEmergencyClientWrapper) Recv() (*traits.PullEmergencyResponse, error) {
	m := new(traits.PullEmergencyResponse)
	if err := c.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type pullEmergencyServerWrapper struct {
	grpc.ServerStream
}

func (s *pullEmergencyServerWrapper) Send(response *traits.PullEmergencyResponse) error {
	return s.ServerStream.SendMsg(response)
}