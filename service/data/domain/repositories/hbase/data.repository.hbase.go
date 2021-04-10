package hbase

import (
	"context"
	"strconv"
	"time"

	"github.com/objforce/objforce/service/data/domain/entities"
	"github.com/objforce/objforce/service/data/domain/repositories"
	"github.com/objforce/objforce/utils"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

const tableName = "mt_datas"

type dataRepository struct {
	client gohbase.Client
}

func NewDataRepository(client gohbase.Client) repositories.DataRepository {
	return &dataRepository{
		client,
	}
}

func (r *dataRepository) Get(c context.Context, orgId, objId, id string, fields []string) (*entities.MTData, error) {
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

	data := &entities.MTData{
		GUID:   objId,
		ObjId:  objId,
		Fields: make(map[string][]byte),
	}

	for _, cell := range rsp.Cells {
		family := string(cell.Family)
		if family == "basic" {
			qualifier := string(cell.Qualifier)
			switch qualifier {
			case "created_at":
				ms := utils.BytesToInt64(cell.Value)
				data.CreatedAt = time.Unix(0, ms*int64(time.Millisecond))
			case "created_by":
				createdBy := string(cell.Value)
				data.CreatedBy = &createdBy
			case "updated_at":
				ms := utils.BytesToInt64(cell.Value)
				data.UpdatedAt = time.Unix(0, ms*int64(time.Millisecond))
			case "updated_by":
				updatedBy := string(cell.Value)
				data.UpdatedBy = &updatedBy
			case "isDeleted":
				data.IsDeleted, _ = strconv.ParseBool(string(cell.Value))
			}
		}

		data.Fields[string(cell.Qualifier)] = cell.Value
	}

	return data, nil
}

/**
 * 获取多条记录
 */
func (r *dataRepository) MultiGet(c context.Context, orgId string, objId string, ids []string, fields []string) []*entities.MTData {
	list := make([]*entities.MTData, len(ids))
	for i, id := range ids {
		data, err := r.Get(c, orgId, objId, id, fields)
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
	rowkey := m.GUID
	values := make(map[string]map[string][]byte)

	basic := map[string][]byte{}
	if m.CreatedBy != nil {
		basic["created_by"] = []byte(*m.CreatedBy)
	}
	basic["created_at"] = utils.Int64ToBytes(m.CreatedAt.UnixNano() / 1e6)
	if m.UpdatedBy != nil {
		basic["updated_by"] = []byte(*m.UpdatedBy)
	}
	basic["updated_at"] = utils.Int64ToBytes(m.UpdatedAt.UnixNano() / 1e6)

	values["basic"] = basic

	ext := map[string][]byte{}
	for fieldName, fieldValue := range m.Fields {
		ext[fieldName] = []byte(fieldValue)
	}

	values["ext"] = ext

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
func (r *dataRepository) Update(c context.Context, m *entities.MTData) error {
	values := make(map[string]map[string][]byte)

	basic := map[string][]byte{}

	if m.UpdatedBy != nil {
		basic["updated_by"] = []byte(*m.UpdatedBy)
	}
	basic["updated_at"] = utils.Int64ToBytes(m.UpdatedAt.UnixNano() / 1e6)
	values["basic"] = basic

	ext := map[string][]byte{}
	for fieldName, fieldValue := range m.Fields {
		ext[fieldName] = []byte(fieldValue)
	}

	values["ext"] = ext

	req, err := hrpc.NewPutStr(c, tableName, m.ObjId, values)
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
	values := make(map[string]map[string][]byte)
	basic := map[string][]byte{}

	if entity.CreatedBy != nil {
		basic["created_by"] = []byte(*entity.CreatedBy)
	}

	basic["created_at"] = utils.Int64ToBytes(entity.CreatedAt.UnixNano() / 1e6)
	if entity.UpdatedBy != nil {
		basic["updated_by"] = []byte(*entity.UpdatedBy)
	}

	basic["updated_at"] = utils.Int64ToBytes(entity.UpdatedAt.UnixNano() / 1e6)
	values["basic"] = basic

	values["ext"] = entity.Fields

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
