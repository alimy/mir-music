package model

import "net/http"

// Album indicate album info
type Album struct {
	Id          int64  `json:"id,omitempty" db:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseYear string `json:"releaseYear" db:"release_year"`
	Genre       string `json:"genre"`
	TrackCount  int    `json:"-" db:"-"`
	AlbumId     int64  `json:"albumId,omitempty" db:"-"`
}

// Albums indicate album slice
type Albums []*Album

// Render implement render.Render(...)
func (*Album) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render implement render.Render(...)
func (Albums) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
