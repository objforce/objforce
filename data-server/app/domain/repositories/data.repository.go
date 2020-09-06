package repositories

import(
	_gorm "github.com/jinzhu/gorm"
	"github.com/xxxmicro/base/domain/repository"
	"github.com/xxxmicro/base/domain/repository/gorm"
)

type DataRepository interface {
	repository.BaseRepository
}

type dataRepository struct {
	repository.BaseRepository
}

func NewDataRepository(db *_gorm.DB) DataRepository {
	return &dataRepository{
		gorm.NewBaseRepository(db),
	}
}