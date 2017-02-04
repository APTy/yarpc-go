// Code generated by thriftrw-plugin-yarpc
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package readonlystoreclient

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common/baseserviceclient"
	"go.uber.org/yarpc"
)

// Interface is a client for the ReadOnlyStore service.
type Interface interface {
	baseserviceclient.Interface

	Integer(
		ctx context.Context,
		Key *string,
		opts ...yarpc.CallOption,
	) (int64, error)
}

// New builds a new client for the ReadOnlyStore service.
//
// 	client := readonlystoreclient.New(dispatcher.ClientConfig("readonlystore"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "ReadOnlyStore",
			ClientConfig: c,
		}, opts...),
		Interface: baseserviceclient.New(c),
	}
}

func init() {
	yarpc.RegisterClientBuilder(func(c transport.ClientConfig) Interface {
		return New(c)
	})
}

type client struct {
	baseserviceclient.Interface

	c thrift.Client
}

func (c client) Integer(
	ctx context.Context,
	_Key *string,
	opts ...yarpc.CallOption,
) (success int64, err error) {

	args := atomic.ReadOnlyStore_Integer_Helper.Args(_Key)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result atomic.ReadOnlyStore_Integer_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = atomic.ReadOnlyStore_Integer_Helper.UnwrapResponse(&result)
	return
}
