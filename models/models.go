package models

import (
	"github.com/alimy/mir-music/models/core"
	"github.com/unisx/logus"
)

// Profile indicate model config
type Profile int

// Profiles items
const (
	MemoryProfile Profile = iota
)

const (
	// used initial albums data from external json file
	fackJsonFileName = "albums.json"
)

// Register setup models profile to init model logic
func Register(profile Profile) error {
	var err error
	switch profile {
	case MemoryProfile:
		err = registerMemoryProfile()
	default:
		logus.Warn("invalid profile so register default memory profile")
		err = registerMemoryProfile()
	}
	return err
}

func registerMemoryProfile() error {
	// setup Album mode factory
	core.RegisterModelFactory(IdAlbum, func() interface{} {
		return &Album{}
	})

	// setup Albums mode factory
	core.RegisterModelFactory(IdAlbums, func() interface{} {
		return &Albums{}
	})

	// initial Repository data
	memRepository = &memOnRepository{
		template:    &memTemplate{},
		albums:      make(map[string]*Album),
		cachedAlums: make(Albums, 0),
	}
	if err := memRepository.inflateFrom(fackJsonFileName); err == nil {
		// setup Repository
		core.RegisterRepository(&core.DefaultRepository{
			OnRepository: memRepository,
		})
	} else {
		return err
	}
	return nil
}
