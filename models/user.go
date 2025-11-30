package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    UserName  string      `json:"user_name"`
    Playlists []PlayList  `json:"playlists" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}



type CreateUserRequest struct{
	UserName string `json:"user_name" binding:"required"`
}


type UpdateUserRequest struct{
	UserName *string `json:"user_name"`
}
