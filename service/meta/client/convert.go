package client

import (
	proto "github.com/objforce/objforce/proto/meta"
	"github.com/objforce/objforce/service/meta/models"
	"github.com/thoas/go-funk"
)

func customObjectToDto(model *models.CustomObject) *proto.CustomObject {
	var deploymentStatus proto.DeploymentStatus
	if model.DeploymentStatus == models.DeploymentStatusDeployed {
		deploymentStatus = proto.DeploymentStatus_Deployed
	} else {
		deploymentStatus = proto.DeploymentStatus_InDevelopment
	}

	dtoFields := make([]*proto.CustomField, len(model.Fields))
	for i, field := range model.Fields {
		dtoFields[i] = customFieldToDto(field)
	}

	var dtoIndexes []*proto.Index
	if model.Indexes != nil {
		dtoIndexes = funk.Map(model.Indexes, func(index *models.Index) *proto.Index {
			return indexToDto(index)
		}).([]*proto.Index)
	}

	return &proto.CustomObject{
		ObjId:              model.ObjId,
		ObjName:            model.ObjName,
		DeploymentStatus:   deploymentStatus,
		Deprecated:         model.Deprecated,
		Description:        model.Description,
		EnableBulkApi:      model.EnableBulkApi,
		EnableSearch:       model.EnableSearch,
		ExternalDataSource: model.ExternalDataSource,
		ExternalName:       model.ExternalName,
		Label:              model.Label,
		Fields:             dtoFields,
		Indexes:            dtoIndexes,
	}
}

func customObjectFromDto(dto *proto.CustomObject) *models.CustomObject {
	var deploymentStatus models.DeploymentStatus
	if dto.DeploymentStatus == proto.DeploymentStatus_Deployed {
		deploymentStatus = models.DeploymentStatusDeployed
	} else {
		deploymentStatus = models.DeploymentStatusInDevelopment
	}

	modelFields := make([]*models.CustomField, len(dto.Fields))
	for i, field := range dto.Fields {
		modelFields[i] = customFieldFromDto(field)
	}

	var modelIndexes []*models.Index
	if dto.Indexes != nil {
		modelIndexes = funk.Map(dto.Indexes, func(index *proto.Index) *models.Index {
			return indexFromDto(index)
		}).([]*models.Index)
	}

	return &models.CustomObject{
		ObjId:              dto.ObjId,
		ObjName:            dto.ObjName,
		DeploymentStatus:   deploymentStatus,
		Deprecated:         dto.Deprecated,
		Description:        dto.Description,
		EnableBulkApi:      dto.EnableBulkApi,
		EnableSearch:       dto.EnableSearch,
		ExternalDataSource: dto.ExternalDataSource,
		ExternalName:       dto.ExternalName,
		Fields:             modelFields,
		Indexes:            modelIndexes,
		Label:              dto.Label,
	}
}

func customFieldFromDto(pb *proto.CustomField) *models.CustomField {
	return &models.CustomField{
		FieldId:                  pb.FieldId,
		FieldName:                pb.FieldName,
		ObjId:                    pb.ObjId,
		OrgId:                    pb.OrgId,
		DefaultValue:             pb.DefaultValue,
		DeleteConstraint:         models.DeleteConstraint(pb.DeleteConstraint),
		Deprecated:               pb.Deprecated,
		Description:              pb.Description,
		DisplayFormat:            pb.DisplayFormat,
		Encrypted:                pb.Encrypted,
		External:                 pb.External,
		ExternalColumnName:       pb.ExternalColumnName,
		Formula:                  pb.Formula,
		IsSortingDisabled:        pb.IsSortingDisabled,
		Label:                    pb.Label,
		Length:                   int(pb.Length),
		Precision:                int(pb.Precision),
		ReferenceTargetField:     pb.ReferenceTargetField,
		ReferenceTo:              pb.ReferenceTo,
		RelationshipLabel:        pb.RelationshipLabel,
		RelationshipName:         pb.RelationshipName,
		RelationshipOrder:        int(pb.RelationshipOrder),
		ReparentableMasterDetail: pb.ReparentableMasterDetail,
		Required:                 pb.Required,
		Scale:                    int(pb.Scale),
		StartingNumber:           int(pb.StartingNumber),
		StripMarkup:              pb.StripMarkup,
		SummarizedField:          pb.SummarizedField,
		Type:                     fieldTypeFromDto(pb.Type),
		IsUnique:                 pb.IsUnique,
		WriteRequiresMasterRead:  pb.WriteRequiresMasterRead,
		Indexed:                  pb.Indexed,
	}
}

