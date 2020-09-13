package services

import(
	"context"
	"github.com/objforce/objforce/data-srv/app/domain/models"
	"github.com/objforce/objforce/data-srv/app/domain/repositories"
	"github.com/objforce/objforce/data-srv/app/dtos"
	"github.com/xxxmicro/base/mapper"
)

type DataService interface {
	Create(c context.Context, dto *dtos.SObject) (*dtos.SObject, error)
	Update(c context.Context, dto *dtos.SObject) (*dtos.SObject, error)
	FindOne(c context.Context, id string) (*dtos.SObject, error)
	Delete(c context.Context, id string) error
}

type dataService struct {
	dataRepository repositories.DataRepository
}

func NewDataService(dataRepository repositories.DataRepository) DataService {
	return &dataService{
		dataRepository,
	}
}

func (s *dataService) Create(c context.Context, dto *dtos.SObject) (*dtos.SObject, error) {
	entity := &models.MTBasicData{}
	mapper.Map(dto, entity)

	err := s.dataRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *dataService) Update(c context.Context, dto *dtos.SObject) (*dtos.SObject, error) {
	entity := &models.MTBasicData{ObjId: dto.ObjId}

	mapper.Map(dto, entity)

	err := s.dataRepository.Update(c, entity, entity)
	if err != nil {
		return nil, err
	}

	err = s.dataRepository.FindOne(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *dataService) FindOne(c context.Context, id string) (*dtos.SObject, error) {
	entity := &models.MTBasicData{GUID: id}
	s.dataRepository.FindOne(c, entity)

	dto := &dtos.SObject{}
	mapper.Map(entity, dto)
	return dto, nil
}

func (s *dataService) Delete(c context.Context, id string) error {
	return s.dataRepository.Delete(c, &models.MTBasicData{GUID: id})
}