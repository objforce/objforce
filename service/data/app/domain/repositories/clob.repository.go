package repositories

import (
	_gorm "github.com/jinzhu/gorm"
	"github.com/xxxmicro/base/domain/repository"
	"github.com/xxxmicro/base/domain/repository/gorm"
)

type ClobRepository interface {
	repository.BaseRepository
}

type clobRepository struct {
	repository.BaseRepository
}

func NewClobRepository(db *_gorm.DB) ClobRepository {
	return &clobRepository{
		gorm.NewBaseRepository(db),
	}
}
