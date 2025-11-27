package repository

import (
	"fmt"

	"github.com/wiwiieie011/songs/models"
	"gorm.io/gorm"
)

type SongRepository interface {
	CreateSong(song *models.Song) error
	GetByID(id uint) (*models.Song, error)
	GetSongsList() ([]models.Song, error)
	GetSongsByCategoryiD(id uint) ([]models.Song, error)
	UpdateSongs(song *models.Song) error
	DeleteSong(id uint) error
}

type songRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) SongRepository {
	return &songRepository{db: db}
}

func (r *songRepository) CreateSong(song *models.Song) error {
	if song == nil {
		return fmt.Errorf("songs is nil")
	}

	return r.db.Create(song).Error
}

func (r *songRepository) GetByID(id uint) (*models.Song, error) {
	var songs models.Song

	if err := r.db.First(&songs, id).Error; err != nil {
		return nil, fmt.Errorf("record not found")
	}

	return &songs, nil
}

func (r *songRepository) GetSongsList() ([]models.Song, error) {
	var songsList []models.Song

	if err := r.db.Find(&songsList).Error; err != nil {
		return nil, fmt.Errorf("not have songs")
	}

	return songsList, nil
}

func (r *songRepository) GetSongsByCategoryiD(id uint) ([]models.Song, error) {
	var songs []models.Song
	if err := r.db.Where("category_id = ?", id).Find(&songs).Error; err != nil {
		return nil, fmt.Errorf("error your links in db")
	}

	return songs, nil
}

func (r *songRepository) UpdateSongs(song *models.Song) error {

	if song == nil {
		return nil
	}

	return r.db.Save(song).Error
}

func (r *songRepository) DeleteSong(id uint) error {

	return r.db.Delete(&models.Song{}, id).Error
}
