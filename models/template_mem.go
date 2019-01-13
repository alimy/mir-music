package models

import "github.com/alimy/mir-music/models/core"

type memTemplate struct{}

func (t *memTemplate) Create(model core.RecyclableCreate) error {
	return model.Create()
}

func (t *memTemplate) Retrieve(model core.RecyclableRetrieve) error {
	return model.Retrieve()
}

func (t *memTemplate) Update(model core.RecyclableUpdate) error {
	return model.Update()
}

func (t *memTemplate) Delete(model core.RecyclableDelete) error {
	return model.Delete()
}
