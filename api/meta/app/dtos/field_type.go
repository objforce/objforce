package dtos

type FieldType string

const(
	FieldTypeAutoNumber FieldType = "AutoNumber"
	FieldTypeLookup FieldType = "Lookup"
	FieldTypeMasterDetail FieldType = "MasterDetail"
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
	FieldTypePercent                       FieldType = "Percent"
	FieldTypePhone                         FieldType = "Phone"
	FieldTypePicklist                      FieldType = "Picklist"
	FieldTypeMultiselectPicklist           FieldType = "MultiselectPicklist"
	FieldTypeSummary                       FieldType = "Summary"
	FieldTypeText                          FieldType = "Text"
	FieldTypeTextArea                      FieldType = "TextArea"
	FieldTypeLongTextArea                  FieldType = "LongTextArea"
	FieldTypeUrl                           FieldType = "Url"
	FieldTypeHierarchy                     FieldType = "Hierarchy"
	FieldTypeFile                          FieldType = "File"
	FieldTypeHtml FieldType = "Html"
	FieldTypeLocation FieldType = "Location"
	FieldTypeTime FieldType = "Time"
)