package models



type CustomField struct {
	Id string `json:"id"`
	ObjId string `json:"objId"`
	DefaultValue string `json:"defaultValue"`
	DeleteConstraint DeleteConstraint `json:"deleteConstraint"`
	Deprecated bool `json:"deprecated"`	// 是否保留使用
	Description string `json:"description"`	// 字段描述
	DisplayFormat string `json:"displayFormat"` // 展示格式
	Encrypted bool `json:encrypted` // 字段是否加密
	EncryptionScheme EncryptionScheme `json:"encryptionScheme"`
	External bool `json:"external"` // 表明是否一个外部映射列
	ExternalColumnName string `json:"externalColumnName"` // 对应的外部数据源 的 列名
	Formula string `json:"formula"` // 应用在字段上的公式
	FormulaTreatBlankAs TreatBlanksAs `json:"formulaTreatBlankAs "` // 如何处理空格
	InlineHelpText string `json:"inlineHelpText"` // 字段的帮助提示
	IsFilteringDisabled bool `json:"isFilteringDisabled"` // Available only for external objects. Indicates whether the custom field is available in filters
	
	/**
		Available only for external objects. Indicates whether the custom field is sortable.
	 */
	IsSortingDisabled bool `json:"isSortingDisabled"`
	
	/**
		Label for the field. You cannot update the label for standard picklist fields, such as the Industry field for accounts.
	 */
	Label string `json:"label"`
	
	/**
		Length of the field.
	 */
	Length int `json:"length"`
	
	/**
		In custom metadata relationships, represents the controlling field that specifies the standard or custom object in an entity definition metadata relationship. 
		Required when creating a field definition or entity particle metadata relationship on a custom metadata type. 
		The object specified in the controlling field determines the values available in its dependent field definition or entity particle. 
		For example, specifying the Account object filters the available fields in the field definition to Account fields only.
   */
	MetadataRelationshipControllingField string `json:"metadataRelationship​ControllingField"`

	/*
		The precision for number values. Precision is the number of digits in a number. For example, the number 256.99 has a precision of 5
	 */
	Precision int `json:"precision"`

	/**
		Available only for indirect lookup relationship fields on external objects. 
		Specifies the custom field on the parent object to match against this indirect lookup relationship field, whose values come from an external data source. The specified custom field on the parent object must have both externalId and unique set to true
	 */
	ReferenceTargetField string `json:"ReferenceTargetField"`
	
	/**
		If specified, indicates a reference this field has to another object
	 */
	ReferenceTo string `json:"referenceTo"`
	
	/**
		Label for the relationship
	 */
	RelationshipLabel string `json:"relationshipLabel"`
	
	/**
		If specified, indicates the value for one-to-many relationships. 
		For example, in the object MyObject that had a relationship to YourObject, the relationship name might be YourObjects.
	 */
	RelationshipName string `json:"relationshipName"`
	
	/**
		This field is valid for all master-detail relationships, but the value is only non-zero for junction objects. 
		A junction object has two master-detail relationships, and is analogous to an association table in a many-to-many relationship. 
		Junction objects must define one parent object as primary (0), the other as secondary (1). 
		The definition of primary or secondary affects delete behavior and inheritance of look and feel, and record ownership for junction objects. 
		For more information, see the Salesforce Help.
		0 or 1 are the only valid values, and 0 is always the value for objects that are not junction objects.
	 */
	RelationshipOrder int `json:"relationshipOrder"`

	/**
	  Indicates whether the child records in a master-detail relationship on a custom object can be reparented to different parent records. The default value is false
	 */
	ReparentableMasterDetail bool `json:"reparentableMasterDetail"`
	/**
	 * Indicates whether the field requires a value on creation (true) or not (false).
	 */
	Required bool `json:"required"`

	/**
	 * The scale for the field. Scale is the number of digits to the right of the decimal point in a number. For example, the number 256.99 has a scale of 2.
	 */
	Scale int `json:"scale"`

	SecurityClassification SecurityClassification `json:"securityClassification"`

	/**
	 * If specified, indicates the starting number for the field. When you create records, Starting Number’s value increments to store the number that will be assigned to the next auto-number field created
	 */
	StartingNumber int `json:"startingNumber"`

	/**
		Set to true to remove markup, or false to preserve markup. Used when converting a rich text area to a long text area.
	 */
	StripMarkup bool `json:"stripMarkup"`

	/**
		Represents the field on the detail row that is being summarized. This field cannot be null unless the summaryOperation value is count.
	 */
	SummarizedField string `json:"summarizedField"`

	/**
	 Represents the set of filter conditions for this field if it is a summary field. This field will be summed on the child if the filter conditions are met.
	 */
	SummaryFilterItems []FilterItem `json:"summaryFilterItems"`

	/**
	 Represents the master-detail field on the child that defines the relationship between the parent and the child
	 */
	SummaryForeignKey string `json:"summaryForeignKey"`

	/**
		Represents the sum operation to be performed. Valid values are enumerated in SummaryOperations.
	 */
	SummaryOperation SummaryOperation `json:"summaryOperation"`

	/**
		Indicates whether the field is enabled for feed tracking (true) or not (false).
		To set this field to true, the enableFeeds field on the associated CustomObject must also be true.
		For more information, see “Customize Chatter Feed Tracking” in the Salesforce Help.
	 */
	TrackFeedHistory bool `json:"trackFeedHistory"`

	/**
		Indicates whether history tracking is enabled for the field (true) or not (false). 
		Also available for standard object fields (picklist and lookup fields only) in API version 30.0 and later.
		To set trackHistory to true, the enableHistory field on the associated standard or custom object must also be true.
		For more information, see “Field History Tracking” in the Salesforce Help.
		Field history tracking isn’t available for external objects.
	 */
	TrackHistory bool `json:"trackHistory"`

	/**
		Indicates whether historical trending data is captured for the field (true) or not (false).
		An object is enabled for historical trending if this attribute is true for at least one field.
	 */
	TrackTrending bool `json:"trackTrending"`

	/**
		Indicates the field type for the field. Valid values are enumerated in FieldType.
		For standard fields on standard objects, the type field is optional. 
		This field is included for some standard field types, such as Picklist or Lookup, but not for others.
		The type field is included for custom fields.
	 */
	Type FieldType `json:"type"`
	/**
		Indicates whether the field is unique (true) or not (false).
	 */
	Unique bool `json:"unique"`

	/**

	 */
	ValueSet ValueSet `json:"valueSet"`

	/**
		Indicates the number of lines displayed for the field.
	 */
	VisibleLines int `json:"visibleLines"`

	/**
		Sets the minimum sharing access level required on the master record to create, edit, or delete child records. 
		This field applies only to master-detail or junction object custom field types.
		true—Allows users with “Read” access to the master record permission to create, edit, or delete child records. This setting makes sharing less restrictive.
		false—Allows users with “Read/Write” access to the master record permission to create, edit, or delete child records. 
		This setting is more restrictive than true, and is the default value.
		For junction objects, the most restrictive access from the two parents is enforced. 
		For example, if you set to true on both master-detail fields, but users have “Read” access to one master record and “Read/Write” access to the other master record, users won't be able to create, edit, or delete child records.
	 */
	WriteRequiresMasterRead bool `json:"writeRequiresMasterRead"`
}