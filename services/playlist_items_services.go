package services

import (
	"github.com/sirupsen/logrus"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/repository"
)

type PlayListItemsServices interface {
	CreatePlayListItems(req models.CreatePlayListItemsRequest) (*models.PlayListItems, error)
	GetByID(pid, uid uint) (*models.PlayListItems, error)
	Delete(pid, sid uint) error
}

type playListItemsServices struct {
	pli      repository.PlayListItemsRepository
	playlist PlayListServices
	song     SongService
	logger   *logrus.Logger
}

func NewPlayListItemsServices(
	pli repository.PlayListItemsRepository,
	playlist PlayListServices,
	song SongService,
	logger *logrus.Logger,
) PlayListItemsServices {
	return &playListItemsServices{
		pli:      pli,
		playlist: playlist,
		song:     song,
		logger:   logger,
	}
}

func (r *playListItemsServices) CreatePlayListItems(req models.CreatePlayListItemsRequest) (*models.PlayListItems, error) {

	if _, err := r.playlist.GetPlaylistByID(req.PlayListID); err != nil {
		r.logger.WithError(err).Error("error CreatePlayListItems in services.playlist_items function")
		return nil, err
	}

	if _, err := r.song.GetSongByID(req.SongID); err != nil {
		r.logger.WithError(err).Error("error CreatePlayListItems in services.playlist_items function")
		return nil, err
	}

	playListItems := &models.PlayListItems{
		SongID:     req.SongID,
		PlayListID: req.PlayListID,
	}

	if err := r.pli.Create(playListItems); err != nil {
		r.logger.WithError(err).Error("error CreatePlayListItems in services.playlist_items function")
		return nil, err
	}

	return playListItems, nil
}

func (r *playListItemsServices) GetByID(pid, sid uint) (*models.PlayListItems, error) {
	playListItems, err := r.pli.GetByID(pid, sid)
	if err != nil {
		r.logger.WithError(err).Error("error GetByID in services.playlist_items function")

		return nil, err
	}

	return playListItems, nil
}

func (r *playListItemsServices) Delete(pid, sid uint) error {
	if err := r.pli.Delete(pid, sid); err != nil {
		r.logger.WithError(err).Error("error Delete in services.playlist_items function")
		return err
	}
	return nil
}
