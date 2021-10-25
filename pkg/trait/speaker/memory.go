package speaker

import (
	"context"

	"github.com/smart-core-os/sc-golang/pkg/memory"

	"github.com/smart-core-os/sc-api/go/traits"
	"github.com/smart-core-os/sc-api/go/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type MemoryDevice struct {
	traits.UnimplementedSpeakerApiServer
	volume *memory.Resource
}

func NewMemoryDevice(initialState *types.AudioLevel) *MemoryDevice {
	return &MemoryDevice{
		volume: memory.NewResource(memory.WithInitialValue(initialState)),
	}
}

func (s *MemoryDevice) Register(server *grpc.Server) {
	traits.RegisterSpeakerApiServer(server, s)
}

func (s *MemoryDevice) GetVolume(_ context.Context, _ *traits.GetSpeakerVolumeRequest) (*types.AudioLevel, error) {
	return s.volume.Get().(*types.AudioLevel), nil
}

func (s *MemoryDevice) UpdateVolume(_ context.Context, request *traits.UpdateSpeakerVolumeRequest) (*types.AudioLevel, error) {
	newValue, err := s.volume.Set(request.Volume, memory.WithUpdateMask(request.UpdateMask), memory.InterceptBefore(func(old, change proto.Message) {
		if request.Delta {
			val := old.(*types.AudioLevel)
			delta := change.(*types.AudioLevel)
			delta.Gain += val.Gain
		}
	}))
	if err != nil {
		return nil, err
	}
	return newValue.(*types.AudioLevel), nil
}

func (s *MemoryDevice) PullVolume(request *traits.PullSpeakerVolumeRequest, server traits.SpeakerApi_PullVolumeServer) error {
	changes, done := s.volume.OnUpdate(server.Context())
	defer done()

	for {
		select {
		case <-server.Context().Done():
			return status.FromContextError(server.Context().Err()).Err()
		case change := <-changes:
			typedChange := &types.AudioLevelChange{
				Name:       request.Name,
				Level:      change.Value.(*types.AudioLevel),
				ChangeTime: change.ChangeTime,
			}
			err := server.Send(&traits.PullSpeakerVolumeResponse{
				Changes: []*types.AudioLevelChange{typedChange},
			})
			if err != nil {
				return err
			}
		}
	}
}