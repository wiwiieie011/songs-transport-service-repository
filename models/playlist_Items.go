package models


type PlayListItems struct {

	SongID uint `json:"song_id" gorm:"primaryKey"`
	Song *Song `json:"song"`
	
	PlayListID uint	`json:"playlist_id" gorm:"primaryKey"`
	PlayList *PlayList `json:"-"`

}


type CreatePlayListItemsRequest struct{
	SongID uint `json:"song_id"  binding:"required"`
	PlayListID uint	`json:"playlist_id"  binding:"required"`
}

type UpdatePlayListItemsRequest struct{
	SongID *uint `json:"song_id"`
	PlayListID  *uint	`json:"playlist_id"`
}