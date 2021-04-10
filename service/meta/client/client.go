package client

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	"github.com/objforce/objforce/proto/meta"
	proto "github.com/objforce/objforce/proto/meta"
	"github.com/objforce/objforce/service/meta/models"
	"github.com/thoas/go-funk"
)

type CustomObjectClient interface {
	Create(c context.Context, req *models.CreateCustomObjectRequest) (*models.CustomObject, error)
	Retrieve(c context.Context, objId string) (*models.CustomObject, error)
	Update(c context.Context, model *models.CustomObject) (*models.CustomObject, error)
	Delete(c context.Context, objId string) error
	FindByObjName(c context.Context, orgId, objName string) (*models.CustomObject, error)
}

type srv struct {
	client meta.CustomObjectService
}

func NewClient(addr string, client client.Client) CustomObjectClient {
	return &srv{
		client: meta.NewCustomObjectService(addr, client),
	}
}

func (m *srv) Create(c context.Context, req *models.CreateCustomObjectRequest) (*models.CustomObject, error) {
	var deploymentStatus proto.DeploymentStatus
	if req.DeploymentStatus == models.DeploymentStatusDeployed {
		deploymentStatus = proto.DeploymentStatus_Deployed
	} else {
		deploymentStatus = proto.DeploymentStatus_InDevelopment
	}

	dtoFields := make([]*proto.CustomField, len(req.Fields))
	for i, field := range req.Fields {
		dtoFields[i] = customFieldToDto(field)
	}

	var dtoIndexes []*proto.Index
	if req.Indexes != nil {
		dtoIndexes = funk.Map(req.Indexes, func(index *models.Index) *proto.Index {
			return indexToDto(index)
		}).([]*proto.Index)
	}

	newDto, err := m.client.Create(c, &proto.CreateCustomObjectRequest{
		OrgId:              req.OrgId,
		ObjName:            req.ObjName,
		DeploymentStatus:   deploymentStatus,
		Deprecated:         req.Deprecated,
		Description:        req.Description,
		EnableBulkApi:      req.EnableBulkApi,
		ExternalDataSource: req.ExternalDataSource,
		ExternalName:       req.ExternalName,
		Fields:             dtoFields,
		Indexes:            dtoIndexes,
		Label:              req.Label,
	})

	if err != nil {
		return nil, err
	}

	return customObjectFromDto(newDto), nil
}

func (m *srv) Retrieve(c context.Context, objId string) (*models.CustomObject, error) {
	dto, err := m.client.Retrieve(c, &proto.GetCustomObjectRequest{
		ObjId: objId,
	})
	if err != nil {
		return nil, err
	}

	return customObjectFromDto(dto), nil
}

func (m *srv) Update(c context.Context, model *models.CustomObject) (*models.CustomObject, error) {
	updatedDto, err := m.client.Update(c, customObjectToDto(model))
	if err != nil {
		return nil, err
	}

	return customObjectFromDto(updatedDto), nil
}

func (m *srv) Delete(c context.Context, objId string) error {
	_, err := m.client.Delete(c, &proto.DeleteCustomObjectRequest{
		ObjId: objId,
	})
	return err
}

func (m *srv) FindByObjName(c context.Context, orgId, objName string) (*models.CustomObject, error) {
	dto, err := m.client.FindByObjName(c, &proto.FindByObjNameRequest{
		OrgId:   orgId,
		ObjName: objName,
	})
	if err != nil {
		return nil, err
	}

	return customObjectFromDto(dto), nil
}