func customFieldToDto(model *models.CustomField) *proto.CustomField {
	return &proto.CustomField{
		FieldId:                  model.FieldId,
		FieldName:                model.FieldName,
		ObjId:                    model.ObjId,
		OrgId:                    model.OrgId,
		DefaultValue:             model.DefaultValue,
		DeleteConstraint:         proto.DeleteConstraint(model.DeleteConstraint),
		Deprecated:               model.Deprecated,
		Description:              model.Description,
		DisplayFormat:            model.DisplayFormat,
		Encrypted:                model.Encrypted,
		External:                 model.External,
		ExternalColumnName:       model.ExternalColumnName,
		Formula:                  model.Formula,
		IsSortingDisabled:        model.IsSortingDisabled,
		Label:                    model.Label,
		Length:                   int32(model.Length),
		Precision:                int32(model.Precision),
		ReferenceTargetField:     model.ReferenceTargetField,
		ReferenceTo:              model.ReferenceTo,
		RelationshipLabel:        model.RelationshipLabel,
		RelationshipName:         model.RelationshipName,
		RelationshipOrder:        int32(model.RelationshipOrder),
		ReparentableMasterDetail: model.ReparentableMasterDetail,
		Required:                 model.Required,
		Scale:                    int32(model.Scale),
		StartingNumber:           int32(model.StartingNumber),
		StripMarkup:              model.StripMarkup,
		SummarizedField:          model.SummarizedField,
		Type:                     fieldTypeToDto(model.Type),
		IsUnique:                 model.IsUnique,
		WriteRequiresMasterRead:  model.WriteRequiresMasterRead,
		Indexed:                  model.Indexed,
	}
}

func indexFromDto(dto *proto.Index) *models.Index {
	var fields []*models.IndexField
	if dto.Fields != nil {
		fields = funk.Map(dto.Fields, func(dto *proto.IndexField) *models.IndexField {
			return &models.IndexField{
				Name:          dto.Name,
				SortDirection: dto.SortDirection,
			}
		}).([]*models.IndexField)
	}

	return &models.Index{
		Fields: fields,
		Label:  dto.Label,
	}
}

func indexToDto(model *models.Index) *proto.Index {
	var fields []*proto.IndexField
	if model.Fields != nil {
		fields = funk.Map(model.Fields, func(model *models.IndexField) *proto.IndexField {
			return &proto.IndexField{
				Name:          model.Name,
				SortDirection: model.SortDirection,
			}
		}).([]*proto.IndexField)
	}

	return &proto.Index{
		Fields: fields,
		Label:  model.Label,
	}
}

