package repositories

import(
	_gorm "github.com/jinzhu/gorm"
	"github.com/xxxmicro/base/domain/repository"
	"github.com/xxxmicro/base/domain/repository/gorm"
)

type customFieldRepository struct {
	repository.BaseRepository
}

type CustomFieldRepository interface {
	repository.BaseRepository
}

func NewCustomFieldRepository(db *_gorm.DB) CustomFieldRepository {
	return &customFieldRepository{
		gorm.NewBaseRepository(db),
	}
}