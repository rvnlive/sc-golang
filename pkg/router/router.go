package router

import (
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Router tracks a registry of gRPC clients.
// Typically used by code generated via the protoc-gen-router plugin.
type Router interface {
	// Add adds a named client to this Router.
	// Add returns the old client associated with this name, or nil if there wasn't one.
	// If HoldsType returns false for the given client, this will panic.
	Add(name string, client interface{}) interface{}
	// HoldsType returns true if this Router holds clients of the specified type.
	HoldsType(client interface{}) bool
	// Remove removes and returns a named client.
	Remove(name string) interface{}
	// Has returns true if this Router has a client with the given name.
	Has(name string) bool
	// Get returns the client for the given name.
	// An error will be returned if no such client exists.
	Get(name string) (interface{}, error)
}

type router struct {
	mu       sync.RWMutex
	registry map[string]interface{} // of type MyServiceClient
	factory  Factory
	fallback Factory

	onCommit OnCommit
}
type Factory func(string) (interface{}, error) // returns the type MyServiceClient

// NewRouter creates a new instance of Router with the given options.
func NewRouter(opts ...Option) Router {
	r := &router{
		registry: make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *router) Add(name string, client interface{}) interface{} {
	r.mu.Lock()
	defer r.mu.Unlock()
	old := r.registry[name]
	r.registry[name] = client
	return old
}

func (r *router) HoldsType(_ interface{}) bool {
	return true
}

func (r *router) Remove(name string) interface{} {
	r.mu.Lock()
	defer r.mu.Unlock()
	old := r.registry[name]
	delete(r.registry, name)
	return old
}

func (r *router) Has(name string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, exists := r.registry[name]
	return exists
}

// Get returns the client identified by the given name.
// If the name is not recognised by r
//  1. A fallback is checked, configured via WithFallback, if found it is returned, else
//  2. The factory is invoked to create a new client, configured via WithFactory.
//     If the factory successfully creates a client, and no concurrent call already created a client, it is remembered
//     and the callback registered via WithOnCommit is notified.
// If no client can be found or created, this returns an error suitable for return by a gRPC method,
// i.e. one representing codes.NotFound.
//
// Note, no locks are held when invoking fallbacks, factories, or callbacks.
func (r *router) Get(name string) (child interface{}, err error) {
	r.mu.RLock()
	child, exists := r.registry[name]
	r.mu.RUnlock()
	if !exists {
		child, exists, err = invoke(name, r.fallback)
	}
	if !exists {
		child, exists, err = invoke(name, r.factory)
		if exists {
			r.mu.Lock()
			// check again
			var newChildRemembered bool
			child2, exists2 := r.registry[name]
			if exists2 {
				child = child2
			} else {
				newChildRemembered = true
				r.registry[name] = child
			}
			r.mu.Unlock()

			if newChildRemembered && r.onCommit != nil {
				r.onCommit(name, child)
			}
		}
	}

	if !exists {
		return nil, status.Error(codes.NotFound, name)
	}
	return
}
func invoke(name string, f Factory) (interface{}, bool, error) {
	if f == nil {
		return nil, false, nil
	}
	child, err := f(name)
	return child, child != nil && err == nil, err
}

type Option func(r *router)

// WithFactory configures a Router to call the given function when Get is called and no existing client is known.
// Prefer using the generated WithMyServiceClientFactory methods in the trait packages.
// The given factory may be called multiple times with the same name if concurrent access is performed.
// Only one returned client will be remembered.
// Use WithOnCommit if you need to trigger side effects as part of your client creation.
func WithFactory(f Factory) Option {
	return func(r *router) {
		r.factory = f
	}
}

// WithFallback configures a Router to ask the given function when Get is called and no existing client is known.
// If WithFallback and WithFactory are both configured, WithFallback will be called first, only using WithFactory if
// WithFallback returns nil or an error.
func WithFallback(f Factory) Option {
	return func(r *router) {
		r.fallback = f
	}
}

// OnCommit is a callback function for use in WithOnCommit.
type OnCommit func(name string, client interface{})

// WithOnCommit registers a func that will be called with the value remembered as part of a WithFactory Factory call.
// Use this if you want to register or otherwise setup side effects for Factory created entries.
func WithOnCommit(onCommit OnCommit) Option {
	return func(r *router) {
		r.onCommit = onCommit
	}
}
