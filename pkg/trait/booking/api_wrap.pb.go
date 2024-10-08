// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package booking

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
)

// WrapApi	adapts a traits.BookingApiServer	and presents it as a traits.BookingApiClient
func WrapApi(server traits.BookingApiServer) traits.BookingApiClient {
	conn := wrap.ServerToClient(traits.BookingApi_ServiceDesc, server)
	client := traits.NewBookingApiClient(conn)
	return &apiWrapper{
		BookingApiClient: client,
		server:           server,
	}
}

type apiWrapper struct {
	traits.BookingApiClient

	server traits.BookingApiServer
}

// compile time check that we implement the interface we need
var _ traits.BookingApiClient = (*apiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *apiWrapper) UnwrapServer() traits.BookingApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *apiWrapper) Unwrap() any {
	return w.UnwrapServer()
}
