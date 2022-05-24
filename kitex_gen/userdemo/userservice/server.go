// Code generated by Kitex v0.3.1. DO NOT EDIT.
package userservice

import (
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler userdemo.UserService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