func fieldTypeFromDto(dto proto.FieldType) models.FieldType {
	var t models.FieldType
	switch dto {
	case proto.FieldType_AutoNumber:
		t = models.FieldTypeAutoNumber
	case proto.FieldType_Lookup:
		t = models.FieldTypeLookup
	case proto.FieldType_MasterDetail:
		t = models.FieldTypeMasterDetail
	case proto.FieldType_MetadataRelationship:
		t = models.FieldTypeMetadataRelationship
	case proto.FieldType_Checkbox:
		t = models.FieldTypeCheckbox
	case proto.FieldType_Currency:
		t = models.FieldTypeCurrency
	case proto.FieldType_Date:
		t = models.FieldTypeDate
	case proto.FieldType_DateTime:
		t = models.FieldTypeDateTime
	case proto.FieldType_Email:
		t = models.FieldTypeEmail
	case proto.FieldType_EncryptedText:
		t = models.FieldTypeEncryptedText
	case proto.FieldType_ExternalLookup:
		t = models.FieldTypeExternalLookup
	case proto.FieldType_IndirectLookup:
		t = models.FieldTypeIndirectLookup
	case proto.FieldType_Number1:
		t = models.FieldTypeNumber1
	case proto.FieldType_Percent:
		t = models.FieldTypePercent
	case proto.FieldType_Phone:
		t = models.FieldTypePhone
	case proto.FieldType_Picklist:
		t = models.FieldTypePicklist
	case proto.FieldType_MultiselectPicklist:
		t = models.FieldTypeMultiselectPicklist
	case proto.FieldType_Summary:
		t = models.FieldTypeSummary
	case proto.FieldType_Text:
		t = models.FieldTypeText
	case proto.FieldType_TextArea:
		t = models.FieldTypeTextArea
	case proto.FieldType_LongTextArea:
		t = models.FieldTypeLongTextArea
	case proto.FieldType_Url:
		t = models.FieldTypeUrl
	case proto.FieldType_Hierarchy:
		t = models.FieldTypeHierarchy
	case proto.FieldType_File:
		t = models.FieldTypeFile
	case proto.FieldType_Html:
		t = models.FieldTypeHtml
	case proto.FieldType_Location:
		t = models.FieldTypeLocation
	case proto.FieldType_Time:
		t = models.FieldTypeTime
	}

	return t
}

func fieldTypeToDto(dto models.FieldType) proto.FieldType {
	var t proto.FieldType
	switch dto {
	case models.FieldTypeAutoNumber:
		t = proto.FieldType_AutoNumber
	case models.FieldTypeLookup:
		t = proto.FieldType_Lookup
	case models.FieldTypeMasterDetail:
		t = proto.FieldType_MasterDetail
	case models.FieldTypeMetadataRelationship:
		t = proto.FieldType_MetadataRelationship
	case models.FieldTypeCheckbox:
		t = proto.FieldType_Checkbox
	case models.FieldTypeCurrency:
		t = proto.FieldType_Currency
	case models.FieldTypeDate:
		t = proto.FieldType_Date
	case models.FieldTypeDateTime:
		t = proto.FieldType_DateTime
	case models.FieldTypeEmail:
		t = proto.FieldType_Email
	case models.FieldTypeEncryptedText:
		t = proto.FieldType_EncryptedText
	case models.FieldTypeExternalLookup:
		t = proto.FieldType_ExternalLookup
	case models.FieldTypeIndirectLookup:
		t = proto.FieldType_IndirectLookup
	case models.FieldTypeNumber1:
		t = proto.FieldType_Number1
	case models.FieldTypePercent:
		t = proto.FieldType_Percent
	case models.FieldTypePhone:
		t = proto.FieldType_Phone
	case models.FieldTypePicklist:
		t = proto.FieldType_Picklist
	case models.FieldTypeMultiselectPicklist:
		t = proto.FieldType_MultiselectPicklist
	case models.FieldTypeSummary:
		t = proto.FieldType_Summary
	case models.FieldTypeText:
		t = proto.FieldType_Text
	case models.FieldTypeTextArea:
		t = proto.FieldType_TextArea
	case models.FieldTypeLongTextArea:
		t = proto.FieldType_LongTextArea
	case models.FieldTypeUrl:
		t = proto.FieldType_Url
	case models.FieldTypeHierarchy:
		t = proto.FieldType_Hierarchy
	case models.FieldTypeFile:
		t = proto.FieldType_File
	case models.FieldTypeHtml:
		t = proto.FieldType_Html
	case models.FieldTypeLocation:
		t = proto.FieldType_Location
	case models.FieldTypeTime:
		t = proto.FieldType_Time
	}

	return t
}
