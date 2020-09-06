package mapper

import (
	"github.com/objforce/objforce/meta-server/app/domain/models"
	"github.com/objforce/objforce/meta-server/app/dtos"
	"github.com/xxxmicro/base/mapper"
	"testing"
)

func TestCustomerField(t *testing.T) {
	dto := &dtos.CustomField{
		Id: "1",
		Label: "name",
		Required: true,
		SummaryOperation: dtos.Count,
		SummaryFilterItems: []*dtos.FilterItem{
			&dtos.FilterItem{
				Field: "name",
				Operation: dtos.FilterOperationEQ,
				Value: "terry",
				ValueField: "name",
			},
		},
	}

	entity := &models.CustomField{}
	err := mapper.Map(dto, entity)
	if err != nil {
		t.Fatal(err)
	}

	entity.Description = "desc"
	err = mapper.Map(entity, dto)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(entity)
}