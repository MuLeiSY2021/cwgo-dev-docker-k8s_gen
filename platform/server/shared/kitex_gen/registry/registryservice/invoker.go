// Code generated by Kitex v0.6.1. DO NOT EDIT.

package registryservice

import (
	registry "github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/registry"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler registry.RegistryService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
