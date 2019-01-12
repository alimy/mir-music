package models

type Album struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseYear string `json:"releaseYear"`
	Genre       string `json:"genre"`
	TrackCount  int    `json:"-"`
	AlbumId     string `json:"albumId"`
}

type Albums struct {
	AlbumSlice []Album
}
