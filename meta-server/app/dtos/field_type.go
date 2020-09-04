package dtos

type FieldType string

const(
	AutoNumber FieldType = "AutoNumber"
	Lookup FieldType = "Lookup"
	MasterDetail FieldType = "MasterDetail"
	MetadataRelationship FieldType = "MetadataRelationship"
	Checkbox FieldType = "Checkbox"
	Currency FieldType = "Currency"
	Date FieldType = "Date"
	DateTime FieldType = "DateTime"
	Email FieldType = "Email"
	EncryptedText FieldType = "EncryptedText"
	ExternalLookup FieldType = "ExternalLookup"
	IndirectLookup FieldType = "IndirectLookup"
	Number1 FieldType = "Number1"
	Percent FieldType = "Percent"
	Phone FieldType = "Phone"
	Picklist FieldType = "Picklist"
	MultiselectPicklist FieldType = "MultiselectPicklist"
	Summary FieldType = "Summary"
	Text FieldType = "Text"
	TextArea FieldType = "TextArea"
	LongTextArea FieldType = "LongTextArea"
	Url FieldType = "Url"
	Hierarchy FieldType = "Hierarchy"
	File FieldType = "File"
	Html FieldType = "Html"
	Location FieldType = "Location"
	Time FieldType = "Time"
)