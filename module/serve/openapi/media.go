package openapi

import (
	"github.com/alimy/mir"
	"github.com/alimy/mir-music/models"
	"github.com/alimy/mir-music/models/core"
	"github.com/unisx/logus"
	"net/http"
)

type media struct {
	group            mir.Group  `mir:"v1"`
	getAlbums        mir.Get    `mir:"/albums/"`
	createAlbums     mir.Put    `mir:"/albums/"`
	updateAlbums     mir.Post   `mir:"/albums/"`
	getAlbumsById    mir.Get    `mir:"/albums/:albumId/"`
	deleteAlbumsById mir.Delete `mir:"/albums/:albumId/"`
}

// GetAlbums GET handler of "/albums/"
func (m *media) GetAlbums(c Context) {
	if albums, ok := core.Model(models.IdAlbums).(*models.Albums); ok {
		logus.Debug("getAlbums")
		core.Retrieve(models.IdAlbums, albums)
		c.JSON(http.StatusOK, albums)
	} else {
		c.String(http.StatusNotFound, "get albums")
	}
}

// CreateAlbums PUT handler of "/albums/"
func (m *media) CreateAlbums(c Context) {
	// TODO
	logus.Debug("create albums")
	c.String(http.StatusCreated, "Albums item created")
}

// UpdateAlbums POST handler of "/albums/"
func (m *media) UpdateAlbums(c Context) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("updateAlbums")
		album.Id = c.Param("albumId")
		core.Update(models.IdAlbum, album)
		c.String(http.StatusCreated, "Albums item updated")
	} else {
		c.String(http.StatusNotFound, "update albums failure")
	}
}

// GetAlbumsById GET handler of "/albums/:albumId/"
func (m *media) GetAlbumsById(c Context) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("getAlbumsById")
		album.Id = c.Param("albumId")
		core.Retrieve(models.IdAlbum, album)
		c.JSON(http.StatusOK, album)
	} else {
		c.String(http.StatusNotFound, "update albums failure")
	}
}

// DeleteAlbumsById DELETE handler of "/albums/:albumId/"
func (m *media) DeleteAlbumsById(c Context) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("deleteAlbumsById")
		album.Id = c.Param("albumId")
		core.Delete(models.IdAlbum, album)
		c.String(http.StatusOK, "Albums item deleted")
	} else {
		c.String(http.StatusNotFound, "delete album failure")
	}
}
