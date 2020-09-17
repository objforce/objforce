package repositories

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/objforce/objforce/data-srv/app/domain/models"
	uuid "github.com/satori/go.uuid"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"strconv"
	"time"
)

const tableName = "mt_datas"

type DataRepository interface {
	Get(c context.Context, orgId string, objId string, guid string, families map[string][]string) (*models.MTData, error)
	Create(c context.Context, m *models.MTData) error
	Update(c context.Context, m *models.MTData) error
	Upsert(c context.Context, m *models.MTData) *models.UpsertResult
	Delete(c context.Context, orgId string, objId string, guid string) error

	MultiGet(c context.Context, orgId string, objId string, ids []string, families map[string][]string) []*models.MTData
	MultiCreate(c context.Context, items []*models.MTData) []*models.SaveResult
	MultiUpdate(c context.Context, items []*models.MTData) []*models.SaveResult
	MultiUpsert(c context.Context, items []*models.MTData) []*models.UpsertResult
	MultiDelete(c context.Context, orgId string, objId string, ids []string) []*models.DeleteResult
}

type dataRepository struct {
	client gohbase.Client
}

func NewDataRepository(client gohbase.Client) DataRepository {
	return &dataRepository{
		client,
	}
}

func (r* dataRepository) Get(c context.Context, orgId string, objId string, guid string, families map[string][]string) (*models.MTData, error) {
	rowkey := marshalRowKey(orgId, objId, guid)

	req, err := hrpc.NewGetStr(c, tableName, rowkey, hrpc.Families(families))
	if err != nil {
		return nil, err
	}

	rsp, err := r.client.Get(req)
	if err != nil {
		return nil, err
	}

	data := &models.MTData{
		GUID: objId,
		OrgId: orgId,
		ObjId: objId,
		Fields: make(map[string]string),
	}

	for _, cell := range rsp.Cells {
		family := string(cell.Family)
		if family == "basic" {
			qualifier := string(cell.Qualifier)
			switch qualifier {
			case "name":
				data.Name = string(cell.Value)
			case "created_at":
				ms := BytesToInt64(cell.Value)
				data.CreatedAt = time.Unix(0, ms * int64(time.Millisecond))
			case "created_by":
				data.CreatedBy = string(cell.Value)
			case "updated_at":
				ms := BytesToInt64(cell.Value)
				data.UpdatedAt = time.Unix(0, ms * int64(time.Millisecond))
			case "updated_by":
				data.UpdatedBy = string(cell.Value)
			case "isDeleted":
				data.IsDeleted, _ = strconv.ParseBool(string(cell.Value))
			}
		}

		data.Fields[string(cell.Qualifier)] = string(cell.Value)
	}

	return data, nil
}

/**
 * 获取多条记录
 */
func (r *dataRepository) MultiGet(c context.Context, orgId string, objId string, ids []string, families map[string][]string) []*models.MTData {
	list := make([]*models.MTData, len(ids))
	for i, id := range ids {
		data, err := r.Get(c, orgId, objId, id, families)
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
func (r *dataRepository) Create(c context.Context, m *models.MTData) error {
	m.GUID = uuid.NewV4().String()
	rowkey := marshalRowKey(m.OrgId, m.ObjId, m.GUID)
	values := make(map[string]map[string][]byte)

	basic := map[string][]byte{
		"created_by": []byte(m.CreatedBy),
		"created_at": Int64ToBytes(m.CreatedAt.UnixNano()/1e6),
		"updated_by": []byte(m.UpdatedBy),
		"updated_at": Int64ToBytes(m.UpdatedAt.UnixNano()/1e6),
	}

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
func (r *dataRepository) Update(c context.Context, m *models.MTData) error {
	m.GUID = uuid.NewV4().String()
	rowkey := marshalRowKey(m.OrgId, m.ObjId, m.GUID)
	values := make(map[string]map[string][]byte)

	basic := map[string][]byte{
		"created_by": []byte(m.CreatedBy),
		"created_at": Int64ToBytes(m.CreatedAt.UnixNano()/1e6),
		"updated_by": []byte(m.UpdatedBy),
		"updated_at": Int64ToBytes(m.UpdatedAt.UnixNano()/1e6),
	}

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
func (r *dataRepository) Upsert(c context.Context, m *models.MTData) *models.UpsertResult {
	m.GUID = uuid.NewV4().String()
	rowkey := marshalRowKey(m.OrgId, m.ObjId, m.GUID)
	values := make(map[string]map[string][]byte)

	basic := map[string][]byte{
		"created_by": []byte(m.CreatedBy),
		"created_at": Int64ToBytes(m.CreatedAt.UnixNano()/1e6),
		"updated_by": []byte(m.UpdatedBy),
		"updated_at": Int64ToBytes(m.UpdatedAt.UnixNano()/1e6),
	}

	values["basic"] = basic

	ext := map[string][]byte{}
	for fieldName, fieldValue := range m.Fields {
		ext[fieldName] = []byte(fieldValue)
	}

	values["ext"] = ext

	req, err := hrpc.NewPutStr(c, tableName, rowkey, values)
	if err != nil {
		return &models.UpsertResult{ Created: false, Error: err, Success: false }
	}

	rsp, err := r.client.Put(req)
	if err != nil {
		return &models.UpsertResult{ Created: false, Error: err, Success: false }
	}

	return &models.UpsertResult{ Created: *rsp.Exists, Error: nil, Success: true }
}

func (r *dataRepository) MultiCreate(c context.Context, items []*models.MTData) []*models.SaveResult {
	results := make([]*models.SaveResult, len(items))
	for i, data := range items {
		success := true

		err := r.Create(c, data)
		if err != nil {
			success = false
		}

		result := &models.SaveResult{
			Id: data.GUID,
			Success: success,
			Error: err,
		}

		results[i] = result
	}
	return results
}

func (r *dataRepository) MultiUpdate(c context.Context, items []*models.MTData) []*models.SaveResult {
	results := make([]*models.SaveResult, len(items))
	for i, data := range items {
		success := true

		err := r.Update(c, data)
		if err != nil {
			success = false
		}

		result := &models.SaveResult{
			Id: data.GUID,
			Success: success,
			Error: err,
		}

		results[i] = result
	}
	return results
}
func (r *dataRepository) MultiUpsert(c context.Context, items []*models.MTData) []*models.UpsertResult {
	results := make([]*models.UpsertResult, len(items))
	for i, data := range items {
		result := r.Upsert(c, data)
		results[i] = result
	}
	return results
}

func (r *dataRepository) Delete(c context.Context, orgId string, objId string, guid string) error {
	rowkey := marshalRowKey(orgId, objId, guid)

	req, err := hrpc.NewDelStr(c, tableName, rowkey, nil)
	if err != nil {
		return err
	}

	_, err = r.client.Delete(req)
	if err != nil {
		return err
	}

	return nil
}

func (r *dataRepository) MultiDelete(c context.Context, orgId string, objId string, ids []string) []*models.DeleteResult {
	results := make([]*models.DeleteResult, len(ids))
	for i, id := range ids {
		success := true
		err := r.Delete(c, orgId, objId, id)
		if err != nil {
			success = false
		}

		results[i] = &models.DeleteResult{
			Error: err,
			Id: id,
			Success: success,
		}
	}
	return results
}


func marshalRowKey(orgId string, objId string, guid string) string {
	return fmt.Sprintf("%s_%s_%s", orgId, objId, guid)
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}