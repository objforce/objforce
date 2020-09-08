package services

import(
	"context"
	"github.com/objforce/objforce/index-server/app/domain/models"
	"github.com/objforce/objforce/index-server/app/domain/repositories"
	"github.com/objforce/objforce/index-server/app/dtos"
	"github.com/xxxmicro/base/mapper"
)

type IndexService interface {
	Create(c context.Context, dto *dtos.Index) (*dtos.Index, error)
	Update(c context.Context, dto *dtos.Index) (*dtos.Index, error)
	FindOne(c context.Context, id string) (*dtos.Index, error)
	Delete(c context.Context, id string) error
}

type indexService struct {
	indexRepository repositories.IndexRepository
}

func NewIndexService(indexRepository repositories.IndexRepository) IndexService {
	return &indexService{
		indexRepository,
	}
}

func (s *indexService) Create(c context.Context, dto *dtos.Index) (*dtos.Index, error) {
	entity := &models.Index{}
	mapper.Map(dto, entity)

	err := s.indexRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *indexService) Update(c context.Context, dto *dtos.Index) (*dtos.Index, error) {
	entity := &models.Index{ObjId: dto.ObjId}

	mapper.Map(dto, entity)

	err := s.indexRepository.Update(c, entity, entity)
	if err != nil {
		return nil, err
	}

	err = s.indexRepository.FindOne(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *indexService) FindOne(c context.Context, id string) (*dtos.Index, error) {
	entity := &models.Index{ObjId: id}
	s.indexRepository.FindOne(c, entity)

	dto := &dtos.Index{}
	mapper.Map(entity, dto)
	return dto, nil
}

func (s *indexService) Delete(c context.Context, id string) error {
	return s.indexRepository.Delete(c, &models.Index{ObjId: id})
}