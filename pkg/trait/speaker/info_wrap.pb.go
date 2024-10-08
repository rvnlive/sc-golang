// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package speaker

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
)

// WrapInfo	adapts a traits.SpeakerInfoServer	and presents it as a traits.SpeakerInfoClient
func WrapInfo(server traits.SpeakerInfoServer) traits.SpeakerInfoClient {
	conn := wrap.ServerToClient(traits.SpeakerInfo_ServiceDesc, server)
	client := traits.NewSpeakerInfoClient(conn)
	return &infoWrapper{
		SpeakerInfoClient: client,
		server:            server,
	}
}

type infoWrapper struct {
	traits.SpeakerInfoClient

	server traits.SpeakerInfoServer
}

// compile time check that we implement the interface we need
var _ traits.SpeakerInfoClient = (*infoWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *infoWrapper) UnwrapServer() traits.SpeakerInfoServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *infoWrapper) Unwrap() any {
	return w.UnwrapServer()
}
