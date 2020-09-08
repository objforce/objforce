package models

type Index struct {
	ObjId string `json:"objId"`
}

func (m *Index) Unique() interface{} {
	return map[string]interface{}{
		"objId": m.ObjId,
	}
}