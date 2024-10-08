// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package count

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
)

// WrapInfo	adapts a traits.CountInfoServer	and presents it as a traits.CountInfoClient
func WrapInfo(server traits.CountInfoServer) traits.CountInfoClient {
	conn := wrap.ServerToClient(traits.CountInfo_ServiceDesc, server)
	client := traits.NewCountInfoClient(conn)
	return &infoWrapper{
		CountInfoClient: client,
		server:          server,
	}
}

type infoWrapper struct {
	traits.CountInfoClient

	server traits.CountInfoServer
}

// compile time check that we implement the interface we need
var _ traits.CountInfoClient = (*infoWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *infoWrapper) UnwrapServer() traits.CountInfoServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *infoWrapper) Unwrap() any {
	return w.UnwrapServer()
}
