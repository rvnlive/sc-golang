// Code generated by protoc-gen-router. DO NOT EDIT.

package speaker

import (
	context "context"
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// InfoRouter is a traits.SpeakerInfoServer that allows routing named requests to specific traits.SpeakerInfoClient
type InfoRouter struct {
	traits.UnimplementedSpeakerInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.SpeakerInfoServer = (*InfoRouter)(nil)

func NewInfoRouter(opts ...router.Option) *InfoRouter {
	return &InfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithSpeakerInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithSpeakerInfoClientFactory(f func(name string) (traits.SpeakerInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (interface{}, error) {
		return f(name)
	})
}

func (r *InfoRouter) Register(server *grpc.Server) {
	traits.RegisterSpeakerInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.SpeakerInfoClient.
func (r *InfoRouter) Add(name string, client interface{}) interface{} {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.SpeakerInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *InfoRouter) HoldsType(client interface{}) bool {
	_, ok := client.(traits.SpeakerInfoClient)
	return ok
}

func (r *InfoRouter) AddSpeakerInfoClient(name string, client traits.SpeakerInfoClient) traits.SpeakerInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.SpeakerInfoClient)
}

func (r *InfoRouter) RemoveSpeakerInfoClient(name string) traits.SpeakerInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.SpeakerInfoClient)
}

func (r *InfoRouter) GetSpeakerInfoClient(name string) (traits.SpeakerInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.SpeakerInfoClient), nil
}

func (r *InfoRouter) DescribeVolume(ctx context.Context, request *traits.DescribeVolumeRequest) (*traits.VolumeSupport, error) {
	child, err := r.GetSpeakerInfoClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DescribeVolume(ctx, request)
}
