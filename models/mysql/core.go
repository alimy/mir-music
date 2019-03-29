package mysql

import (
	"fmt"
	"github.com/alimy/mir-music/models/model"
	"github.com/unisx/logus"
)

const (
	sqlGetMainPage     = "SELECT * FROM albums limit 50"
	sqlAddAlbum        = "INSERT INTO albums (artist, title, release_year, genre) VALUES (?, ?, ?, ?)"
	sqlUpdateAlbum     = "UPDATE albums SET artist=:artist, title=:title, release_year=:release_year, genre=:genre WHERE id=:id"
	sqlGetAlbumById    = "SELECT * FROM albums WHERE id=?"
	sqlDeleteAlbumById = "DELETE FROM albums WHERE id=?"
)

// GetMainPage get main page data
func (r *mysqlRepository) GetMainPage() (interface{}, error) {
	albums := model.Albums{}
	err := r.Select(&albums, sqlGetMainPage)
	if err != nil {
		return nil, err
	}
	for _, album := range albums {
		album.AlbumId = album.Id
	}
	return albums, nil
}

func (r *mysqlRepository) AddAlbum(album *model.Album) error {
	if album == nil {
		return fmt.Errorf("album is nil")
	}
	logus.Debug("addAlbum", logus.Any("album", album))
	_, err := r.Exec(sqlAddAlbum, album.Artist, album.Title, album.ReleaseYear, album.Genre)
	return err

}

func (r *mysqlRepository) UpdateAlbum(album *model.Album) error {
	if album == nil {
		return fmt.Errorf("album is nil")
	}
	album.Id = album.AlbumId
	_, err := r.NamedExec(sqlUpdateAlbum, album)
	return err
}

func (r *mysqlRepository) GetAlbumById(id int64) (*model.Album, error) {
	album := &model.Album{}
	err := r.Get(album, sqlGetAlbumById, id)
	if err != nil {
		return nil, err
	}
	album.AlbumId = album.Id
	return album, nil
}

func (r *mysqlRepository) DeleteAlbumById(id int64) error {
	_, err := r.Exec(sqlDeleteAlbumById, id)
	return err
}

func (r *mysqlRepository) GetProfiles() *model.AppInfo {
	return &model.AppInfo{
		Info: &model.Profiles{Profile: "Mir music info", Services: "provide music info search"},
	}
}
