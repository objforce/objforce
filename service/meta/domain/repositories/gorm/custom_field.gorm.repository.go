package gorm

import (
	"github.com/duolacloud/microbase/domain/repository"
	"github.com/duolacloud/microbase/domain/repository/gorm"
	"github.com/objforce/objforce/service/meta/domain/repositories"
	"go.uber.org/dig"
)

type CustomFieldRepositoryTarget struct {
	dig.In
	DataSourceProvider repository.DataSourceProvider `name:"gorm"`
}

type customFieldRepository struct {
	repository.BaseRepository
}

func NewCustomFieldRepository(target CustomFieldRepositoryTarget) repositories.CustomFieldRepository {
	return &customFieldRepository{
		gorm.NewBaseRepository(target.DataSourceProvider),
	}
}
