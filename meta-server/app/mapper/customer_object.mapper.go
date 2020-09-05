package mapper

import (
	"github.com/objforce/meta-server/app/domain/models"
	"github.com/objforce/meta-server/app/dtos"
	"github.com/petersunbag/coven"
)

type CustomerObjectMapper struct {
	c1 *coven.Converter
	c2 *coven.Converter
}

var CUSTOMER_OBJECT_MAPPER *CustomerObjectMapper

func init() {
	CUSTOMER_OBJECT_MAPPER = newCustomerObjectMapper()
}

func newCustomerObjectMapper() *CustomerObjectMapper {
	c1, err := coven.NewConverter(models.CustomObject{}, dtos.CustomObject{})
	if err != nil {
		panic(err)
	}

	c2, err := coven.NewConverter(dtos.CustomObject{}, models.CustomObject{})
	if err != nil {
		panic(err)
	}

	return &CustomerObjectMapper{
		c1: c1,
		c2: c2,
	}
}

func (s *CustomerObjectMapper) ConvertToEntity(src interface{}, dst interface{}) error {
	return s.c1.Convert(dst, src)
}

func (s *CustomerObjectMapper) ConvertToDto(src interface{}, dst interface{}) error {
	return s.c1.Convert(dst, src)
}