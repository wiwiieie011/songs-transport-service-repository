package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model

	SongName  string `json:"song_name"`
	Author    string `json:"author"`
	GroupName string `json:"group_name"`

	CategoryID uint      `json:"category_id" gorm:"index"`
	Category   *Category `json:"-"`
}

type CreateSongRequest struct {
	SongName   string `json:"song_name"`
	Author     string `json:"author"`
	GroupName  string `json:"group_name"`
	CategoryID uint   `json:"category_id"`
}

type UpdateSongRequest struct {
	SongName   *string `json:"song_name"`
	Author     *string `json:"author"`
	GroupName  *string `json:"group_name"`
	CategoryID *uint   `json:"category_id"`
}
