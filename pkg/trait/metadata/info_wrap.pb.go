// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package metadata

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
)

// WrapInfo	adapts a traits.MetadataInfoServer	and presents it as a traits.MetadataInfoClient
func WrapInfo(server traits.MetadataInfoServer) traits.MetadataInfoClient {
	return &infoWrapper{server}
}

type infoWrapper struct {
	server traits.MetadataInfoServer
}

// compile time check that we implement the interface we need
var _ traits.MetadataInfoClient = (*infoWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *infoWrapper) UnwrapServer() traits.MetadataInfoServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *infoWrapper) Unwrap() interface{} {
	return w.UnwrapServer()
}
