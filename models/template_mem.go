package models

import "github.com/alimy/mir-music/models/core"

type memTemplate struct{}

func (t *memTemplate) Create(model core.RecyclableCrud) error {
	return model.Create()
}

func (t *memTemplate) Retrieve(model core.RecyclableCrud) error {
	return model.Retrieve()
}

func (t *memTemplate) Update(model core.RecyclableCrud) error {
	return model.Update()
}

func (t *memTemplate) Delete(model core.RecyclableCrud) error {
	return model.Delete()
}
