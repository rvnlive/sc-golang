// Code generated by protoc-gen-router. DO NOT EDIT.

package occupancysensor

import (
	context "context"
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// InfoRouter is a traits.OccupancySensorInfoServer that allows routing named requests to specific traits.OccupancySensorInfoClient
type InfoRouter struct {
	traits.UnimplementedOccupancySensorInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.OccupancySensorInfoServer = (*InfoRouter)(nil)

func NewInfoRouter(opts ...router.Option) *InfoRouter {
	return &InfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithOccupancySensorInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithOccupancySensorInfoClientFactory(f func(name string) (traits.OccupancySensorInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *InfoRouter) Register(server *grpc.Server) {
	traits.RegisterOccupancySensorInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.OccupancySensorInfoClient.
func (r *InfoRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.OccupancySensorInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *InfoRouter) HoldsType(client any) bool {
	_, ok := client.(traits.OccupancySensorInfoClient)
	return ok
}

func (r *InfoRouter) AddOccupancySensorInfoClient(name string, client traits.OccupancySensorInfoClient) traits.OccupancySensorInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.OccupancySensorInfoClient)
}

func (r *InfoRouter) RemoveOccupancySensorInfoClient(name string) traits.OccupancySensorInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.OccupancySensorInfoClient)
}

func (r *InfoRouter) GetOccupancySensorInfoClient(name string) (traits.OccupancySensorInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.OccupancySensorInfoClient), nil
}

func (r *InfoRouter) DescribeOccupancy(ctx context.Context, request *traits.DescribeOccupancyRequest) (*traits.OccupancySupport, error) {
	child, err := r.GetOccupancySensorInfoClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DescribeOccupancy(ctx, request)
}
