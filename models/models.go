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

// Register setup models profile to init model logic
func Register(profile Profile) {
	switch profile {
	case MemoryProfile:
		registerMemoryProfile()
	default:
		logus.Warn("invalid profile so register default memory profile")
		registerMemoryProfile()
	}
}

func registerMemoryProfile() {
	core.RegisterModelFactory(IdAlbum, func() interface{} {
		return &Album{}
	})
	core.RegisterModelFactory(IdAlbums, func() interface{} {
		return &Albums{}
	})
	core.RegisterRepository(&core.DefaultRepository{
		OnRepository: &memOnRepository{
			template: &memTemplate{},
			Albums:   make(map[string]*Album),
		}})
}
