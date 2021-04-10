package hbase

import (
	"context"

	"github.com/objforce/objforce/service/data/domain/entities"
	"github.com/objforce/objforce/service/data/domain/repositories"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"github.com/vmihailenco/msgpack"
)

type dataRepository struct {
	client gohbase.Client
}

func NewDataRepository(client gohbase.Client) repositories.DataRepository {
	return &dataRepository{
		client,
	}
}

func (r *dataRepository) Retrieve(c context.Context, orgId, objId, id string, fields []string) (*entities.MTData, error) {
	tableName := objId
	rowkey := id

	families := map[string][]string{
		"basic": {
			"created_at", "created_by", "updated_at", "updated_by",
		},
		"ext": fields,
	}

	req, err := hrpc.NewGetStr(c, tableName, rowkey, hrpc.Families(families))
	if err != nil {
		return nil, err
	}

	rsp, err := r.client.Get(req)
	if err != nil {
		return nil, err
	}

	var fieldValues map[string][]byte
	for _, cell := range rsp.Cells {
		fieldValues[string(cell.Qualifier)] = cell.Value
	}

	goFieldValues, err := unmarshalFields(fieldValues)
	if err != nil {
		return nil, err
	}

	return &entities.MTData{
		GUID:   objId,
		ObjId:  objId,
		Fields: goFieldValues,
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
func (r *dataRepository) Create(c context.Context, entity *entities.MTData) error {
	tableName := entity.ObjId
	rowkey := entity.GUID
	values := make(map[string]map[string][]byte)

	basic := map[string][]byte{}
	var err error
	basic["created_at"], err = msgpack.Marshal(entity.CreatedAt)
	if err != nil {
		return nil
	}

	basic["updated_at"], err = msgpack.Marshal(entity.UpdatedAt)
	if err != nil {
		return nil
	}

	if entity.CreatedBy != nil {
		basic["created_by"], _ = msgpack.Marshal(*entity.CreatedBy)
	}
	if entity.UpdatedBy != nil {
		basic["updated_by"], _ = msgpack.Marshal(*entity.UpdatedBy)
	}

	values["basic"] = basic

	values["ext"], err = marshalFields(entity.Fields)
	if err != nil {
		return nil
	}

	req, err := hrpc.NewPutStr(c, tableName, rowkey, values)
	if err != nil {
		return err
	}

	_, err = r.client.Put(req)
	if err != nil {
		return err
	}

	return nil
}

/**
 * 更新
 */
func (r *dataRepository) Update(c context.Context, entity *entities.MTData) error {
	tableName := entity.ObjId

	values := make(map[string]map[string][]byte)

	basic := map[string][]byte{}

	if entity.UpdatedBy != nil {
		basic["updated_by"], _ = msgpack.Marshal(*entity.UpdatedBy)
	}
	basic["updated_at"], _ = msgpack.Marshal(entity.UpdatedAt)
	values["basic"] = basic

	var err error
	values["ext"], err = marshalFields(entity.Fields)
	if err != nil {
		return err
	}

	req, err := hrpc.NewPutStr(c, tableName, entity.ObjId, values)
	if err != nil {
		return err
	}

	_, err = r.client.Put(req)
	if err != nil {
		return err
	}

	return nil
}

/**
 * 更新
 */
func (r *dataRepository) Upsert(c context.Context, entity *entities.MTData) *entities.UpsertResult {
	tableName := entity.ObjId

	values := make(map[string]map[string][]byte)
	basic := map[string][]byte{}

	if entity.CreatedBy != nil {
		basic["created_by"], _ = msgpack.Marshal(*entity.CreatedBy)
	}

	basic["created_at"], _ = msgpack.Marshal(entity.CreatedAt)
	if entity.UpdatedBy != nil {
		basic["updated_by"], _ = msgpack.Marshal(*entity.UpdatedBy)
	}

	basic["updated_at"], _ = msgpack.Marshal(entity.UpdatedAt)
	values["basic"] = basic

	var err error

	values["ext"], err = marshalFields(entity.Fields)
	if err != nil {
		return &entities.UpsertResult{Created: false, Errors: []error{err}, Success: false}
	}

	req, err := hrpc.NewPutStr(c, tableName, entity.ObjId, values)
	if err != nil {
		return &entities.UpsertResult{Created: false, Errors: []error{err}, Success: false}
	}

	rsp, err := r.client.Put(req)
	if err != nil {
		return &entities.UpsertResult{Created: false, Errors: []error{err}, Success: false}
	}

	return &entities.UpsertResult{Created: *rsp.Exists, Errors: nil, Success: true}
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
	req, err := hrpc.NewDelStr(c, objId, id, nil)
	if err != nil {
		return err
	}

	_, err = r.client.Delete(req)
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
