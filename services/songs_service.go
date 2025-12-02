package services

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/repository"
)

type SongService interface {
	CreateSong(req models.CreateSongRequest) (*models.Song, error)
	GetSongByID(id uint) (*models.Song, error)
	GetSongs() ([]models.Song, error)
	GetSongsByCategoryiD(id uint) ([]models.Song, error)
	UpdateSong(id uint, req models.UpdateSongRequest) (*models.Song, error)
	DeleteSong(id uint) error
}

type songService struct {
	song   repository.SongRepository
	logger *logrus.Logger
}

func NewSongService(song repository.SongRepository, logger *logrus.Logger) SongService {
	return &songService{
		song:   song,
		logger: logger,
	}
}

func (r *songService) CreateSong(req models.CreateSongRequest) (*models.Song, error) {
	if err := r.validateSong(req); err != nil {
		return nil, err
	}

	song := &models.Song{
		SongName:   req.SongName,
		Author:     req.Author,
		GroupName:  req.GroupName,
		CategoryID: req.CategoryID,
	}

	if err := r.song.CreateSong(song); err != nil {
		r.logger.WithError(err).Error("error CreateSong in services.songs function")
		return nil, fmt.Errorf("error create song")
	}

	return song, nil
}

func (r *songService) GetSongByID(id uint) (*models.Song, error) {
	song, err := r.song.GetByID(id)
	if err != nil {
		r.logger.WithError(err).Error("error GetSongByID in services.songs function")
		return nil, fmt.Errorf("record not found by id")
	}

	return song, nil
}

func (r *songService) GetSongs() ([]models.Song, error) {
	list, err := r.song.GetSongsList()
	if err != nil {
		r.logger.WithError(err).Error("error GetSongByID in services.songs function")
		return nil, fmt.Errorf("record not found")
	}

	return list, nil
}

func (r *songService) GetSongsByCategoryiD(id uint) ([]models.Song, error) {
	songs, err := r.song.GetSongsByCategoryiD(id)
	if err != nil {
		r.logger.WithError(err).Error("error GetSongByCategoryID in services.songs function")
		return nil, fmt.Errorf("not found song category")
	}

	return songs, nil
}

func (r *songService) UpdateSong(id uint, req models.UpdateSongRequest) (*models.Song, error) {
	song, err := r.song.GetByID(id)
	if err != nil {
		r.logger.WithError(err).Error("error UpdateSong in services.songs function")
		return nil, fmt.Errorf("record not found by id")
	}

	r.applySongUpdate(song, req)

	if err := r.song.UpdateSongs(song); err != nil {
		r.logger.WithError(err).Error("error UpdateSong in services.songs function")
		return nil, fmt.Errorf("failed update song by id")
	}

	return song, nil
}

func (r *songService) DeleteSong(id uint) error {
	if _, err := r.song.GetByID(id); err != nil {
		r.logger.WithError(err).Error("error UpdateSong in services.songs function")
		return fmt.Errorf("record by delete not found")
	}

	return r.song.DeleteSong(id)
}

func (r *songService) validateSong(req models.CreateSongRequest) error {

	if req.SongName == "" {
		return fmt.Errorf("поле song_name не должно быть пустым")
	}

	if req.Author == "" {
		return fmt.Errorf("поле author не должно быть пустым")
	}

	if req.GroupName == "" {
		return fmt.Errorf("поле group_name не должно быть пустым")
	}

	return nil
}

func (r *songService) applySongUpdate(song *models.Song, req models.UpdateSongRequest) {

	if req.SongName != nil {
		song.SongName = *req.SongName
	}

	if req.Author != nil {
		song.Author = *req.Author
	}

	if req.GroupName != nil {
		song.GroupName = *req.GroupName
	}

	if req.CategoryID != nil {
		song.CategoryID = *req.CategoryID
	}
}
