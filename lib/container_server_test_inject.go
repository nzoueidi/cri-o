// +build test
// All *_inject.go files are meant to be used by tests only. Purpose of this
// files is to provide a way to inject mocked data into the current setup.

package lib

import (
	"github.com/cri-o/cri-o/oci"
	"github.com/cri-o/cri-o/pkg/storage"
)

// SetStorageRuntimeServer sets the runtime server for the ContainerServer
func (c *ContainerServer) SetStorageRuntimeServer(server storage.RuntimeServer) {
	c.storageRuntimeServer = server
}

// SetStorageImageServer sets the ImageServer for the ContainerServer
func (c *ContainerServer) SetStorageImageServer(server storage.ImageServer) {
	c.storageImageServer = server
}

// SetRuntime can be used to explicitly specify a runtime.
func (c *ContainerServer) SetRuntime(runtime oci.RuntimeImpl) {
	c.runtime = runtime
}
