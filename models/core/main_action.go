package core

import (
	"github.com/alimy/mir-music/models/model"
)

// MainAction indicator mina service interface
type MainAction interface {
	GetMainPage() (interface{}, error)
	AddAlbum(album *model.Album) error
	UpdateAlbum(album *model.Album) error
	GetAlbumById(id int64) (*model.Album, error)
	DeleteAlbumById(id int64) error
}
