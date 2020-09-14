package repositories

import(
	"context"
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
		Fields: make(map[string]interface{}),
	}

	for _, cell := range rsp.Cells {
		family := string(cell.Family)
		if family == "basic" {
			qualifier := string(cell.Qualifier)
			switch qualifier {
			case "name":
				data.Name = string(cell.Value)
			case "created_at":
				data.CreatedAt = time.Unix(0, 0)
			case "created_by":
				data.CreatedBy = string(cell.Value)
			case "updated_at":
				data.UpdatedAt = time.Unix(0, 0)
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

func (r *dataRepository) Put(c context.Context, m *models.MTData) error {
	m.GUID = uuid.NewV4().String()
	rowkey := marshalRowKey(m.OrgId, m.ObjId, m.GUID)

	values := map[string]map[string][]byte{
		"basic": {
			"created_by": []byte(m.CreatedBy),
			// "created_at": {m.CreatedAt.UnixNano()},
		},
	}

	ext := map[string][]byte{}
	for fieldName, fieldValue := range m.Fields {
		ext[fieldName] = []byte(fieldValue)
	}

	values["ext"] = ext

	_, err := hrpc.NewPutStr(c, tableName, rowkey, values)
	if err != nil {
		return err
	}

	return nil
}


func marshalRowKey(orgId string, objId string, guid string) string {
	return fmt.Sprintf("%s_%s_%s", orgId, objId, guid)
}