// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package motionsensor

import (
	traits "github.com/smart-core-os/sc-api/go/traits"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
)

// WrapSensorInfo	adapts a traits.MotionSensorSensorInfoServer	and presents it as a traits.MotionSensorSensorInfoClient
func WrapSensorInfo(server traits.MotionSensorSensorInfoServer) traits.MotionSensorSensorInfoClient {
	conn := wrap.ServerToClient(traits.MotionSensorSensorInfo_ServiceDesc, server)
	client := traits.NewMotionSensorSensorInfoClient(conn)
	return &sensorInfoWrapper{
		MotionSensorSensorInfoClient: client,
		server:                       server,
	}
}

type sensorInfoWrapper struct {
	traits.MotionSensorSensorInfoClient

	server traits.MotionSensorSensorInfoServer
}

// compile time check that we implement the interface we need
var _ traits.MotionSensorSensorInfoClient = (*sensorInfoWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *sensorInfoWrapper) UnwrapServer() traits.MotionSensorSensorInfoServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *sensorInfoWrapper) Unwrap() any {
	return w.UnwrapServer()
}
