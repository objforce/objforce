package repositories

import(
	_gorm "github.com/jinzhu/gorm"
	"github.com/xxxmicro/base/domain/repository"
	"github.com/xxxmicro/base/domain/repository/gorm"
)

type customObjectRepository struct {
	repository.BaseRepository
}

type CustomObjectRepository interface {
	repository.BaseRepository
}

func NewCustomObjectRepository(db *_gorm.DB) CustomObjectRepository {
	return &customObjectRepository{
		gorm.NewBaseRepository(db),
	}
}