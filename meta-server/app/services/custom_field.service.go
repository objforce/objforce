package services

import(
	"github.com/objforce/meta-server/app/models"
	"github.com/objforce/meta-server/app/repositories"
)

type customFieldService struct {
	customFieldRepository repositories.CustomFieldRepository
}

func (s *customFieldService) Create(customField models.CustomField) models.CustomField {
	newCustomField := s.customFieldRepository.Create(customField)
	return newCustomField
}

type CustomFieldService interface {
	Create(customField models.CustomField) models.CustomField
}

func NewCustomFieldService(customFieldRepository repositories.CustomFieldRepository) CustomFieldService {
	return &customFieldService{
		customFieldRepository,
	}
}