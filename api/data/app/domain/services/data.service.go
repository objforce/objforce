package services

import (
	"context"

	"github.com/goinggo/mapstructure"
	"github.com/objforce/objforce/proto/data"
	"github.com/objforce/objforce/service/data/app/domain/models"
)

type DataService interface {
	Create(c context.Context, object *models.SObject) (*models.SObject, error)
	Update(c context.Context, objct *models.SObject) (*models.SObject, error)
	Upsert(c context.Context, object *models.SObject) (*models.UpsertResult, error)
	Get(c context.Context, req *models.GetSObjectRequest) (*models.SObject, error)
	Delete(c context.Context, req *models.DeleteSObjectRequest) error
}

type dataService struct {
	objectService data.SObjectService
}

func NewDataService(objectService data.SObjectService) DataService {
	return &dataService{
		objectService,
	}
}

func (s *dataService) Create(c context.Context, object *models.SObject) (*models.SObject, error) {
	orgId := c.Value("orgId").(string)

	pbObject := &data.SObject{
		OrgId: orgId,
	}
	mapstructure.Decode(object, pbObject)

	pb := &data.CreateSObjectRequest{Object: pbObject}

	pbNewObject, err := s.objectService.Create(c, pb)
	if err != nil {
		return nil, err
	}

	newObject := &models.SObject{}
	mapstructure.Decode(pbNewObject, newObject)

	return newObject, nil
}

func (s *dataService) Update(c context.Context, object *models.SObject) (*models.SObject, error) {
	orgId := c.Value("orgId").(string)

	pbObject := &data.SObject{
		OrgId: orgId,
	}
	mapstructure.Decode(object, pbObject)

	pbUpdatedObject, err := s.objectService.Update(c, &data.UpdateSObjectRequest{Object: pbObject})
	if err != nil {
		return nil, err
	}

	updatedObject := &models.SObject{}
	mapstructure.Decode(pbUpdatedObject, updatedObject)
	return updatedObject, nil
}

func (s *dataService) Upsert(c context.Context, object *models.SObject) (*models.UpsertSObjectResult, error) {
	orgId := c.Value("orgId").(string)

	pbObject := &data.SObject{
		OrgId: orgId,
	}
	mapstructure.Decode(object, pbObject)

	pbRsp, err := s.objectService.Upsert(c, &data.UpsertSObjectRequest{Object: pbObject})
	if err != nil {
		return nil, err
	}

	upsertResult := &models.UpsertResult{}
	mapstructure.Decode(pbRsp, upsertResult)

	return upsertResult, nil
}

func (s *dataService) Get(c context.Context, req *models.GetSObjectRequest) (*models.SObject, error) {
	orgId := c.Value("orgId").(string)

	pbReq := &data.GetSObjectRequest{}
	pbReq.OrgId = orgId
	mapstructure.Decode(req, pbReq)

	pbObject, err := s.objectService.Get(c, pbReq)
	if err != nil {
		return nil, err
	}

	var object *models.SObject
	mapstructure.Decode(pbObject, object)
	return object, nil
}

func (s *dataService) Delete(c context.Context, req *models.DeleteSObjectRequest) error {
	orgId := c.Value("orgId").(string)

	_, err := s.objectService.Delete(c, &data.DeleteSObjectRequest{OrgId: orgId, Type: req.Type, Id: req.Id})
	if err != nil {
		return err
	}

	return nil
}
