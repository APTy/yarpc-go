// Code generated by protoc-gen-yarpc-go
// source: internal/examples/protobuf/examplepb/example.proto
// DO NOT EDIT!

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

package examplepb

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/yarpcproto"
)

var _ = ioutil.NopCloser

// KeyValueYARPCClient is the YARPC client-side interface for the KeyValue service.
type KeyValueYARPCClient interface {
	GetValue(context.Context, *GetValueRequest, ...yarpc.CallOption) (*GetValueResponse, error)
	SetValue(context.Context, *SetValueRequest, ...yarpc.CallOption) (*SetValueResponse, error)
}

// NewKeyValueYARPCClient builds a new YARPC client for the KeyValue service.
func NewKeyValueYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) KeyValueYARPCClient {
	return &_KeyValueYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.internal.examples.protobuf.example.KeyValue",
			ClientConfig: clientConfig,
			Options:      options,
		},
	)}
}

// KeyValueYARPCServer is the YARPC server-side interface for the KeyValue service.
type KeyValueYARPCServer interface {
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)
}

// BuildKeyValueYARPCProcedures prepares an implementation of the KeyValue service for YARPC registration.
func BuildKeyValueYARPCProcedures(server KeyValueYARPCServer) []transport.Procedure {
	handler := &_KeyValueYARPCHandler{server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.yarpc.internal.examples.protobuf.example.KeyValue",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "GetValue",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:     handler.GetValue,
							NewRequest: newKeyValueServiceGetValueYARPCRequest,
						},
					),
				},
				{
					MethodName: "SetValue",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:     handler.SetValue,
							NewRequest: newKeyValueServiceSetValueYARPCRequest,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

type _KeyValueYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_KeyValueYARPCCaller) GetValue(ctx context.Context, request *GetValueRequest, options ...yarpc.CallOption) (*GetValueResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "GetValue", request, newKeyValueServiceGetValueYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*GetValueResponse)
	if !ok {
		return nil, protobuf.CastError(emptyKeyValueServiceGetValueYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_KeyValueYARPCCaller) SetValue(ctx context.Context, request *SetValueRequest, options ...yarpc.CallOption) (*SetValueResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "SetValue", request, newKeyValueServiceSetValueYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*SetValueResponse)
	if !ok {
		return nil, protobuf.CastError(emptyKeyValueServiceSetValueYARPCResponse, responseMessage)
	}
	return response, err
}

type _KeyValueYARPCHandler struct {
	server KeyValueYARPCServer
}

