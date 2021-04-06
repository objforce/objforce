package services

import (
	"context"

	"github.com/goinggo/mapstructure"
	"github.com/micro/go-micro/v2/client"
	"github.com/objforce/objforce/proto/meta"
	"github.com/objforce/objforce/service/meta/app/domain/entities"
	"github.com/objforce/objforce/service/meta/app/domain/repositories"
	"github.com/objforce/objforce/service/meta/app/models"
	"github.com/objforce/objforce/service/meta/config"
)

type CustomObjectService interface {
	Create(c context.Context, model *models.CustomObject) (*models.CustomObject, error)
	Upsert(c context.Context, model *models.CustomObject) (*models.CustomObject, error)
	Update(c context.Context, model *models.CustomObject) (*models.CustomObject, error)
	Get(c context.Context, id string) (*models.CustomObject, error)
	Delete(c context.Context, id string) error
}

type customObjectService struct {
	client                 client.Client
	customObjectRepository repositories.CustomObjectRepository
}

func NewCustomObjectService(client client.Client, customObjectRepository repositories.CustomObjectRepository) CustomObjectService {
	s := &customObjectService{
		client,
		customObjectRepository,
	}

	s1 := CustomObjectServiceEventWrapper{}
	s1.Wrap(s)
	return s
}

func (s *customObjectService) Create(c context.Context, model *models.CustomObject) (*models.CustomObject, error) {
	entity := &entities.MTObject{}
	mapstructure.Decode(model, entity)

	err := s.customObjectRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapstructure.Decode(entity, model)

	return model, nil
}

func (s *customObjectService) Upsert(c context.Context, model *models.CustomObject) (*models.CustomObject, error) {
	entity := &entities.MTObject{ObjId: model.ObjId}

	mapstructure.Decode(model, entity)

	if err := s.customObjectRepository.Upsert(c, entity); err != nil {
		return nil, err
	}

	mapstructure.Decode(entity, model)

	return model, nil
}

func (s *customObjectService) Update(c context.Context, model *models.CustomObject) (*models.CustomObject, error) {
	entity := &entities.MTObject{ObjId: model.ObjId}

	mapstructure.Decode(model, entity)

	if err := s.customObjectRepository.Update(c, entity); err != nil {
		return nil, err
	}

	entity, err := s.customObjectRepository.Get(c, entity.ObjId)
	if err != nil {
		return nil, err
	}

	mapstructure.Decode(entity, model)

	return model, nil
}

func (s *customObjectService) Get(c context.Context, objId string) (*models.CustomObject, error) {
	entity, err := s.customObjectRepository.Get(c, objId)
	if err != nil {
		return nil, err
	}

	model := &models.CustomObject{}
	mapstructure.Decode(entity, model)
	return model, nil
}

func (s *customObjectService) Delete(c context.Context, objId string) error {
	return s.customObjectRepository.Delete(c, objId)
}

type CustomObjectServiceEventWrapper struct {
	ref CustomObjectService
}

func (s *CustomObjectServiceEventWrapper) Wrap(ref CustomObjectService) {
	s.ref = ref
}

func (s *CustomObjectServiceEventWrapper) Create(c context.Context, model *models.CustomObject) (*models.CustomObject, error) {
	newModel, err := s.ref.Create(c, model)
	if err != nil {
		return nil, err
	}

	pb := &meta.CustomObject{}
	mapstructure.Decode(newModel, pb)
	err = client.Publish(c, client.NewMessage(config.CustomObjectCreatedTopic, pb))
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

func (s *CustomObjectServiceEventWrapper) Upsert(c context.Context, model *models.CustomObject) (*models.CustomObject, error) {
	newModel, err := s.ref.Upsert(c, model)
	if err != nil {
		return nil, err
	}

	pb := &meta.CustomObject{}
	mapstructure.Decode(newModel, pb)
	err = client.Publish(c, client.NewMessage(config.CustomObjectUpsertedTopic, pb))
	if err != nil {
		return newModel, err
	}

	return newModel, err
}

func (s *CustomObjectServiceEventWrapper) Update(c context.Context, model *models.CustomObject) (*models.CustomObject, error) {
	updatedModel, err := s.ref.Update(c, model)
	if err != nil {
		return nil, err
	}

	pb := &meta.CustomObject{}
	mapstructure.Decode(updatedModel, pb)
	err = client.Publish(c, client.NewMessage(config.CustomObjectUpdatedTopic, pb))
	if err != nil {
		return updatedModel, err
	}
	return updatedModel, err
}

func (s *CustomObjectServiceEventWrapper) Get(c context.Context, id string) (*models.CustomObject, error) {
	return s.ref.Get(c, id)
}
func (s *CustomObjectServiceEventWrapper) Delete(c context.Context, id string) error {
	model, err := s.ref.Get(c, id)
	if err != nil {
		return err
	}

	err = s.Delete(c, id)
	if err != nil {
		return err
	}

	pb := &meta.CustomObject{}
	mapstructure.Decode(model, pb)
	err = client.Publish(c, client.NewMessage(config.CustomObjectDeletedTopic, pb))
	if err != nil {
		return err
	}

	return nil
}
