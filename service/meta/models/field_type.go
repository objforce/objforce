package models

import (
	"encoding/binary"
	"errors"
	"math"
)

type FieldType string

const (
	FieldTypeAutoNumber           FieldType = "AutoNumber"
	FieldTypeLookup               FieldType = "Lookup"
	FieldTypeMasterDetail         FieldType = "MasterDetail"
	FieldTypeMetadataRelationship FieldType = "MetadataRelationship"
	FieldTypeCheckbox             FieldType = "Checkbox"
	FieldTypeCurrency             FieldType = "Currency"
	FieldTypeDate                 FieldType = "Date"
	FieldTypeDateTime             FieldType = "DateTime"
	FieldTypeEmail                FieldType = "Email"
	FieldTypeEncryptedText        FieldType = "EncryptedText"
	FieldTypeExternalLookup       FieldType = "ExternalLookup"
	FieldTypeIndirectLookup       FieldType = "IndirectLookup"
	FieldTypeNumber1              FieldType = "Number1"
	FieldTypePercent              FieldType = "Percent"
	FieldTypePhone                FieldType = "Phone"
	FieldTypePicklist             FieldType = "Picklist"
	FieldTypeMultiselectPicklist  FieldType = "MultiselectPicklist"
	FieldTypeSummary              FieldType = "Summary"
	FieldTypeText                 FieldType = "Text"
	FieldTypeTextArea             FieldType = "TextArea"
	FieldTypeLongTextArea         FieldType = "LongTextArea"
	FieldTypeUrl                  FieldType = "Url"
	FieldTypeHierarchy            FieldType = "Hierarchy"
	FieldTypeFile                 FieldType = "File"
	FieldTypeHtml                 FieldType = "Html"
	FieldTypeLocation             FieldType = "Location"
	FieldTypeTime                 FieldType = "Time"
)

var FieldDataTypeGoDataTypeMapping = map[FieldType]string{
	// FieldTypeAddress: "string",
	// FieldTypeAnyType:                    "interface{}",
	// FieldTypeCalculated:                 "string",
	// FieldTypeCombobox:                   "string",
	FieldTypeCurrency: "json.Number",
	// FieldTypeDataCategoryGroupReference: "string",
	FieldTypeEmail:         "string",
	FieldTypeEncryptedText: "string",
	// FieldTypeID:                         "string",
	// FieldTypeIDJunctionIdList:             "[]string",
	FieldTypeLocation: "string",
	// FieldTypeMasterrecord:               "string",
	FieldTypeMultiselectPicklist: "string",
	FieldTypePercent:             "json.Number",
	FieldTypePhone:               "string",
	FieldTypePicklist:            "string",
	// FieldTypeReference:                  "string",
	FieldTypeTextArea: "string",
	FieldTypeUrl:      "string",
}

func (f FieldType) ConvertToGoDataType() string {
	if v, ok := FieldDataTypeGoDataTypeMapping[f]; ok {
		return v
	}
	return "string"
}

func (t FieldType) Marshal(v interface{}) ([]byte, error) {
	var buf []byte
	goDataType := t.ConvertToGoDataType()

	switch goDataType {
	case "string":
		{
			buf = []byte(v.(string))
		}
	case "json.Number":
		{
			switch v.(type) {
			case int64, uint64:
				buf = nil // TODO
			case float32:
				buf = Float32ToBytes(v.(float32))
			case float64:
				buf = Float64ToBytes(v.(float64))
			}
		}
	}
	return buf, nil
}

func (t FieldType) Unmarshal(v []byte) (interface{}, error) {
	goDataType := t.ConvertToGoDataType()

	switch goDataType {
	case "string":
		{
			return string(v), nil
		}
	case "json.Number":
		{
			return nil, errors.New("not impl")
		}
	}
	return nil, errors.New("type not found")
}

func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}
