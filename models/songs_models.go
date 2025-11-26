package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	
	SongName  string
	Author    string
	GroupName string

	CategoryID uint	
	Category *Category
}

type CreateSongRequest struct {
	SongName  string
	Author    string
	GroupName string
	CategoryID Category
}

type UpdateSongRequest struct {
	SongName  *string
	Author    *string
	GroupName *string
	CategoryID uint	
}
