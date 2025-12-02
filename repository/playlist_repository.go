package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/wiwiieie011/songs/models"
	"gorm.io/gorm"
)

type PlayListRepository interface {
	Create(playlist *models.PlayList) error
	GetAll() ([]models.PlayList, error)
	GetByID(id uint) (*models.PlayList, error)
	Update(playlist *models.PlayList) error
	Delete(id uint) error
}

type playListRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewPlayListRepository(db *gorm.DB, logger *logrus.Logger) PlayListRepository {
	return &playListRepository{
		db:     db,
		logger: logger,
	}

}

func (r *playListRepository) Create(playlist *models.PlayList) error {
	if playlist == nil {
		return fmt.Errorf("error create playlist")
	}

	return r.db.Create(playlist).Error
}

func (r *playListRepository) GetAll() ([]models.PlayList, error) {
	var playlists []models.PlayList
	if err := r.db.Find(&playlists).Error; err != nil {
		r.logger.WithError(err).Error("error  GetAll repository.playlist function")
		return nil, fmt.Errorf("record ont ound")
	}
	return playlists, nil
}

func (r *playListRepository) GetByID(id uint) (*models.PlayList, error) {
	var playlist models.PlayList
	if err := r.db.Preload("Items.Song").First(&playlist, id).Error; err != nil {
		r.logger.WithError(err).Error("error  GetByID repository.playlist function")
		return nil, fmt.Errorf("not found playlist ID")
	}

	return &playlist, nil
}

func (r *playListRepository) Update(playlist *models.PlayList) error {

	if playlist == nil {
		return fmt.Errorf("error update")
	}

	return r.db.Save(playlist).Error
}

func (r *playListRepository) Delete(id uint) error {
	return r.db.Delete(&models.PlayList{}, id).Error
}
