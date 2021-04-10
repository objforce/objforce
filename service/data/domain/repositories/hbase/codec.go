package hbase

import (
	"github.com/vmihailenco/msgpack"
)

func marshalFields(fields map[string]interface{}) (map[string][]byte, error) {
	var r map[string][]byte
	var err error
	for fieldName, fieldValue := range fields {
		r[fieldName], err = msgpack.Marshal(fieldValue)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func unmarshalFields(fields map[string][]byte) (map[string]interface{}, error) {
	var r map[string]interface{}
	var err error
	for fieldName, fieldValue := range fields {
		var v interface{} // TODO 这里要根据字段类型来声明不同 golang类型
		err = msgpack.Unmarshal(fieldValue, &v)
		if err != nil {
			return nil, err
		}
		r[fieldName] = v
	}
	return r, nil
}
