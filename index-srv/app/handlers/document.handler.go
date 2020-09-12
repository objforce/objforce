package handlers

import (
	"context"
	"github.com/objforce/objforce/index-srv/app/domain/services"
	"github.com/objforce/objforce/index-srv/app/dtos"
	index "github.com/objforce/objforce/index-srv/proto/index/gen-go"
	"github.com/xxxmicro/base/mapper"
)

type DocumentHandler struct {
	documentService services.DocumentService
}

func NewDocumentHandler(documentService services.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		documentService,
	}
}

func (h *DocumentHandler) Create(c context.Context, req *index.Document, rsp *index.Document) error {
	dto := &dtos.Document{}
	mapper.Map(req, dto)

	err := h.documentService.Upsert(c, dto)
	if err != nil {
		return err
	}

	return nil
}

func (h *DocumentHandler) Update(c context.Context, req *index.Document, rsp *index.Document) error {
	dto := &dtos.Document{}
	mapper.Map(req, dto)
	err := h.documentService.Update(c, dto)
	if err != nil {
		return err
	}

	return nil
}

func (h *DocumentHandler) Retrieve(c context.Context, req *index.FindDocumentRequest, rsp *index.Document) error {
	doc, err := h.documentService.Retrieve(c, req.ObjId, req.Id)
	if err != nil {
		return err
	}

	mapper.Map(doc, rsp)

	return nil
}

func (h *DocumentHandler) Delete(c context.Context, req *index.DeleteDocumentRequest, rsp *index.Document) error {
	err := h.documentService.Delete(c, req.ObjId, req.Id)
	if err != nil {
		return err
	}

	return nil
}
