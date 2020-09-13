package repositories

import(
	"github.com/jinzhu/gorm"
	"github.com/objforce/objforce/meta-srv/app/domain/models"
)

type extendTableMetaRepository struct {
	db *gorm.DB
}

type ExtendTableMetaRepository interface {
	All() ([]*models.MTExtendTableMeta, error)
}

func NewExtendTableMetaRepository(db *gorm.DB) ExtendTableMetaRepository {
	return &extendTableMetaRepository{
		db,
	}
}

func (r *extendTableMetaRepository) All() ([]*models.MTExtendTableMeta, error) {
	items := []*models.MTExtendTableMeta{}

	if err := r.db.Model(&models.MTExtendTableMeta{}).Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}