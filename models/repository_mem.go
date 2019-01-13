package models

import (
	"encoding/json"
	"fmt"
	"github.com/alimy/mir-music/models/core"
	"github.com/alimy/mir-music/pkg/assets"
	"github.com/google/uuid"
)

var memRepository *memOnRepository

type memOnRepository struct {
	template    core.Template
	albums      map[string]*Album
	cachedAlums Albums
}

func (r *memOnRepository) OnCreate(model core.RecyclableCreate) error {
	return r.template.Create(model)
}

func (r *memOnRepository) OnRetrieve(model core.RecyclableRetrieve) error {
	return r.template.Retrieve(model)
}

func (r *memOnRepository) OnUpdate(model core.RecyclableUpdate) error {
	return r.template.Update(model)
}

func (r *memOnRepository) OnDelete(model core.RecyclableDelete) error {
	return r.template.Delete(model)
}

func (r *memOnRepository) addAlbum(a *Album) error {
	if id, err := uuid.NewUUID(); err == nil {
		a.Id = id.String()
		r.albums[a.Id] = a
	} else {
		return err
	}
	return nil
}

func (r *memOnRepository) getAlbum(a *Album) error {
	if album, exist := r.albums[a.Id]; exist {
		a.Title = album.Title
		a.Artist = album.Artist
		a.Genre = album.Genre
		a.ReleaseYear = album.ReleaseYear
		a.TrackCount = album.TrackCount
		a.AlbumId = album.AlbumId
	} else {
		return fmt.Errorf("no exist")
	}
	return nil
}

func (r *memOnRepository) updateAlbum(a *Album) error {
	if album, exist := r.albums[a.Id]; exist {
		album.Title = a.Title
		album.Artist = a.Artist
		album.Genre = a.Genre
		album.ReleaseYear = a.ReleaseYear
	} else {
		return fmt.Errorf("no exist")
	}
	return nil
}

func (r *memOnRepository) deleteAlbum(a *Album) error {
	if _, exist := r.albums[a.Id]; exist {
		delete(r.albums, a.Id)
		for i, album := range r.cachedAlums {
			if i == 0 {
				r.cachedAlums = r.cachedAlums[:0]
			}
			if album.Id == a.Id {
				copy(r.cachedAlums[:i], r.cachedAlums[i+1:])
			}
		}
	}
	return nil
}

func (r *memOnRepository) getAlbums(a *Albums) error {
	*a = r.cachedAlums
	return nil
}

func (r *memOnRepository) inflateFrom(name string) error {
	if data, err := assets.Asset(name); err == nil {
		return r.inflate(data)
	} else {
		return err
	}
}

func (r *memOnRepository) inflate(data []byte) error {
	if err := json.Unmarshal(data, &r.cachedAlums); err == nil {
		for _, album := range r.cachedAlums {
			if id, err := uuid.NewUUID(); err == nil {
				album.Id = id.String()
				album.AlbumId = album.Id
				r.albums[album.Id] = album
			} else {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}
