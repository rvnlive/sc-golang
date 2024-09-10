// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package publication

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
)

// WrapApi	adapts a traits.PublicationApiServer	and presents it as a traits.PublicationApiClient
func WrapApi(server traits.PublicationApiServer) traits.PublicationApiClient {
	conn := wrap.ServerToClient(traits.PublicationApi_ServiceDesc, server)
	client := traits.NewPublicationApiClient(conn)
	return &apiWrapper{
		PublicationApiClient: client,
		server:               server,
	}
}

type apiWrapper struct {
	traits.PublicationApiClient

	server traits.PublicationApiServer
}

// compile time check that we implement the interface we need
var _ traits.PublicationApiClient = (*apiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *apiWrapper) UnwrapServer() traits.PublicationApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *apiWrapper) Unwrap() any {
	return w.UnwrapServer()
}
