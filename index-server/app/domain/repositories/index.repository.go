package repositories

import(
	_gorm "github.com/jinzhu/gorm"
	"github.com/xxxmicro/base/domain/repository"
	"github.com/xxxmicro/base/domain/repository/gorm"
)

type indexRepository struct {
	repository.BaseRepository
}

type IndexRepository interface {
	repository.BaseRepository
}

func NewIndexRepository(db *_gorm.DB) IndexRepository {
	return &indexRepository{
		gorm.NewBaseRepository(db),
	}
}