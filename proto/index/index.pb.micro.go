// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/index/index.proto

package index

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
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

// Api Endpoints for DocumentService service

func NewDocumentServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DocumentService service

type DocumentService interface {
	Create(ctx context.Context, in *Document, opts ...client.CallOption) (*Document, error)
	Update(ctx context.Context, in *Document, opts ...client.CallOption) (*Document, error)
	Retrieve(ctx context.Context, in *FindDocumentRequest, opts ...client.CallOption) (*Document, error)
	Delete(ctx context.Context, in *DeleteDocumentRequest, opts ...client.CallOption) (*Document, error)
}

type documentService struct {
	c    client.Client
	name string
}

func NewDocumentService(name string, c client.Client) DocumentService {
	return &documentService{
		c:    c,
		name: name,
	}
}

func (c *documentService) Create(ctx context.Context, in *Document, opts ...client.CallOption) (*Document, error) {
	req := c.c.NewRequest(c.name, "DocumentService.Create", in)
	out := new(Document)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentService) Update(ctx context.Context, in *Document, opts ...client.CallOption) (*Document, error) {
	req := c.c.NewRequest(c.name, "DocumentService.Update", in)
	out := new(Document)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentService) Retrieve(ctx context.Context, in *FindDocumentRequest, opts ...client.CallOption) (*Document, error) {
	req := c.c.NewRequest(c.name, "DocumentService.Retrieve", in)
	out := new(Document)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentService) Delete(ctx context.Context, in *DeleteDocumentRequest, opts ...client.CallOption) (*Document, error) {
	req := c.c.NewRequest(c.name, "DocumentService.Delete", in)
	out := new(Document)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DocumentService service

type DocumentServiceHandler interface {
	Create(context.Context, *Document, *Document) error
	Update(context.Context, *Document, *Document) error
	Retrieve(context.Context, *FindDocumentRequest, *Document) error
	Delete(context.Context, *DeleteDocumentRequest, *Document) error
}

func RegisterDocumentServiceHandler(s server.Server, hdlr DocumentServiceHandler, opts ...server.HandlerOption) error {
	type documentService interface {
		Create(ctx context.Context, in *Document, out *Document) error
		Update(ctx context.Context, in *Document, out *Document) error
		Retrieve(ctx context.Context, in *FindDocumentRequest, out *Document) error
		Delete(ctx context.Context, in *DeleteDocumentRequest, out *Document) error
	}
	type DocumentService struct {
		documentService
	}
	h := &documentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DocumentService{h}, opts...))
}

type documentServiceHandler struct {
	DocumentServiceHandler
}

func (h *documentServiceHandler) Create(ctx context.Context, in *Document, out *Document) error {
	return h.DocumentServiceHandler.Create(ctx, in, out)
}

func (h *documentServiceHandler) Update(ctx context.Context, in *Document, out *Document) error {
	return h.DocumentServiceHandler.Update(ctx, in, out)
}

func (h *documentServiceHandler) Retrieve(ctx context.Context, in *FindDocumentRequest, out *Document) error {
	return h.DocumentServiceHandler.Retrieve(ctx, in, out)
}

func (h *documentServiceHandler) Delete(ctx context.Context, in *DeleteDocumentRequest, out *Document) error {
	return h.DocumentServiceHandler.Delete(ctx, in, out)
}
