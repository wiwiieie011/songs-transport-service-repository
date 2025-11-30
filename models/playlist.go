package models

import "gorm.io/gorm"

	type PlayList struct {
    gorm.Model
    Name   string       `json:"name"`
    UserID uint         `json:"user_id"`
    User   *User        `json:"-" gorm:"constraint:OnDelete:CASCADE"`  // связь к родителю
    Items  []PlayListItems `json:"items" gorm:"constraint:OnDelete:CASCADE"`  // все песни плейлиста удаляются вместе с ним
}


type CreatePlayListRequest struct{
	Name string `json:"name" binding:"required"`
	UserID uint `json:"user_id" binding:"required"`
}

type UpdatePlayListRequest struct{
	Name *string `json:"name"`
}