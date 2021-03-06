// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/meta/meta.proto

package meta

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CustomObjectService service

func NewCustomObjectServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CustomObjectService service

type CustomObjectService interface {
	Create(ctx context.Context, in *CreateCustomObjectRequest, opts ...client.CallOption) (*CustomObject, error)
	Retrieve(ctx context.Context, in *GetCustomObjectRequest, opts ...client.CallOption) (*CustomObject, error)
	Update(ctx context.Context, in *CustomObject, opts ...client.CallOption) (*CustomObject, error)
	Delete(ctx context.Context, in *DeleteCustomObjectRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	FindByObjName(ctx context.Context, in *FindByObjNameRequest, opts ...client.CallOption) (*CustomObject, error)
}

type customObjectService struct {
	c    client.Client
	name string
}

func NewCustomObjectService(name string, c client.Client) CustomObjectService {
	return &customObjectService{
		c:    c,
		name: name,
	}
}

func (c *customObjectService) Create(ctx context.Context, in *CreateCustomObjectRequest, opts ...client.CallOption) (*CustomObject, error) {
	req := c.c.NewRequest(c.name, "CustomObjectService.Create", in)
	out := new(CustomObject)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customObjectService) Retrieve(ctx context.Context, in *GetCustomObjectRequest, opts ...client.CallOption) (*CustomObject, error) {
	req := c.c.NewRequest(c.name, "CustomObjectService.Retrieve", in)
	out := new(CustomObject)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customObjectService) Update(ctx context.Context, in *CustomObject, opts ...client.CallOption) (*CustomObject, error) {
	req := c.c.NewRequest(c.name, "CustomObjectService.Update", in)
	out := new(CustomObject)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customObjectService) Delete(ctx context.Context, in *DeleteCustomObjectRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "CustomObjectService.Delete", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customObjectService) FindByObjName(ctx context.Context, in *FindByObjNameRequest, opts ...client.CallOption) (*CustomObject, error) {
	req := c.c.NewRequest(c.name, "CustomObjectService.FindByObjName", in)
	out := new(CustomObject)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomObjectService service

type CustomObjectServiceHandler interface {
	Create(context.Context, *CreateCustomObjectRequest, *CustomObject) error
	Retrieve(context.Context, *GetCustomObjectRequest, *CustomObject) error
	Update(context.Context, *CustomObject, *CustomObject) error
	Delete(context.Context, *DeleteCustomObjectRequest, *emptypb.Empty) error
	FindByObjName(context.Context, *FindByObjNameRequest, *CustomObject) error
}

func RegisterCustomObjectServiceHandler(s server.Server, hdlr CustomObjectServiceHandler, opts ...server.HandlerOption) error {
	type customObjectService interface {
		Create(ctx context.Context, in *CreateCustomObjectRequest, out *CustomObject) error
		Retrieve(ctx context.Context, in *GetCustomObjectRequest, out *CustomObject) error
		Update(ctx context.Context, in *CustomObject, out *CustomObject) error
		Delete(ctx context.Context, in *DeleteCustomObjectRequest, out *emptypb.Empty) error
		FindByObjName(ctx context.Context, in *FindByObjNameRequest, out *CustomObject) error
	}
	type CustomObjectService struct {
		customObjectService
	}
	h := &customObjectServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CustomObjectService{h}, opts...))
}

type customObjectServiceHandler struct {
	CustomObjectServiceHandler
}

func (h *customObjectServiceHandler) Create(ctx context.Context, in *CreateCustomObjectRequest, out *CustomObject) error {
	return h.CustomObjectServiceHandler.Create(ctx, in, out)
}

func (h *customObjectServiceHandler) Retrieve(ctx context.Context, in *GetCustomObjectRequest, out *CustomObject) error {
	return h.CustomObjectServiceHandler.Retrieve(ctx, in, out)
}

func (h *customObjectServiceHandler) Update(ctx context.Context, in *CustomObject, out *CustomObject) error {
	return h.CustomObjectServiceHandler.Update(ctx, in, out)
}

