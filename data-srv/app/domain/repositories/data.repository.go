package repositories

import(
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
	Get(c context.Context, orgId string, objId string, guid string) (*models.MTData, error)
	Put(c context.Context, m *models.MTData) error
	Delete(c context.Context, orgId string, objId string, guid string) error
}

type dataRepository struct {
	client gohbase.Client
}

func NewDataRepository(client gohbase.Client) DataRepository {
	return &dataRepository{
		client,
	}
}

func (r* dataRepository) Get(c context.Context, orgId string, objId string, guid string) (*models.MTData, error) {
	rowkey := marshalRowKey(orgId, objId, guid)

	req, err := hrpc.NewGetStr(c, tableName, rowkey)
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
func (r *dataRepository) MultiGet(c context.Context, ids []string, fields []string) ([]*models.MTData, error) {
	req, err := hrpc.NewGetStr()
	if err != nil {
		return nil, err
	}

	r.client.
}

/**
 * 创建
 * 局部更新，只修改部分列
 */
func (r *dataRepository) Put(c context.Context, m *models.MTData) error {
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