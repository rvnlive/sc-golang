package onoff

import (
	"context"

	"github.com/smart-core-os/sc-api/go/traits"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ModelServer struct {
	traits.UnimplementedOnOffApiServer
	model *Model
}

func NewModelServer(model *Model) *ModelServer {
	return &ModelServer{model: model}
}

func (s *ModelServer) Register(server *grpc.Server) {
	traits.RegisterOnOffApiServer(server, s)
}

func (s *ModelServer) GetOnOff(_ context.Context, _ *traits.GetOnOffRequest) (*traits.OnOff, error) {
	return s.model.GetOnOff()
}

func (s *ModelServer) UpdateOnOff(_ context.Context, request *traits.UpdateOnOffRequest) (*traits.OnOff, error) {
	return s.model.UpdateOnOff(request.OnOff)
}

func (s *ModelServer) PullOnOff(request *traits.PullOnOffRequest, server traits.OnOffApi_PullOnOffServer) error {
	updates, done := s.model.PullOnOff(server.Context(), nil)
	defer done()

	for {
		select {
		case <-server.Context().Done():
			return status.FromContextError(server.Context().Err()).Err()
		case update := <-updates:
			change := &traits.PullOnOffResponse_Change{
				Name:       request.Name,
				ChangeTime: timestamppb.New(update.ChangeTime),
				OnOff:      update.Value,
			}

			err := server.Send(&traits.PullOnOffResponse{
				Changes: []*traits.PullOnOffResponse_Change{change},
			})
			if err != nil {
				return err
			}
		}
	}
}