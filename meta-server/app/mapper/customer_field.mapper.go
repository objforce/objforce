package mapper

import (
	"github.com/objforce/meta-server/app/domain/models"
	"github.com/objforce/meta-server/app/dtos"
	"github.com/petersunbag/coven"
)

type CustomerFieldMapper struct {
	c1 *coven.Converter
	c2 *coven.Converter
}

var CUSTOMER_FIELD_MAPPER *CustomerFieldMapper

func init() {
	CUSTOMER_FIELD_MAPPER = newCustomerFieldMapper()
}

func newCustomerFieldMapper() *CustomerFieldMapper {
	c1, err := coven.NewConverter(models.CustomField{}, dtos.CustomField{})
	if err != nil {
		panic(err)
	}

	c2, err := coven.NewConverter(models.CustomField{}, dtos.CustomField{})
	if err != nil {
		panic(err)
	}

	return &CustomerFieldMapper{
		c1: c1,
		c2: c2,
	}
}

func (s *CustomerFieldMapper) ConvertToEntity(src interface{}, dst interface{}) error {
	return s.c1.Convert(dst, src)
}

func (s *CustomerFieldMapper) ConvertToDto(src interface{}, dst interface{}) error {
	return s.c1.Convert(dst, src)
}