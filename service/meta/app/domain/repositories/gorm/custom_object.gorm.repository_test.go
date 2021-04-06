package gorm

import (
	"context"
	"testing"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/memory"
	"github.com/objforce/objforce/service/meta/app/domain/models"
	"github.com/xxxmicro/base/database/gorm"
)

func TestCustomObjectRepository_Create(t *testing.T) {
	source := memory.NewSource(
		memory.WithJSON([]byte(`
		{
			"db": {
				"driver": "mysql",
				"connection_string": "root:root@tcp(localhost:3306)/objforce_meta?charset=utf8mb4&parseTime=True&loc=Local"
			}
		}`)))

	config, err := config.NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	err = config.Load(source)
	if err != nil {
		t.Fatal(err)
	}

	db, err := gorm.NewDbProvider(config)
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(
		&models.MTObject{},
		&models.MTField{},
	)

	orgId := "1234567890123456789012345678901234567890"

	customObjectRepository := NewCustomObjectRepository(db)

	customObject := &models.MTObject{
		ObjName: "customers",
		OrgId:   orgId,
		Fields: []*models.MTField{
			{
				OrgId:     orgId,
				FieldName: "age",
				Type:      models.FieldTypeAutoNumber,
			},
			{
				OrgId:     orgId,
				FieldName: "name",
				Type:      models.FieldTypeText,
			},
		},
	}

	err = customObjectRepository.Upsert(context.Background(), customObject)
	if err != nil {
		t.Fatal(err)
	}

	customObject, err = customObjectRepository.Retrieve(context.Background(), customObject.ObjId)
	if err != nil {
		t.Fatal(err)
	}

	customObject.Fields = []*models.MTField{
		{
			OrgId:     orgId,
			FieldName: "balance",
			Type:      models.FieldTypeCurrency,
		},
		{
			FieldId:   customObject.Fields[1].FieldId,
			OrgId:     orgId,
			FieldName: "name1",
			Type:      models.FieldTypeText,
		},
	}

	err = customObjectRepository.Update(context.Background(), customObject)
	if err != nil {
		t.Fatal(err)
	}
}
