package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `json:"name"`
	Songs []Song `json:"-"`
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	Name *string
}
