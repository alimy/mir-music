package models

import (
	"github.com/alimy/mir-music/models/core"
	"github.com/pkg/errors"
	"github.com/unisx/logus"
)

var (
	_ core.RecyclableCrud = &Album{}
	_ core.RecyclableCrud = &Albums{}
)

func (m *Album) Create() error {
	// TODO
	logus.Debug("album create")
	return errors.New("invalide operator")
}

func (m *Album) Retrieve() error {
	// TODO
	logus.Debug("album retrieve")
	return errors.New("invalide operator")
}

func (m *Album) Update() error {
	// TODO
	logus.Debug("album update")
	return errors.New("invalide operator")
}

func (m *Album) Delete() error {
	// TODO
	logus.Debug("album delete")
	return errors.New("invalide operator")
}

func (m *Album) Reset() {
	logus.Debug("album reset")
	// TODO
}

func (m *Albums) Create() error {
	// TODO
	logus.Debug("albums create")
	return errors.New("invalide operator")
}

func (m *Albums) Retrieve() error {
	// TODO
	logus.Debug("albums retrieve")
	return errors.New("invalide operator")
}

func (m *Albums) Update() error {
	// TODO
	logus.Debug("albums update")
	return errors.New("invalide operator")
}

func (m *Albums) Delete() error {
	// TODO
	logus.Debug("albums delete")
	return errors.New("invalide operator")
}

func (m *Albums) Reset() {
	logus.Debug("albums reset")
	// TODO
}
