package services

import (
	"context"

	"github.com/goinggo/mapstructure"
	"github.com/objforce/objforce/service/meta/domain/entities"
	"github.com/objforce/objforce/service/meta/domain/repositories"
	"github.com/objforce/objforce/service/meta/models"
)

type RelationshipService interface {
	Create(c context.Context, customField *models.Relationship) (*models.Relationship, error)
	Update(c context.Context, customField *models.Relationship) (*models.Relationship, error)
	Get(c context.Context, id string) (*models.Relationship, error)
	Delete(c context.Context, id string) error
}

type relationshipService struct {
	relationshipRepository repositories.RelationshipRepository
}

func NewRelationshipService(relationshipRepository repositories.RelationshipRepository) RelationshipService {
	return &relationshipService{
		relationshipRepository,
	}
}

func (s *relationshipService) Create(c context.Context, model *models.Relationship) (*models.Relationship, error) {
	entity := &entities.MTRelationship{}
	mapstructure.Decode(model, &entity)

	err := s.relationshipRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapstructure.Decode(entity, model)

	return model, nil
}

func (s *relationshipService) Update(c context.Context, model *models.Relationship) (*models.Relationship, error) {
	entity := &entities.MTRelationship{RelationId: model.RelationId}

	mapstructure.Decode(model, &entity)

	err := s.relationshipRepository.Update(c, entity)
	if err != nil {
		return nil, err
	}

	entity, err = s.relationshipRepository.Get(c, entity.RelationId)
	if err != nil {
		return nil, err
	}

	mapstructure.Decode(entity, &model)

	return model, nil
}

func (s *relationshipService) Get(c context.Context, id string) (*models.Relationship, error) {
	entity := &entities.MTRelationship{RelationId: id}
	entity, err := s.relationshipRepository.Get(c, id)
	if err != nil {
		return nil, err
	}

	model := &models.Relationship{}
	mapstructure.Decode(entity, model)
	return model, nil
}

func (s *relationshipService) Delete(c context.Context, id string) error {
	return s.relationshipRepository.Delete(c, id)
}
