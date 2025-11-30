package services

import (
	"fmt"

	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/repository"
)

type PlayListServices interface {
	 CreatePlayList(req models.CreatePlayListRequest) (*models.PlayList, error)
	GetAllPlaylists() ([]models.PlayList, error)
	GetPlaylistByID(id uint) (*models.PlayList, error)
	UpdatePlaylistByID(id uint, req models.UpdatePlayListRequest) (*models.PlayList, error)
	DeletePlayList(id uint) error
}

type playListServices struct {
	playlist repository.PlayListRepository
	us       UserServices
}

func NewPlayListServices(
	playlist repository.PlayListRepository,
	us UserServices,
) PlayListServices {
	return &playListServices{
		playlist: playlist,
		us:       us,
	}
}


func (r *playListServices) CreatePlayList(req models.CreatePlayListRequest) (*models.PlayList, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("empty playlist name")
	}

	_, err := r.us.GetByID(req.UserID)
	if err != nil {
		return nil, err
	}

	playlist := &models.PlayList{
		Name:   req.Name,
		UserID: req.UserID,
	}

	if err := r.playlist.Create(playlist); err != nil {
		return nil, err
	}

	return playlist, nil
}

func (r *playListServices) GetAllPlaylists() ([]models.PlayList, error) {
	playlists, err := r.playlist.GetAll()
	if err != nil {
		return nil, err
	}
	return playlists, nil
}

func (r *playListServices) GetPlaylistByID(id uint) (*models.PlayList, error) {
	playlist, err := r.playlist.GetByID(id)
	if err != nil {
		return nil, err
	}

	return playlist, nil

}

func (r *playListServices) UpdatePlaylistByID(id uint, req models.UpdatePlayListRequest) (*models.PlayList, error) {
	playlist, err := r.playlist.GetByID(id)
	if err != nil {
		return nil, err
	}

	r.applyPlaylist(playlist, req)

	if err := r.playlist.Update(playlist); err != nil {
		return nil, err
	}

	return playlist, nil
}

func (r *playListServices) DeletePlayList(id uint) error {
	err := r.playlist.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *playListServices) applyPlaylist(playlist *models.PlayList, req models.UpdatePlayListRequest) {
	if req.Name != nil {
		playlist.Name = *req.Name
	}
}
