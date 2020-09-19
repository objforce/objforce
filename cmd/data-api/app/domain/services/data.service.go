package services

import(
	"context"
	"github.com/objforce/objforce/cmd/data-api/app/dtos"
	data "github.com/objforce/objforce/idl/data/gen-go"
	"github.com/xxxmicro/base/mapper"
)

type DataService interface {
	Create(c context.Context, dtoObjects []*dtos.SObject) (*dtos.CreateSObjectResponse, error)
	Update(c context.Context, dto *dtos.UpdateSObjectRequest) (*dtos.UpdateSObjectResponse, error)
	Upsert(c context.Context, dto *dtos.UpsertSObjectRequest) (*dtos.UpsertSObjectResponse, error)
	Retrieve(c context.Context, dtoReq *dtos.RetrieveSObjectRequest) (*dtos.RetrieveSObjectResponse, error)
	Delete(c context.Context, dtoReq *dtos.DeleteSObjectRequest) (*dtos.DeleteSObjectResponse, error)
}

type dataService struct {
	objectService data.SObjectService
}

func NewDataService(objectService data.SObjectService) DataService {
	return &dataService{
		objectService,
	}
}

func (s *dataService) Create(c context.Context, dtoObjects []*dtos.SObject) (*dtos.CreateSObjectResponse, error) {
	orgId := c.Value("orgId").(string)

	pbObjects := make([]*data.SObject, len(dtoObjects))
	mapper.Map(dtoObjects, pbObjects)
	for _, pbObject := range pbObjects {
		pbObject.OrgId = orgId
	}

	pb := &data.CreateSObjectRequest{Objects: pbObjects}

	pbRsp, err := s.objectService.Create(c, pb)
	if err != nil {
		return nil, err
	}

	dtoRsp := &dtos.CreateSObjectResponse{}
	mapper.Map(pbRsp, dtoRsp)

	return dtoRsp, nil
}

func (s *dataService) Update(c context.Context, dtoReq *dtos.UpdateSObjectRequest) (*dtos.UpdateSObjectResponse, error) {
	orgId := c.Value("orgId").(string)

	pbReq := &data.UpdateSObjectRequest{}
	mapper.Map(dtoReq, pbReq)
	for _, pbObject := range pbReq.Objects {
		pbObject.OrgId = orgId
	}

	pbRsp, err := s.objectService.Update(c, pbReq)
	if err != nil {
		return nil, err
	}

	dtoRsp := &dtos.UpdateSObjectResponse{}
	mapper.Map(pbRsp, dtoRsp)
	return dtoRsp, nil
}

func (s *dataService) Upsert(c context.Context, dtoReq *dtos.UpsertSObjectRequest) (*dtos.UpsertSObjectResponse, error) {
	orgId := c.Value("orgId").(string)

	pbReq := &data.UpsertSObjectRequest{}
	mapper.Map(dtoReq, pbReq)
	for _, pbObject := range pbReq.Objects {
		pbObject.OrgId = orgId
	}

	pbRsp, err := s.objectService.Upsert(c, pbReq)
	if err != nil {
		return nil, err
	}

	dtoRsp := &dtos.UpsertSObjectResponse{}
	mapper.Map(pbRsp, dtoRsp)
	return dtoRsp, nil
}


func (s *dataService) Retrieve(c context.Context, dtoReq *dtos.RetrieveSObjectRequest) (*dtos.RetrieveSObjectResponse, error) {
	orgId := c.Value("orgId").(string)

	pbReq := &data.RetrieveSObjectRequest{}
	mapper.Map(dtoReq, pbReq)
	pbReq.OrgId = orgId

	pbRsp, err := s.objectService.Retrieve(c, pbReq)
	if err != nil {
		return nil, err
	}

	if pbRsp.Results == nil {
		return nil, nil
	}

	dtoRsp := &dtos.RetrieveSObjectResponse{}
	mapper.Map(pbRsp, dtoRsp)
	return dtoRsp, nil
}

func (s *dataService) Delete(c context.Context, dtoReq *dtos.DeleteSObjectRequest) (*dtos.DeleteSObjectResponse, error) {
	orgId := c.Value("orgId").(string)

	pbRsp, err := s.objectService.Delete(c, &data.DeleteSObjectRequest{OrgId: orgId, Type: dtoReq.Type, Ids: dtoReq.Ids})
	if err != nil {
		return nil, err
	}

	dtoRsp := &dtos.DeleteSObjectResponse{}
	mapper.Map(dtoRsp, pbRsp)
	return dtoRsp, nil
}