func (h *_KeyValueYARPCHandler) GetValue(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *GetValueRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*GetValueRequest)
		if !ok {
			return nil, protobuf.CastError(emptyKeyValueServiceGetValueYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.GetValue(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func (h *_KeyValueYARPCHandler) SetValue(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *SetValueRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*SetValueRequest)
		if !ok {
			return nil, protobuf.CastError(emptyKeyValueServiceSetValueYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.SetValue(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newKeyValueServiceGetValueYARPCRequest() proto.Message {
	return &GetValueRequest{}
}

func newKeyValueServiceGetValueYARPCResponse() proto.Message {
	return &GetValueResponse{}
}

func newKeyValueServiceSetValueYARPCRequest() proto.Message {
	return &SetValueRequest{}
}

func newKeyValueServiceSetValueYARPCResponse() proto.Message {
	return &SetValueResponse{}
}

var (
	emptyKeyValueServiceGetValueYARPCRequest  = &GetValueRequest{}
	emptyKeyValueServiceGetValueYARPCResponse = &GetValueResponse{}
	emptyKeyValueServiceSetValueYARPCRequest  = &SetValueRequest{}
	emptyKeyValueServiceSetValueYARPCResponse = &SetValueResponse{}
)

// SinkYARPCClient is the YARPC client-side interface for the Sink service.
type SinkYARPCClient interface {
	Fire(context.Context, *FireRequest, ...yarpc.CallOption) (yarpc.Ack, error)
}

// NewSinkYARPCClient builds a new YARPC client for the Sink service.
func NewSinkYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) SinkYARPCClient {
	return &_SinkYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.internal.examples.protobuf.example.Sink",
			ClientConfig: clientConfig,
			Options:      options,
		},
	)}
}

// SinkYARPCServer is the YARPC server-side interface for the Sink service.
type SinkYARPCServer interface {
	Fire(context.Context, *FireRequest) error
}

// BuildSinkYARPCProcedures prepares an implementation of the Sink service for YARPC registration.
func BuildSinkYARPCProcedures(server SinkYARPCServer) []transport.Procedure {
	handler := &_SinkYARPCHandler{server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName:        "uber.yarpc.internal.examples.protobuf.example.Sink",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{
				{
					MethodName: "Fire",
					Handler: protobuf.NewOnewayHandler(
						protobuf.OnewayHandlerParams{
							Handle:     handler.Fire,
							NewRequest: newSinkServiceFireYARPCRequest,
						},
					),
				},
			},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

type _SinkYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_SinkYARPCCaller) Fire(ctx context.Context, request *FireRequest, options ...yarpc.CallOption) (yarpc.Ack, error) {
	return c.streamClient.CallOneway(ctx, "Fire", request, options...)
}

type _SinkYARPCHandler struct {
	server SinkYARPCServer
}

func (h *_SinkYARPCHandler) Fire(ctx context.Context, requestMessage proto.Message) error {
	var request *FireRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*FireRequest)
		if !ok {
			return protobuf.CastError(emptySinkServiceFireYARPCRequest, requestMessage)
		}
	}
	return h.server.Fire(ctx, request)
}

func newSinkServiceFireYARPCRequest() proto.Message {
	return &FireRequest{}
}

func newSinkServiceFireYARPCResponse() proto.Message {
	return &yarpcproto.Oneway{}
}

var (
	emptySinkServiceFireYARPCRequest  = &FireRequest{}
	emptySinkServiceFireYARPCResponse = &yarpcproto.Oneway{}
)

// FooYARPCClient is the YARPC client-side interface for the Foo service.
type FooYARPCClient interface {
	EchoOut(context.Context, ...yarpc.CallOption) (FooServiceEchoOutYARPCClient, error)
	EchoIn(context.Context, *EchoInRequest, ...yarpc.CallOption) (FooServiceEchoInYARPCClient, error)
	EchoBoth(context.Context, ...yarpc.CallOption) (FooServiceEchoBothYARPCClient, error)
}

// FooServiceEchoOutYARPCClient sends EchoOutRequests and receives the single EchoOutResponse when sending is done.
type FooServiceEchoOutYARPCClient interface {
	Context() context.Context
	RequestMeta() *transport.RequestMeta
	ResponseMeta() *transport.ResponseMeta
	Send(*EchoOutRequest) error
	CloseAndRecv() (*EchoOutResponse, error)
}

// FooServiceEchoInYARPCClient receives EchoInResponses, returning io.EOF when the stream is complete.
type FooServiceEchoInYARPCClient interface {
	Context() context.Context
	RequestMeta() *transport.RequestMeta
	ResponseMeta() *transport.ResponseMeta
	Recv() (*EchoInResponse, error)
}

// FooServiceEchoBothYARPCClient sends EchoBothRequests and receives EchoBothResponses, returning io.EOF when the stream is complete.
type FooServiceEchoBothYARPCClient interface {
	Context() context.Context
	RequestMeta() *transport.RequestMeta
	ResponseMeta() *transport.ResponseMeta
	Send(*EchoBothRequest) error
	Recv() (*EchoBothResponse, error)
	CloseSend() error
}

// NewFooYARPCClient builds a new YARPC client for the Foo service.
func NewFooYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) FooYARPCClient {
	return &_FooYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.internal.examples.protobuf.example.Foo",
			ClientConfig: clientConfig,
			Options:      options,
		},
	)}
}

