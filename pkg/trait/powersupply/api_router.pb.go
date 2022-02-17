// Code generated by protoc-gen-router. DO NOT EDIT.

package powersupply

import (
	context "context"
	fmt "fmt"
	traits "github.com/smart-core-os/sc-api/go/traits"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

// ApiRouter is a traits.PowerSupplyApiServer that allows routing named requests to specific traits.PowerSupplyApiClient
type ApiRouter struct {
	traits.UnimplementedPowerSupplyApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ traits.PowerSupplyApiServer = (*ApiRouter)(nil)

func NewApiRouter(opts ...router.Option) *ApiRouter {
	return &ApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithPowerSupplyApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithPowerSupplyApiClientFactory(f func(name string) (traits.PowerSupplyApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (interface{}, error) {
		return f(name)
	})
}

func (r *ApiRouter) Register(server *grpc.Server) {
	traits.RegisterPowerSupplyApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type traits.PowerSupplyApiClient.
func (r *ApiRouter) Add(name string, client interface{}) interface{} {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a traits.PowerSupplyApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *ApiRouter) HoldsType(client interface{}) bool {
	_, ok := client.(traits.PowerSupplyApiClient)
	return ok
}

func (r *ApiRouter) AddPowerSupplyApiClient(name string, client traits.PowerSupplyApiClient) traits.PowerSupplyApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(traits.PowerSupplyApiClient)
}

func (r *ApiRouter) RemovePowerSupplyApiClient(name string) traits.PowerSupplyApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(traits.PowerSupplyApiClient)
}

func (r *ApiRouter) GetPowerSupplyApiClient(name string) (traits.PowerSupplyApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(traits.PowerSupplyApiClient), nil
}

func (r *ApiRouter) GetPowerCapacity(ctx context.Context, request *traits.GetPowerCapacityRequest) (*traits.PowerCapacity, error) {
	child, err := r.GetPowerSupplyApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetPowerCapacity(ctx, request)
}

func (r *ApiRouter) PullPowerCapacity(request *traits.PullPowerCapacityRequest, server traits.PowerSupplyApi_PullPowerCapacityServer) error {
	child, err := r.GetPowerSupplyApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullPowerCapacity(reqCtx, request)
	if err != nil {
		return err
	}

	// send the stream header
	header, err := stream.Header()
	if err != nil {
		return err
	}
	if err = server.SendHeader(header); err != nil {
		return err
	}

	// send all the messages
	// false means the error is from the child, true means the error is from the caller
	var callerError bool
	for {
		// Impl note: we could improve throughput here by issuing the Recv and Send in different goroutines, but we're doing
		// it synchronously until we have a need to change the behaviour

		var msg *traits.PullPowerCapacityResponse
		msg, err = stream.Recv()
		if err != nil {
			break
		}

		err = server.Send(msg)
		if err != nil {
			callerError = true
			break
		}
	}

	// err is guaranteed to be non-nil as it's the only way to exit the loop
	if callerError {
		// cancel the request
		reqDone()
		return err
	} else {
		if trailer := stream.Trailer(); trailer != nil {
			server.SetTrailer(trailer)
		}
		if err == io.EOF {
			return nil
		}
		return err
	}
}

func (r *ApiRouter) ListDrawNotifications(ctx context.Context, request *traits.ListDrawNotificationsRequest) (*traits.ListDrawNotificationsResponse, error) {
	child, err := r.GetPowerSupplyApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.ListDrawNotifications(ctx, request)
}

func (r *ApiRouter) CreateDrawNotification(ctx context.Context, request *traits.CreateDrawNotificationRequest) (*traits.DrawNotification, error) {
	child, err := r.GetPowerSupplyApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.CreateDrawNotification(ctx, request)
}

func (r *ApiRouter) UpdateDrawNotification(ctx context.Context, request *traits.UpdateDrawNotificationRequest) (*traits.DrawNotification, error) {
	child, err := r.GetPowerSupplyApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.UpdateDrawNotification(ctx, request)
}

func (r *ApiRouter) DeleteDrawNotification(ctx context.Context, request *traits.DeleteDrawNotificationRequest) (*emptypb.Empty, error) {
	child, err := r.GetPowerSupplyApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DeleteDrawNotification(ctx, request)
}

func (r *ApiRouter) PullDrawNotifications(request *traits.PullDrawNotificationsRequest, server traits.PowerSupplyApi_PullDrawNotificationsServer) error {
	child, err := r.GetPowerSupplyApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullDrawNotifications(reqCtx, request)
	if err != nil {
		return err
	}

	// send the stream header
	header, err := stream.Header()
	if err != nil {
		return err
	}
	if err = server.SendHeader(header); err != nil {
		return err
	}

	// send all the messages
	// false means the error is from the child, true means the error is from the caller
	var callerError bool
	for {
		// Impl note: we could improve throughput here by issuing the Recv and Send in different goroutines, but we're doing
		// it synchronously until we have a need to change the behaviour

		var msg *traits.PullDrawNotificationsResponse
		msg, err = stream.Recv()
		if err != nil {
			break
		}

		err = server.Send(msg)
		if err != nil {
			callerError = true
			break
		}
	}

	// err is guaranteed to be non-nil as it's the only way to exit the loop
	if callerError {
		// cancel the request
		reqDone()
		return err
	} else {
		if trailer := stream.Trailer(); trailer != nil {
			server.SetTrailer(trailer)
		}
		if err == io.EOF {
			return nil
		}
		return err
	}
}
