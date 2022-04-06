// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package inputselect

import (
	context "context"
	traits "github.com/smart-core-os/sc-api/go/traits"
	grpc "google.golang.org/grpc"
)

// WrapInfo	adapts a traits.InputSelectInfoServer	and presents it as a traits.InputSelectInfoClient
func WrapInfo(server traits.InputSelectInfoServer) traits.InputSelectInfoClient {
	return &infoWrapper{server}
}

type infoWrapper struct {
	server traits.InputSelectInfoServer
}

// compile time check that we implement the interface we need
var _ traits.InputSelectInfoClient = (*infoWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *infoWrapper) UnwrapServer() traits.InputSelectInfoServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *infoWrapper) Unwrap() interface{} {
	return w.UnwrapServer()
}

func (w *infoWrapper) DescribeInput(ctx context.Context, req *traits.DescribeInputRequest, _ ...grpc.CallOption) (*traits.InputSupport, error) {
	return w.server.DescribeInput(ctx, req)
}