// FooYARPCServer is the YARPC server-side interface for the Foo service.
type FooYARPCServer interface {
	EchoOut(FooServiceEchoOutYARPCServer) (*EchoOutResponse, error)
	EchoIn(*EchoInRequest, FooServiceEchoInYARPCServer) error
	EchoBoth(FooServiceEchoBothYARPCServer) error
}

// FooServiceEchoOutYARPCServer receives EchoOutRequests.
type FooServiceEchoOutYARPCServer interface {
	Context() context.Context
	RequestMeta() *transport.RequestMeta
	SetResponseMeta(*transport.ResponseMeta)
	Recv() (*EchoOutRequest, error)
}

// FooServiceEchoInYARPCServer sends EchoInResponses.
type FooServiceEchoInYARPCServer interface {
	Context() context.Context
	RequestMeta() *transport.RequestMeta
	SetResponseMeta(*transport.ResponseMeta)
	Send(*EchoInResponse) error
}

// FooServiceEchoBothYARPCServer receives EchoBothRequests and sends EchoBothResponse.
type FooServiceEchoBothYARPCServer interface {
	Context() context.Context
	RequestMeta() *transport.RequestMeta
	SetResponseMeta(*transport.ResponseMeta)
	Recv() (*EchoBothRequest, error)
	Send(*EchoBothResponse) error
}

// BuildFooYARPCProcedures prepares an implementation of the Foo service for YARPC registration.
func BuildFooYARPCProcedures(server FooYARPCServer) []transport.Procedure {
	handler := &_FooYARPCHandler{server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName:         "uber.yarpc.internal.examples.protobuf.example.Foo",
			UnaryHandlerParams:  []protobuf.BuildProceduresUnaryHandlerParams{},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{
				{
					MethodName: "EchoBoth",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.EchoBoth,
						},
					),
				},

				{
					MethodName: "EchoIn",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.EchoIn,
						},
					),
				},

				{
					MethodName: "EchoOut",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.EchoOut,
						},
					),
				},
			},
		},
	)
}

type _FooYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_FooYARPCCaller) EchoOut(ctx context.Context, options ...yarpc.CallOption) (FooServiceEchoOutYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "EchoOut", options...)
	if err != nil {
		return nil, err
	}
	return &_FooServiceEchoOutYARPCClient{stream: stream}, nil
}

func (c *_FooYARPCCaller) EchoIn(ctx context.Context, request *EchoInRequest, options ...yarpc.CallOption) (FooServiceEchoInYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "EchoIn", options...)
	if err != nil {
		return nil, err
	}
	reader, closer, err := protobuf.ToReader(request, stream.RequestMeta().Encoding)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return nil, err
	}
	if err := stream.SendMsg(&transport.StreamMessage{ReadCloser: ioutil.NopCloser(reader)}); err != nil {
		return nil, err
	}
	return &_FooServiceEchoInYARPCClient{stream: stream}, nil
}

func (c *_FooYARPCCaller) EchoBoth(ctx context.Context, options ...yarpc.CallOption) (FooServiceEchoBothYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "EchoBoth", options...)
	if err != nil {
		return nil, err
	}
	return &_FooServiceEchoBothYARPCClient{stream: stream}, nil
}

type _FooYARPCHandler struct {
	server FooYARPCServer
}

func (h *_FooYARPCHandler) EchoOut(serverStream transport.ServerStream) error {
	response, err := h.server.EchoOut(&_FooServiceEchoOutYARPCServer{serverStream: serverStream})
	if err != nil {
		return err
	}
	reader, closer, err := protobuf.ToReader(response, serverStream.RequestMeta().Encoding)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return err
	}
	return serverStream.SendMsg(&transport.StreamMessage{ReadCloser: ioutil.NopCloser(reader)})
}

func (h *_FooYARPCHandler) EchoIn(serverStream transport.ServerStream) error {
	src, err := serverStream.RecvMsg()
	if err != nil {
		return err
	}
	requestMessage, err := protobuf.ToProtoMessage(src, serverStream.RequestMeta().Encoding, newFooServiceEchoInYARPCRequest)
	if requestMessage == nil {
		return err
	}
	request, ok := requestMessage.(*EchoInRequest)
	if !ok {
		return protobuf.CastError(emptyFooServiceEchoInYARPCRequest, requestMessage)
	}
	return h.server.EchoIn(request, &_FooServiceEchoInYARPCServer{serverStream: serverStream})
}

