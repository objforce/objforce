package handlers

import "github.com/objforce/objforce/service/index/app/domain/services"

type IndexHandler struct {
	indexService services.IndexService
}

func NewIndexHandler(indexService services.IndexService) *IndexHandler {
	return &IndexHandler{
		indexService,
	}
}
