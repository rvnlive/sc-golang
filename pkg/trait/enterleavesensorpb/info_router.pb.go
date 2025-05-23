// Code generated by protoc-gen-router. DO NOT EDIT.

package enterleavesensorpb

import (
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// InfoRouter is a traits.EnterLeaveSensorInfoServer that allows routing named requests to specific traits.EnterLeaveSensorInfoClient
type InfoRouter struct {
	traits.UnimplementedEnterLeaveSensorInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.EnterLeaveSensorInfoServer = (*InfoRouter)(nil)

func NewInfoRouter(opts ...router.Option) *InfoRouter {
	return &InfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithEnterLeaveSensorInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithEnterLeaveSensorInfoClientFactory(f func(name string) (traits.EnterLeaveSensorInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *InfoRouter) Register(server grpc.ServiceRegistrar) {
	traits.RegisterEnterLeaveSensorInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.EnterLeaveSensorInfoClient.
func (r *InfoRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.EnterLeaveSensorInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *InfoRouter) HoldsType(client any) bool {
	_, ok := client.(traits.EnterLeaveSensorInfoClient)
	return ok
}

func (r *InfoRouter) AddEnterLeaveSensorInfoClient(name string, client traits.EnterLeaveSensorInfoClient) traits.EnterLeaveSensorInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.EnterLeaveSensorInfoClient)
}

func (r *InfoRouter) RemoveEnterLeaveSensorInfoClient(name string) traits.EnterLeaveSensorInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.EnterLeaveSensorInfoClient)
}

func (r *InfoRouter) GetEnterLeaveSensorInfoClient(name string) (traits.EnterLeaveSensorInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.EnterLeaveSensorInfoClient), nil
}
