package repository

import (
	"fmt"

	"github.com/wiwiieie011/songs/models"
	"gorm.io/gorm"
)

type PlayListItemsRepository interface {
	Create(items *models.PlayListItems) error
	GetByID(pid, sid uint) (*models.PlayListItems, error)
	Delete(pid, sid uint) error
}

type playListItemsRepository struct {
	db *gorm.DB
}

func NewPlayListItemsRepository(db *gorm.DB) PlayListItemsRepository {
	return &playListItemsRepository{db: db}
}

func (r *playListItemsRepository) Create(items *models.PlayListItems) error {
	if items == nil {
		return fmt.Errorf("playlist items is nil")
	}

	return r.db.Create(items).Error
}

func (r *playListItemsRepository) GetByID(pid, sid uint) (*models.PlayListItems, error) {
	var playlist models.PlayListItems
	if err := r.db.Where("playlist_id = ? AND song_id = ?", pid, sid).First(&playlist).Error; err != nil {
		return nil, fmt.Errorf("not found")
	}
	return &playlist, nil
}

func (r *playListItemsRepository) Delete(pid, sid uint) error {
	return r.db.Where("playlist_id = ? AND song_id = ?", pid, sid).Delete(&models.PlayListItems{}).Error
}
