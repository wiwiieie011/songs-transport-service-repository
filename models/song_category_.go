package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string
	SongID uint

	Song *Song
	List []Song `json:"list"`
}
