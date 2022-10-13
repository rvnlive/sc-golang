// Code generated by protoc-gen-router. DO NOT EDIT.

package brightnesssensor

import (
	context "context"
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// InfoRouter is a traits.BrightnessSensorInfoServer that allows routing named requests to specific traits.BrightnessSensorInfoClient
type InfoRouter struct {
	traits.UnimplementedBrightnessSensorInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.BrightnessSensorInfoServer = (*InfoRouter)(nil)

func NewInfoRouter(opts ...router.Option) *InfoRouter {
	return &InfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithBrightnessSensorInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithBrightnessSensorInfoClientFactory(f func(name string) (traits.BrightnessSensorInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *InfoRouter) Register(server *grpc.Server) {
	traits.RegisterBrightnessSensorInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.BrightnessSensorInfoClient.
func (r *InfoRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.BrightnessSensorInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *InfoRouter) HoldsType(client any) bool {
	_, ok := client.(traits.BrightnessSensorInfoClient)
	return ok
}

func (r *InfoRouter) AddBrightnessSensorInfoClient(name string, client traits.BrightnessSensorInfoClient) traits.BrightnessSensorInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.BrightnessSensorInfoClient)
}

func (r *InfoRouter) RemoveBrightnessSensorInfoClient(name string) traits.BrightnessSensorInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.BrightnessSensorInfoClient)
}

func (r *InfoRouter) GetBrightnessSensorInfoClient(name string) (traits.BrightnessSensorInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.BrightnessSensorInfoClient), nil
}

func (r *InfoRouter) DescribeAmbientBrightness(ctx context.Context, request *traits.DescribeAmbientBrightnessRequest) (*traits.AmbientBrightnessSupport, error) {
	child, err := r.GetBrightnessSensorInfoClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DescribeAmbientBrightness(ctx, request)
}
