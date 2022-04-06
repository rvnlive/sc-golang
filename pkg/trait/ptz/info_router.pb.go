// Code generated by protoc-gen-router. DO NOT EDIT.

package ptz

import (
	context "context"
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// InfoRouter is a traits.PtzInfoServer that allows routing named requests to specific traits.PtzInfoClient
type InfoRouter struct {
	traits.UnimplementedPtzInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.PtzInfoServer = (*InfoRouter)(nil)

func NewInfoRouter(opts ...router.Option) *InfoRouter {
	return &InfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithPtzInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithPtzInfoClientFactory(f func(name string) (traits.PtzInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (interface{}, error) {
		return f(name)
	})
}

func (r *InfoRouter) Register(server *grpc.Server) {
	traits.RegisterPtzInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.PtzInfoClient.
func (r *InfoRouter) Add(name string, client interface{}) interface{} {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.PtzInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *InfoRouter) HoldsType(client interface{}) bool {
	_, ok := client.(traits.PtzInfoClient)
	return ok
}

func (r *InfoRouter) AddPtzInfoClient(name string, client traits.PtzInfoClient) traits.PtzInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.PtzInfoClient)
}

func (r *InfoRouter) RemovePtzInfoClient(name string) traits.PtzInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.PtzInfoClient)
}

func (r *InfoRouter) GetPtzInfoClient(name string) (traits.PtzInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.PtzInfoClient), nil
}

func (r *InfoRouter) DescribePtz(ctx context.Context, request *traits.DescribePtzRequest) (*traits.PtzSupport, error) {
	child, err := r.GetPtzInfoClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DescribePtz(ctx, request)
}