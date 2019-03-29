package core

import "github.com/alimy/mir-music/models/model"

// ProfileAction indicate profiles action
type ProfileAction interface {
	GetProfiles() *model.AppInfo
}
