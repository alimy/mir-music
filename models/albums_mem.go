package models

import (
	"github.com/alimy/mir-music/models/core"
)

var (
	_ core.RecyclableCrud     = &Album{}
	_ core.RecyclableRetrieve = &Albums{}
)

// Create add album to repository
func (m *Album) Create() error {
	return memRepository.addAlbum(m)
}

// Retrieve get album from repository
func (m *Album) Retrieve() error {
	return memRepository.getAlbum(m)
}

// Update update album by id to repository
func (m *Album) Update() error {
	return memRepository.updateAlbum(m)
}

// Delete remove album by id from repository
func (m *Album) Delete() error {
	return memRepository.deleteAlbum(m)
}

// Reset reset album used to recycle to pool
func (m *Album) Reset() {
	m.Id = ""
	m.Title = ""
	m.Artist = ""
	m.ReleaseYear = ""
	m.AlbumId = ""
	m.TrackCount = 0
}

// Retrieve get albums
func (m *Albums) Retrieve() error {
	return memRepository.getAlbums(m)
}

// Reset reset albums
func (m *Albums) Reset() {
	*m = (*m)[:0]
}
