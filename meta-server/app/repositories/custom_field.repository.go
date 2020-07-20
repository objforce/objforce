package repositories

import(
	"github.com/objforce/meta-server/app/models"
	"github.com/objforce/meta-server/app/providers"
)

type customFieldRepository struct {
	database providers.DatabaseProvider
}

func (r *customFieldRepository) Create(customField models.CustomField) models.CustomField {
	db := r.database.Connect()
	defer db.Close()

	db.Create(&customField)
	return customField
}

type CustomFieldRepository interface {
	Create(customField models.CustomField) models.CustomField
}

func NewCustomFieldRepository(database providers.DatabaseProvider) CustomFieldRepository {
	return &customFieldRepository{
		database,
	}
}