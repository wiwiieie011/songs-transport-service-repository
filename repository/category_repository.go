package repository

import (
	"fmt"

	"github.com/wiwiieie011/songs/models"
	"gorm.io/gorm"
)

type CategoryRepo interface {
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) Create(category *models.Category) error {
	if category == nil {
		// return fmt.Errorf("error create category")
		return nil
	}

	return r.db.Create(category).Error
}

func (r *categoryRepo) GetByID(id uint) (*models.Category, error) {
	var category models.Category

	if err := r.db.First(&category, id).Error; err != nil {
		return nil, fmt.Errorf("record not found")
	}

	return &category, nil
}