func (h *_FooYARPCHandler) EchoBoth(serverStream transport.ServerStream) error {
	return h.server.EchoBoth(&_FooServiceEchoBothYARPCServer{serverStream: serverStream})
}

type _FooServiceEchoOutYARPCClient struct {
	stream transport.ClientStream
}

func (c *_FooServiceEchoOutYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_FooServiceEchoOutYARPCClient) RequestMeta() *transport.RequestMeta {
	return c.stream.RequestMeta()
}

func (c *_FooServiceEchoOutYARPCClient) ResponseMeta() *transport.ResponseMeta {
	return c.stream.ResponseMeta()
}

func (c *_FooServiceEchoOutYARPCClient) Send(request *EchoOutRequest) error {
	reader, closer, err := protobuf.ToReader(request, c.stream.RequestMeta().Encoding)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return err
	}
	return c.stream.SendMsg(&transport.StreamMessage{ReadCloser: ioutil.NopCloser(reader)})
}

func (c *_FooServiceEchoOutYARPCClient) CloseAndRecv() (*EchoOutResponse, error) {
	if err := c.stream.Close(); err != nil {
		return nil, err
	}
	src, err := c.stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	responseMessage, err := protobuf.ToProtoMessage(src, c.stream.RequestMeta().Encoding, newFooServiceEchoOutYARPCResponse)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*EchoOutResponse)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoOutYARPCResponse, responseMessage)
	}
	return response, err
}

type _FooServiceEchoInYARPCClient struct {
	stream transport.ClientStream
}

func (c *_FooServiceEchoInYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_FooServiceEchoInYARPCClient) RequestMeta() *transport.RequestMeta {
	return c.stream.RequestMeta()
}

func (c *_FooServiceEchoInYARPCClient) ResponseMeta() *transport.ResponseMeta {
	return c.stream.ResponseMeta()
}

func (c *_FooServiceEchoInYARPCClient) Recv() (*EchoInResponse, error) {
	src, err := c.stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	responseMessage, err := protobuf.ToProtoMessage(src, c.stream.RequestMeta().Encoding, newFooServiceEchoInYARPCResponse)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*EchoInResponse)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoInYARPCResponse, responseMessage)
	}
	return response, err
}

type _FooServiceEchoBothYARPCClient struct {
	stream transport.ClientStream
}

func (c *_FooServiceEchoBothYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_FooServiceEchoBothYARPCClient) RequestMeta() *transport.RequestMeta {
	return c.stream.RequestMeta()
}

func (c *_FooServiceEchoBothYARPCClient) ResponseMeta() *transport.ResponseMeta {
	return c.stream.ResponseMeta()
}

func (c *_FooServiceEchoBothYARPCClient) Send(request *EchoBothRequest) error {
	reader, closer, err := protobuf.ToReader(request, c.stream.RequestMeta().Encoding)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return err
	}
	return c.stream.SendMsg(&transport.StreamMessage{ReadCloser: ioutil.NopCloser(reader)})
}

func (c *_FooServiceEchoBothYARPCClient) Recv() (*EchoBothResponse, error) {
	src, err := c.stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	responseMessage, err := protobuf.ToProtoMessage(src, c.stream.RequestMeta().Encoding, newFooServiceEchoBothYARPCResponse)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*EchoBothResponse)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoBothYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_FooServiceEchoBothYARPCClient) CloseSend() error {
	return c.stream.Close()
}

type _FooServiceEchoOutYARPCServer struct {
	serverStream transport.ServerStream
}

func (s *_FooServiceEchoOutYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_FooServiceEchoOutYARPCServer) RequestMeta() *transport.RequestMeta {
	return s.serverStream.RequestMeta()
}

