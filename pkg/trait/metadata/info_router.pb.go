// Code generated by protoc-gen-router. DO NOT EDIT.

package metadata

import (
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// InfoRouter is a traits.MetadataInfoServer that allows routing named requests to specific traits.MetadataInfoClient
type InfoRouter struct {
	traits.UnimplementedMetadataInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.MetadataInfoServer = (*InfoRouter)(nil)

func NewInfoRouter(opts ...router.Option) *InfoRouter {
	return &InfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithMetadataInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithMetadataInfoClientFactory(f func(name string) (traits.MetadataInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (interface{}, error) {
		return f(name)
	})
}

func (r *InfoRouter) Register(server *grpc.Server) {
	traits.RegisterMetadataInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.MetadataInfoClient.
func (r *InfoRouter) Add(name string, client interface{}) interface{} {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.MetadataInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *InfoRouter) HoldsType(client interface{}) bool {
	_, ok := client.(traits.MetadataInfoClient)
	return ok
}

func (r *InfoRouter) AddMetadataInfoClient(name string, client traits.MetadataInfoClient) traits.MetadataInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.MetadataInfoClient)
}

func (r *InfoRouter) RemoveMetadataInfoClient(name string) traits.MetadataInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.MetadataInfoClient)
}

func (r *InfoRouter) GetMetadataInfoClient(name string) (traits.MetadataInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.MetadataInfoClient), nil
}
