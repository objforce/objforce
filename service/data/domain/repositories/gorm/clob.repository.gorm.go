package gorm

import (
	"github.com/duolacloud/microbase/domain/repository"
	"github.com/duolacloud/microbase/domain/repository/gorm"
	"github.com/objforce/objforce/service/data/domain/repositories"
	"go.uber.org/dig"
)

type ClobRepositoryTarget struct {
	dig.In
	DataSourceProvider repository.DataSourceProvider `name:"gorm"`
}

type clobRepository struct {
	repository.BaseRepository
}

func NewClobRepository(target ClobRepositoryTarget) repositories.ClobRepository {
	return &clobRepository{
		gorm.NewBaseRepository(target.DataSourceProvider),
	}
}
