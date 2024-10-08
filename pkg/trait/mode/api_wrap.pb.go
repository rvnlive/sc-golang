// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package mode

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
)

// WrapApi	adapts a traits.ModeApiServer	and presents it as a traits.ModeApiClient
func WrapApi(server traits.ModeApiServer) traits.ModeApiClient {
	conn := wrap.ServerToClient(traits.ModeApi_ServiceDesc, server)
	client := traits.NewModeApiClient(conn)
	return &apiWrapper{
		ModeApiClient: client,
		server:        server,
	}
}

type apiWrapper struct {
	traits.ModeApiClient

	server traits.ModeApiServer
}

// compile time check that we implement the interface we need
var _ traits.ModeApiClient = (*apiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *apiWrapper) UnwrapServer() traits.ModeApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *apiWrapper) Unwrap() any {
	return w.UnwrapServer()
}
