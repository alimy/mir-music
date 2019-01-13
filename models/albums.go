package models

// Album indicate album info
type Album struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseYear string `json:"releaseYear"`
	Genre       string `json:"genre"`
	TrackCount  int    `json:"-"`
	AlbumId     string `json:"albumId,omitempty"`
}

// Albums indicate album slice
type Albums []*Album