func (h *customObjectServiceHandler) Delete(ctx context.Context, in *DeleteCustomObjectRequest, out *emptypb.Empty) error {
	return h.CustomObjectServiceHandler.Delete(ctx, in, out)
}

func (h *customObjectServiceHandler) FindByObjName(ctx context.Context, in *FindByObjNameRequest, out *CustomObject) error {
	return h.CustomObjectServiceHandler.FindByObjName(ctx, in, out)
}

// Api Endpoints for CustomFieldService service

func NewCustomFieldServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CustomFieldService service

type CustomFieldService interface {
	Create(ctx context.Context, in *CustomField, opts ...client.CallOption) (*CustomField, error)
	Update(ctx context.Context, in *CustomField, opts ...client.CallOption) (*CustomField, error)
	Get(ctx context.Context, in *GetCustomFieldRequest, opts ...client.CallOption) (*CustomField, error)
	Delete(ctx context.Context, in *DeleteCustomFieldRequest, opts ...client.CallOption) (*CustomField, error)
}

type customFieldService struct {
	c    client.Client
	name string
}

func NewCustomFieldService(name string, c client.Client) CustomFieldService {
	return &customFieldService{
		c:    c,
		name: name,
	}
}

func (c *customFieldService) Create(ctx context.Context, in *CustomField, opts ...client.CallOption) (*CustomField, error) {
	req := c.c.NewRequest(c.name, "CustomFieldService.Create", in)
	out := new(CustomField)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customFieldService) Update(ctx context.Context, in *CustomField, opts ...client.CallOption) (*CustomField, error) {
	req := c.c.NewRequest(c.name, "CustomFieldService.Update", in)
	out := new(CustomField)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customFieldService) Get(ctx context.Context, in *GetCustomFieldRequest, opts ...client.CallOption) (*CustomField, error) {
	req := c.c.NewRequest(c.name, "CustomFieldService.Get", in)
	out := new(CustomField)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customFieldService) Delete(ctx context.Context, in *DeleteCustomFieldRequest, opts ...client.CallOption) (*CustomField, error) {
	req := c.c.NewRequest(c.name, "CustomFieldService.Delete", in)
	out := new(CustomField)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomFieldService service

type CustomFieldServiceHandler interface {
	Create(context.Context, *CustomField, *CustomField) error
	Update(context.Context, *CustomField, *CustomField) error
	Get(context.Context, *GetCustomFieldRequest, *CustomField) error
	Delete(context.Context, *DeleteCustomFieldRequest, *CustomField) error
}

func RegisterCustomFieldServiceHandler(s server.Server, hdlr CustomFieldServiceHandler, opts ...server.HandlerOption) error {
	type customFieldService interface {
		Create(ctx context.Context, in *CustomField, out *CustomField) error
		Update(ctx context.Context, in *CustomField, out *CustomField) error
		Get(ctx context.Context, in *GetCustomFieldRequest, out *CustomField) error
		Delete(ctx context.Context, in *DeleteCustomFieldRequest, out *CustomField) error
	}
	type CustomFieldService struct {
		customFieldService
	}
	h := &customFieldServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CustomFieldService{h}, opts...))
}

type customFieldServiceHandler struct {
	CustomFieldServiceHandler
}

func (h *customFieldServiceHandler) Create(ctx context.Context, in *CustomField, out *CustomField) error {
	return h.CustomFieldServiceHandler.Create(ctx, in, out)
}

func (h *customFieldServiceHandler) Update(ctx context.Context, in *CustomField, out *CustomField) error {
	return h.CustomFieldServiceHandler.Update(ctx, in, out)
}

func (h *customFieldServiceHandler) Get(ctx context.Context, in *GetCustomFieldRequest, out *CustomField) error {
	return h.CustomFieldServiceHandler.Get(ctx, in, out)
}

func (h *customFieldServiceHandler) Delete(ctx context.Context, in *DeleteCustomFieldRequest, out *CustomField) error {
	return h.CustomFieldServiceHandler.Delete(ctx, in, out)
}
