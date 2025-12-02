package services

import (
	"fmt"

	"github.com/sirupsen/logrus"
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
	user   repository.UserRepository
	logger *logrus.Logger
}

func NewUserService(user repository.UserRepository, logger *logrus.Logger) UserServices {
	return &userServices{
		user:   user,
		logger: logger,
	}
}

func (r *userServices) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	if req.UserName == "" {
		return nil, fmt.Errorf("user name is nil")
	}

	user := &models.User{
		UserName: req.UserName,
	}

	if err := r.user.CreateUser(user); err != nil {
		r.logger.WithError(err).Error("error CreateUser in services.user function")
		return nil, err
	}

	return user, nil
}

func (r *userServices) GetUserByID(id uint) (*models.User, error) {
	user, err := r.user.GetUserByID(id)
	if err != nil {
		r.logger.WithError(err).Error("error GetUserByID in services.user function")
		return nil, err
	}

	return user, nil
}
func (r *userServices) GetByID(id uint) (*models.User, error) {
	user, err := r.user.GetByID(id)
	if err != nil {
		r.logger.WithError(err).Error("error GetByID in services.user function")
		return nil, err
	}

	return user, nil
}

func (r *userServices) GetAllUsers() ([]models.User, error) {
	list, err := r.user.GetAllUsers()
	if err != nil {
		r.logger.WithError(err).Error("error GetAllUsers in services.user function")
		return nil, err
	}

	return list, nil
}

func (r *userServices) UpdatsUser(id uint, req models.UpdateUserRequest) (*models.User, error) {
	user, err := r.user.GetUserByID(id)
	if err != nil {
		r.logger.WithError(err).Error("error GetUserByID in services.user function")
		return nil, err
	}
	r.apllyUpdate(user, req)

	if err := r.user.Update(user); err != nil {
		r.logger.WithError(err).Error("error UpdateUser in services.user function")
		return nil, err
	}

	return user, nil
}

func (r *userServices) DeleteUser(id uint) error {
	if err := r.user.Delete(id); err != nil {
		r.logger.WithError(err).Error("error DeleteUser in services.user function")
		return fmt.Errorf("record by delete not found")
	}

	return r.user.Delete(id)
}

func (r *userServices) apllyUpdate(user *models.User, req models.UpdateUserRequest) {

	if req.UserName != nil {
		user.UserName = *req.UserName
	}
}
