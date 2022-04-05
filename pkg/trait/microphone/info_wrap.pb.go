// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package microphone

import (
	context "context"
	traits "github.com/smart-core-os/sc-api/go/traits"
	grpc "google.golang.org/grpc"
)

// WrapInfo	adapts a traits.MicrophoneInfoServer	and presents it as a traits.MicrophoneInfoClient
func WrapInfo(server traits.MicrophoneInfoServer) traits.MicrophoneInfoClient {
	return &infoWrapper{server}
}

type infoWrapper struct {
	server traits.MicrophoneInfoServer
}

// compile time check that we implement the interface we need
var _ traits.MicrophoneInfoClient = (*infoWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *infoWrapper) UnwrapServer() traits.MicrophoneInfoServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *infoWrapper) Unwrap() interface{} {
	return w.UnwrapServer()
}

func (w *infoWrapper) DescribeGain(ctx context.Context, req *traits.DescribeGainRequest, _ ...grpc.CallOption) (*traits.GainSupport, error) {
	return w.server.DescribeGain(ctx, req)
}
