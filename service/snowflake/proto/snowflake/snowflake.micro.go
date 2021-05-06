// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/snowflake/snowflake.proto

package go_micro_srv_snowflake

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Snowflake service

type SnowflakeService interface {
	GenerateOnlyId(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type snowflakeService struct {
	c    client.Client
	name string
}

func NewSnowflakeService(name string, c client.Client) SnowflakeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.snowflake"
	}
	return &snowflakeService{
		c:    c,
		name: name,
	}
}

func (c *snowflakeService) GenerateOnlyId(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Snowflake.GenerateOnlyId", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Snowflake service

type SnowflakeHandler interface {
	GenerateOnlyId(context.Context, *Request, *Response) error
}

func RegisterSnowflakeHandler(s server.Server, hdlr SnowflakeHandler, opts ...server.HandlerOption) error {
	type snowflake interface {
		GenerateOnlyId(ctx context.Context, in *Request, out *Response) error
	}
	type Snowflake struct {
		snowflake
	}
	h := &snowflakeHandler{hdlr}
	return s.Handle(s.NewHandler(&Snowflake{h}, opts...))
}

type snowflakeHandler struct {
	SnowflakeHandler
}

func (h *snowflakeHandler) GenerateOnlyId(ctx context.Context, in *Request, out *Response) error {
	return h.SnowflakeHandler.GenerateOnlyId(ctx, in, out)
}
