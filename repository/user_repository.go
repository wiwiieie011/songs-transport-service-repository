package repository

import (
	"fmt"

	"github.com/wiwiieie011/songs/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	 CreateUser(user *models.User) error
	 GetAllUsers() ([]models.User,error)
	 GetByID(id uint) (*models.User, error)
	 GetUserByID(id uint) (*models.User, error)
	 Update(user *models.User) error
	 Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB)	UserRepository{
	return  &userRepository{db: db}
}


func (r *userRepository) CreateUser(user *models.User) error{
	if user == nil {
		return fmt.Errorf("user is null")
	}

	return r.db.Create(user).Error
}

func (r *userRepository) GetAllUsers() ([]models.User,error){
	var users []models.User
	if err:= r.db.Find(&users).Error; err!=nil{
		return nil, fmt.Errorf("record ont ound")
	}
	return  users,nil

}

func (r *userRepository) GetUserByID(id uint) (*models.User, error){
		var user models.User
		if err:= r.db.Preload("Playlists.Items.Song").First(&user, id).Error; err !=nil{
			return nil, fmt.Errorf("not found")
		}

		return  &user, nil
}

func (r *userRepository) GetByID(id uint) (*models.User, error){
		var user models.User
		if err:= r.db.First(&user, id).Error; err !=nil{
			return nil, fmt.Errorf("not found")
		}

		return  &user, nil
}



func (r *userRepository) Update(user *models.User) error{
	if user == nil{
		return  fmt.Errorf("error update")
	}

	return r.db.Save(user).Error
}


func (r *userRepository) Delete(id uint) error{
	return r.db.Delete(&models.User{}, id).Error
}