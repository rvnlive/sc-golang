// Code generated by protoc-gen-router. DO NOT EDIT.

package light

import (
	context "context"
	traits "github.com/smart-core-os/sc-api/go/traits"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	sync "sync"
)

// InfoRouter is a traits.LightInfoServer that allows routing named requests to specific traits.LightInfoClient
type InfoRouter struct {
	traits.UnimplementedLightInfoServer

	mu       sync.Mutex
	registry map[string]traits.LightInfoClient
	// Factory can be used to dynamically create api clients if requests come in for devices we haven't seen.
	Factory func(string) (traits.LightInfoClient, error)
}

// compile time check that we implement the interface we need
var _ traits.LightInfoServer = (*InfoRouter)(nil)

func NewInfoRouter() *InfoRouter {
	return &InfoRouter{
		registry: make(map[string]traits.LightInfoClient),
	}
}

func (r *InfoRouter) Register(server *grpc.Server) {
	traits.RegisterLightInfoServer(server, r)
}

func (r *InfoRouter) Add(name string, client traits.LightInfoClient) traits.LightInfoClient {
	r.mu.Lock()
	defer r.mu.Unlock()
	old := r.registry[name]
	r.registry[name] = client
	return old
}

func (r *InfoRouter) Remove(name string) traits.LightInfoClient {
	r.mu.Lock()
	defer r.mu.Unlock()
	old := r.registry[name]
	delete(r.registry, name)
	return old
}

func (r *InfoRouter) Has(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.registry[name]
	return exists
}

func (r *InfoRouter) Get(name string) (traits.LightInfoClient, error) {
	r.mu.Lock()
	child, exists := r.registry[name]
	defer r.mu.Unlock()
	if !exists {
		if r.Factory != nil {
			child, err := r.Factory(name)
			if err != nil {
				return nil, err
			}
			r.registry[name] = child
			return child, nil
		}
		return nil, status.Error(codes.NotFound, name)
	}
	return child, nil
}

func (r *InfoRouter) DescribeBrightness(ctx context.Context, request *traits.DescribeBrightnessRequest) (*traits.BrightnessSupport, error) {
	child, err := r.Get(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DescribeBrightness(ctx, request)
}
