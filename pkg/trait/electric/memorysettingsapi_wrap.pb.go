// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package electric

import (
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
)

// WrapMemorySettingsApi	adapts a MemorySettingsApiServer	and presents it as a MemorySettingsApiClient
func WrapMemorySettingsApi(server MemorySettingsApiServer) MemorySettingsApiClient {
	conn := wrap.ServerToClient(MemorySettingsApi_ServiceDesc, server)
	client := NewMemorySettingsApiClient(conn)
	return &memorySettingsApiWrapper{
		MemorySettingsApiClient: client,
		server:                  server,
	}
}

type memorySettingsApiWrapper struct {
	MemorySettingsApiClient

	server MemorySettingsApiServer
}

// compile time check that we implement the interface we need
var _ MemorySettingsApiClient = (*memorySettingsApiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *memorySettingsApiWrapper) UnwrapServer() MemorySettingsApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *memorySettingsApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}