func (s *_FooServiceEchoOutYARPCServer) SetResponseMeta(responseMeta *transport.ResponseMeta) {
	s.serverStream.SetResponseMeta(responseMeta)
}

func (s *_FooServiceEchoOutYARPCServer) Recv() (*EchoOutRequest, error) {
	src, err := s.serverStream.RecvMsg()
	if err != nil {
		return nil, err
	}
	responseMessage, err := protobuf.ToProtoMessage(src, s.serverStream.RequestMeta().Encoding, newFooServiceEchoOutYARPCRequest)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*EchoOutRequest)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoOutYARPCRequest, responseMessage)
	}
	return response, err
}

type _FooServiceEchoInYARPCServer struct {
	serverStream transport.ServerStream
}

func (s *_FooServiceEchoInYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_FooServiceEchoInYARPCServer) RequestMeta() *transport.RequestMeta {
	return s.serverStream.RequestMeta()
}

func (s *_FooServiceEchoInYARPCServer) SetResponseMeta(responseMeta *transport.ResponseMeta) {
	s.serverStream.SetResponseMeta(responseMeta)
}

func (s *_FooServiceEchoInYARPCServer) Send(response *EchoInResponse) error {
	reader, closer, err := protobuf.ToReader(response, s.serverStream.RequestMeta().Encoding)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return err
	}
	return s.serverStream.SendMsg(&transport.StreamMessage{ReadCloser: ioutil.NopCloser(reader)})
}

type _FooServiceEchoBothYARPCServer struct {
	serverStream transport.ServerStream
}

func (s *_FooServiceEchoBothYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_FooServiceEchoBothYARPCServer) RequestMeta() *transport.RequestMeta {
	return s.serverStream.RequestMeta()
}

func (s *_FooServiceEchoBothYARPCServer) SetResponseMeta(responseMeta *transport.ResponseMeta) {
	s.serverStream.SetResponseMeta(responseMeta)
}

func (s *_FooServiceEchoBothYARPCServer) Recv() (*EchoBothRequest, error) {
	src, err := s.serverStream.RecvMsg()
	if err != nil {
		return nil, err
	}
	responseMessage, err := protobuf.ToProtoMessage(src, s.serverStream.RequestMeta().Encoding, newFooServiceEchoBothYARPCRequest)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*EchoBothRequest)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoBothYARPCRequest, responseMessage)
	}
	return response, err
}

func (s *_FooServiceEchoBothYARPCServer) Send(response *EchoBothResponse) error {
	reader, closer, err := protobuf.ToReader(response, s.serverStream.RequestMeta().Encoding)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return err
	}
	return s.serverStream.SendMsg(&transport.StreamMessage{ReadCloser: ioutil.NopCloser(reader)})
}

func newFooServiceEchoOutYARPCRequest() proto.Message {
	return &EchoOutRequest{}
}

func newFooServiceEchoOutYARPCResponse() proto.Message {
	return &EchoOutResponse{}
}

func newFooServiceEchoInYARPCRequest() proto.Message {
	return &EchoInRequest{}
}

func newFooServiceEchoInYARPCResponse() proto.Message {
	return &EchoInResponse{}
}

func newFooServiceEchoBothYARPCRequest() proto.Message {
	return &EchoBothRequest{}
}

func newFooServiceEchoBothYARPCResponse() proto.Message {
	return &EchoBothResponse{}
}

var (
	emptyFooServiceEchoOutYARPCRequest   = &EchoOutRequest{}
	emptyFooServiceEchoOutYARPCResponse  = &EchoOutResponse{}
	emptyFooServiceEchoInYARPCRequest    = &EchoInRequest{}
	emptyFooServiceEchoInYARPCResponse   = &EchoInResponse{}
	emptyFooServiceEchoBothYARPCRequest  = &EchoBothRequest{}
	emptyFooServiceEchoBothYARPCResponse = &EchoBothResponse{}
)

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) KeyValueYARPCClient {
			return NewKeyValueYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) SinkYARPCClient {
			return NewSinkYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) FooYARPCClient {
			return NewFooYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
