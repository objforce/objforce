package mapper

import(
	"github.com/objforce/meta-server/app/domain/models"
	"github.com/objforce/meta-server/app/dtos"
	"github.com/petersunbag/coven"
	"testing"
)

func TestCustomerField(t *testing.T) {
	c1, err := coven.NewConverter(models.CustomField{}, dtos.CustomField{})
	c2, err := coven.NewConverter(dtos.CustomField{}, models.CustomField{})
	if err != nil {
		panic(err)
	}

	dto := &dtos.CustomField{
		Id: "1",
		Label: "name",
		Required: true,
		SummaryOperation: dtos.Count,
	}

	entity := &models.CustomField{}
	err = c1.Convert(entity, dto)
	if err != nil {
		t.Fatal(err)
	}

	entity.Description = "desc"
	err = c2.Convert(dto, entity)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(entity)
}