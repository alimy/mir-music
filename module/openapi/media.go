package openapi

import (
	"github.com/alimy/mir"
	"github.com/alimy/mir-music/models/core"
	"github.com/alimy/mir-music/models/model"
	"github.com/alimy/mir-music/pkg/json"
	"github.com/unisx/logus"
	"net/http"
	"strconv"
)

type Media struct {
	group            mir.Group  `mir:"v1"`
	getAlbums        mir.Get    `mir:"/albums/"`
	createAlbums     mir.Put    `mir:"/albums/"`
	updateAlbums     mir.Post   `mir:"/albums/"`
	getAlbumsById    mir.Get    `mir:"/albums/:albumId/"`
	deleteAlbumsById mir.Delete `mir:"/albums/:albumId/"`

	*core.Context
}

// GetAlbums GET handler of "/albums/"
func (m *Media) GetAlbums(c Context) {
	m.Retrieve(c, core.RdsMainPage, m.Repo.GetMainPage)
}

// CreateAlbums PUT handler of "/albums/"
func (m *Media) CreateAlbums(c Context) {
	defer c.Request.Body.Close()

	album := &model.Album{}
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(album)

	if err := m.Repo.AddAlbum(album); err != nil {
		m.ErrInternalServer(c, err.Error())
		logus.Debug("create albums failure", logus.ErrorField(err))
		return
	}
	m.Expire(core.RdsMainPage)
	logus.Debug("create albums success")
	c.String(http.StatusCreated, "Albums item created")
}

// UpdateAlbums POST handler of "/albums/"
func (m *Media) UpdateAlbums(c Context) {
	defer c.Request.Body.Close()

	album := &model.Album{}
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(album)

	if album.Id == 0 {
		if err := m.Repo.AddAlbum(album); err != nil {
			m.ErrInternalServer(c, err.Error())
			logus.Debug("create albums failure", logus.ErrorField(err))
			return
		}
		m.Expire(core.RdsMainPage)
		logus.Debug("create albums success")
		c.String(http.StatusCreated, "albums item created")
		return
	}

	if err := m.Repo.UpdateAlbum(album); err != nil {
		m.ErrInternalServer(c, err.Error())
		logus.Debug("update albums failure", logus.ErrorField(err))
		return
	}
	m.Expire(core.RdsMainPage)
	logus.Debug("update albums success")
	c.String(http.StatusCreated, "albums item created")
}

// GetAlbumsById GET handler of "/albums/:albumId/"
func (m *Media) GetAlbumsById(c Context) {
	albumId := c.Param("albumId")
	id, err := strconv.ParseInt(albumId, 10, 0)
	if err != nil {
		m.ErrInternalServer(c, err.Error())
		logus.Debug("parse id failure", logus.ErrorField(err))
		return
	}
	album, err := m.Repo.GetAlbumById(id)
	if err != nil {
		m.ErrInternalServer(c, err.Error())
		logus.Debug("get album by id failure", logus.Int64("id", id), logus.ErrorField(err))
		return
	}
	c.JSON(http.StatusOK, album)
	logus.Debug("get albums by id success", logus.Int64("id", id))
}

// DeleteAlbumsById DELETE handler of "/albums/:albumId/"
func (m *Media) DeleteAlbumsById(c Context) {
	albumId := c.Param("albumId")
	id, err := strconv.ParseInt(albumId, 10, 0)
	if err != nil {
		m.ErrInternalServer(c, err.Error())
		logus.Debug("parse id failure", logus.ErrorField(err))
		return
	}
	err = m.Repo.DeleteAlbumById(id)
	if err != nil {
		m.ErrInternalServer(c, err.Error())
		logus.Debug("delete album failure", logus.Int64("id", id), logus.ErrorField(err))
		return
	}
	m.Expire(core.RdsMainPage)
	logus.Debug("delete album by id success", logus.Int64("id", id))
	c.String(http.StatusOK, "album deleted")
}
