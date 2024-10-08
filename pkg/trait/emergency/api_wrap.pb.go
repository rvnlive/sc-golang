// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package emergency

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
)

// WrapApi	adapts a traits.EmergencyApiServer	and presents it as a traits.EmergencyApiClient
func WrapApi(server traits.EmergencyApiServer) traits.EmergencyApiClient {
	conn := wrap.ServerToClient(traits.EmergencyApi_ServiceDesc, server)
	client := traits.NewEmergencyApiClient(conn)
	return &apiWrapper{
		EmergencyApiClient: client,
		server:             server,
	}
}

type apiWrapper struct {
	traits.EmergencyApiClient

	server traits.EmergencyApiServer
}

// compile time check that we implement the interface we need
var _ traits.EmergencyApiClient = (*apiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *apiWrapper) UnwrapServer() traits.EmergencyApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *apiWrapper) Unwrap() any {
	return w.UnwrapServer()
}
