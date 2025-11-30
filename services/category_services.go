package services

import (
	"fmt"

	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/repository"
)

type CategoryService interface {
	CreateCategory(req models.CreateCategoryRequest) (*models.Category, error)
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
	UpdateCategory(id uint, req models.UpdateCategoryRequest) (*models.Category, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	category repository.CategoryRepo
}

func NewCategoryService(category repository.CategoryRepo) CategoryService {
	return &categoryService{
		category: category,
	}
}

func (r *categoryService) CreateCategory(req models.CreateCategoryRequest) (*models.Category, error) {

	if req.Name == "" {
		return nil, fmt.Errorf("empty category name")
	}

	category := &models.Category{
		Name: req.Name,
	}

	if err := r.category.Create(category); err != nil {
		return nil, fmt.Errorf("error create category")
	}

	return category, nil
}

func (r *categoryService) GetAll() ([]models.Category, error) {
	list, err := r.category.GetAll()
	if err != nil {
		return nil, fmt.Errorf("no found category list")
	}
	return list, nil
}

func (r *categoryService) GetByID(id uint) (*models.Category, error) {
	category, err := r.category.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid category id")
	}

	return category, nil
}

func (r *categoryService) UpdateCategory(id uint, req models.UpdateCategoryRequest) (*models.Category, error) {
	category, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	r.applyCategory(category, req)

	if err := r.category.UpdateCategory(category); err != nil {
		return nil, fmt.Errorf("failed update song by id")
	}

	return category, nil
}

func (r *categoryService) DeleteCategory(id uint) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}

	return r.category.DeleteCategory(id)
}

func (r *categoryService) applyCategory(category *models.Category, req models.UpdateCategoryRequest) {
	if req.Name != nil {
		category.Name = *req.Name
	}
}
