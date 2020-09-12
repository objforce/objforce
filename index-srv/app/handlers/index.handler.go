package handlers

import "github.com/objforce/objforce/index-srv/app/domain/services"

type IndexHandler struct {
	indexService services.IndexService
}

func NewIndexHandler(indexService services.IndexService) *IndexHandler {
	return &IndexHandler{
		indexService,
	}
}
