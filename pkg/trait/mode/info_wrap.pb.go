// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package mode

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
)

// WrapInfo	adapts a traits.ModeInfoServer	and presents it as a traits.ModeInfoClient
func WrapInfo(server traits.ModeInfoServer) traits.ModeInfoClient {
	conn := wrap.ServerToClient(traits.ModeInfo_ServiceDesc, server)
	client := traits.NewModeInfoClient(conn)
	return &infoWrapper{
		ModeInfoClient: client,
		server:         server,
	}
}

type infoWrapper struct {
	traits.ModeInfoClient

	server traits.ModeInfoServer
}

// compile time check that we implement the interface we need
var _ traits.ModeInfoClient = (*infoWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *infoWrapper) UnwrapServer() traits.ModeInfoServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *infoWrapper) Unwrap() any {
	return w.UnwrapServer()
}
