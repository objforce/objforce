package tablestore

import (
	"context"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/duolacloud/microbase/multitenancy"
	"github.com/objforce/objforce/service/data/domain/entities"
	"github.com/objforce/objforce/service/data/domain/repositories"
)

type DataRepositoryTarget struct {
	Tenancy multitenancy.Tenancy `name:"tablestore.tenancy"`
}

type dataRepository struct {
	target DataRepositoryTarget
}

func NewDataRepository(target DataRepositoryTarget) repositories.DataRepository {
	return &dataRepository{
		target,
	}
}

func (r *dataRepository) clientBy(c context.Context, orgId string) (*tablestore.TableStoreClient, error) {
	res, err := r.target.Tenancy.ResourceFor(c, orgId)
	if err != nil {
		return nil, err
	}
	return res.(*tablestore.TableStoreClient), nil
}

func (r *dataRepository) Retrieve(c context.Context, orgId, objId, id string, fields []string) (*entities.MTData, error) {
	client, err := r.clientBy(c, orgId)
	if err != nil {
		return nil, err
	}

	tableName := objId

	req := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	pk := new(tablestore.PrimaryKey)
	pk.AddPrimaryKeyColumn("id", id)
	criteria.PrimaryKey = pk
	req.SingleRowQueryCriteria = criteria
	req.SingleRowQueryCriteria.TableName = tableName
	req.SingleRowQueryCriteria.MaxVersion = 1

	rsp, err := client.GetRow(req)
	if err != nil {
		return nil, err
	}

	var fieldValues map[string]interface{}
	for _, column := range rsp.Columns {
		fieldValues[column.ColumnName] = column.Value
	}

	return &entities.MTData{
		GUID:   objId,
		ObjId:  objId,
		Fields: fieldValues,
	}, nil
}

/**
 * 获取多条记录
 */
func (r *dataRepository) MultiGet(c context.Context, orgId string, objId string, ids []string, fields []string) []*entities.MTData {
	list := make([]*entities.MTData, len(ids))
	for i, id := range ids {
		data, err := r.Retrieve(c, orgId, objId, id, fields)
		if err != nil {
			continue
		}
		list[i] = data
	}
	return list
}

/**
 * 创建
 * 局部更新，只修改部分列
 */
func (r *dataRepository) Create(c context.Context, m *entities.MTData) error {
	client, err := r.clientBy(c, m.OrgId)
	if err != nil {
		return err
	}

	tableName := m.ObjId

	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = tableName
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("id", m.GUID)

	if m.CreatedBy != nil {
		putRowChange.AddColumn("created_by", *m.CreatedBy)
	}

	putRowChange.AddColumn("created_at", m.CreatedAt)
	if m.UpdatedBy != nil {
		putRowChange.AddColumn("updated_by", *m.UpdatedBy)
	}
	putRowChange.AddColumn("updated_at", m.UpdatedAt)

	for fieldName, fieldValue := range m.Fields {
		putRowChange.AddColumn(fieldName, fieldValue)
	}

	_, err = client.PutRow(putRowRequest)
	if err != nil {
		return err
	}

	return nil
}

/**
 * 更新
 */
func (r *dataRepository) Update(c context.Context, entity *entities.MTData) error {
	client, err := r.clientBy(c, entity.OrgId)
	if err != nil {
		return err
	}

	tableName := entity.ObjId

	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = tableName
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("id", entity.GUID)

	if entity.UpdatedBy != nil {
		updateRowChange.PutColumn("updated_by", *entity.UpdatedBy)
	}
	updateRowChange.PutColumn("updated_at", entity.UpdatedAt)

	for fieldName, fieldValue := range entity.Fields {
		updateRowChange.PutColumn(fieldName, fieldValue)
	}

	_, err = client.UpdateRow(updateRowRequest)
	if err != nil {
		return err
	}

	return nil
}

/**
 * 更新
 */
func (r *dataRepository) Upsert(c context.Context, entity *entities.MTData) *entities.UpsertResult {
	returndEntity, err := r.Retrieve(c, entity.OrgId, entity.ObjId, entity.GUID, []string{"id"})
	if err != nil {
		return &entities.UpsertResult{Errors: []error{err}}
	}

	if returndEntity != nil {
		err := r.Update(c, entity)
		if err != nil {
			return &entities.UpsertResult{Created: false, Errors: []error{err}, Success: false}
		} else {
			return &entities.UpsertResult{Created: false, Errors: nil, Success: false}
		}
	}

	err = r.Create(c, entity)
	if err != nil {
		return &entities.UpsertResult{Created: false, Errors: []error{err}, Success: false}
	}

	return &entities.UpsertResult{Created: true, Errors: nil, Success: true}
}

func (r *dataRepository) MultiCreate(c context.Context, items []*entities.MTData) []*entities.SaveResult {
	results := make([]*entities.SaveResult, len(items))
	for i, data := range items {
		success := true

		err := r.Create(c, data)
		if err != nil {
			success = false
		}

		result := &entities.SaveResult{
			Id:      data.GUID,
			Success: success,
			Error:   err,
		}

		results[i] = result
	}
	return results
}

func (r *dataRepository) MultiUpdate(c context.Context, items []*entities.MTData) []*entities.SaveResult {
	results := make([]*entities.SaveResult, len(items))
	for i, data := range items {
		success := true

		err := r.Update(c, data)
		if err != nil {
			success = false
		}

		result := &entities.SaveResult{
			Id:      data.GUID,
			Success: success,
			Error:   err,
		}

		results[i] = result
	}
	return results
}
func (r *dataRepository) MultiUpsert(c context.Context, items []*entities.MTData) []*entities.UpsertResult {
	results := make([]*entities.UpsertResult, len(items))
	for i, data := range items {
		result := r.Upsert(c, data)
		results[i] = result
	}
	return results
}

func (r *dataRepository) Delete(c context.Context, orgId string, objId string, id string) error {
	client, err := r.clientBy(c, orgId)
	if err != nil {
		return err
	}

	deleteRowReq := new(tablestore.DeleteRowRequest)
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)
	deleteRowReq.DeleteRowChange.TableName = objId
	deletePk := new(tablestore.PrimaryKey)
	deletePk.AddPrimaryKeyColumn("id", id)
	deleteRowReq.DeleteRowChange.PrimaryKey = deletePk

	_, err = client.DeleteRow(deleteRowReq)
	if err != nil {
		return err
	}

	return nil
}

func (r *dataRepository) MultiDelete(c context.Context, orgId string, objId string, ids []string) []*entities.DeleteResult {
	results := make([]*entities.DeleteResult, len(ids))
	for i, id := range ids {
		success := true
		err := r.Delete(c, orgId, objId, id)
		if err != nil {
			success = false
		}

		results[i] = &entities.DeleteResult{
			Error:   err,
			Id:      id,
			Success: success,
		}
	}
	return results
}
