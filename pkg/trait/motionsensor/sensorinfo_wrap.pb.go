// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package motionsensor

import (
	context "context"
	traits "github.com/smart-core-os/sc-api/go/traits"
	grpc "google.golang.org/grpc"
)

// WrapSensorInfo	adapts a traits.MotionSensorSensorInfoServer	and presents it as a traits.MotionSensorSensorInfoClient
func WrapSensorInfo(server traits.MotionSensorSensorInfoServer) traits.MotionSensorSensorInfoClient {
	return &sensorInfoWrapper{server}
}

type sensorInfoWrapper struct {
	server traits.MotionSensorSensorInfoServer
}

// compile time check that we implement the interface we need
var _ traits.MotionSensorSensorInfoClient = (*sensorInfoWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *sensorInfoWrapper) UnwrapServer() traits.MotionSensorSensorInfoServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *sensorInfoWrapper) Unwrap() interface{} {
	return w.UnwrapServer()
}

func (w *sensorInfoWrapper) DescribeMotionDetection(ctx context.Context, req *traits.DescribeMotionDetectionRequest, _ ...grpc.CallOption) (*traits.MotionDetectionSupport, error) {
	return w.server.DescribeMotionDetection(ctx, req)
}
