package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/wiwiieie011/songs/models"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint) error
}

type categoryRepo struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewCategoryRepository(db *gorm.DB, logger *logrus.Logger) CategoryRepo {
	return &categoryRepo{
		db:     db,
		logger: logger,
	}
}

func (r *categoryRepo) Create(category *models.Category) error {
	if category == nil {
		r.logger.Error("error in Create repository.category db")
		return fmt.Errorf("error create category")
	}

	return r.db.Create(category).Error
}

func (r *categoryRepo) GetByID(id uint) (*models.Category, error) {
	var category models.Category

	if err := r.db.Preload("Songs").First(&category, id).Error; err != nil {
		r.logger.WithError(err).Error("error  GetByID repository.category function")
		return nil, fmt.Errorf("record not found")
	}
	return &category, nil
}

func (r *categoryRepo) GetAll() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Find(&categories).Error; err != nil {
		r.logger.WithError(err).Error("error in GetAll repository.category function")
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepo) UpdateCategory(category *models.Category) error {
	if category == nil {
		r.logger.Error("error in UpdateCategory repository.category function")
		return nil
	}

	return r.db.Save(category).Error
}

func (r *categoryRepo) DeleteCategory(id uint) error {
	if err := r.db.Delete(&models.Category{}, id).Error; err != nil {
		r.logger.WithError(err).Error("error in Delete repository.category function")
		return fmt.Errorf("not have deleted category")
	}
	return nil
}
