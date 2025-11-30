package services

import (
	"fmt"

	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/repository"
)

type UserServices interface {
	CreateUser(req models.CreateUserRequest) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdatsUser(id uint, req models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id uint) error
}

type userServices struct {
	user repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserServices {
	return &userServices{user: user}
}

func (r *userServices) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	if req.UserName == "" {
		return nil, fmt.Errorf("user name is nil")
	}

	user := &models.User{
		UserName: req.UserName,
	}

	if err := r.user.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userServices) GetUserByID(id uint) (*models.User, error) {
	user, err := r.user.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userServices) GetByID(id uint) (*models.User, error) {
	user, err := r.user.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userServices) GetAllUsers() ([]models.User, error) {
	list, err := r.user.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *userServices) UpdatsUser(id uint, req models.UpdateUserRequest) (*models.User, error) {
	user, err := r.user.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	r.apllyUpdate(user, req)

	if err := r.user.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userServices) DeleteUser(id uint) error {
	if _, err := r.user.GetUserByID(id); err != nil {
		return fmt.Errorf("record by delete not found")
	}

	return r.user.Delete(id)
}

func (r *userServices) apllyUpdate(user *models.User, req models.UpdateUserRequest) {

	if req.UserName != nil {
		user.UserName = *req.UserName
	}
}
