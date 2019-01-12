package core

import (
	"errors"
	"github.com/unisx/logus"
	"sync"
)

var (
	modelPools = make(map[int]*sync.Pool)
	repository Repository
)

// Repository data operator
type Repository interface {
	Create(int, RecyclableCrud) error
	Retrieve(int, RecyclableCrud) error
	Update(int, RecyclableCrud) error
	Delete(int, RecyclableCrud) error
}

// OnRepository data operator
type OnRepository interface {
	OnCreate(RecyclableCrud) error
	OnRetrieve(RecyclableCrud) error
	OnUpdate(RecyclableCrud) error
	OnDelete(RecyclableCrud) error
}

// Template mode data template operator
type Template interface {
	Create(RecyclableCrud) error
	Retrieve(RecyclableCrud) error
	Update(RecyclableCrud) error
	Delete(RecyclableCrud) error
}

// Crud mode data operator
type Crud interface {
	Create() error
	Retrieve() error
	Update() error
	Delete() error
}

// Recycler used for Recycle(...)
type Recycler interface {
	Reset()
}

// RecyclableCrud Crud and Recycler interface
type RecyclableCrud interface {
	Crud
	Recycler
}

// ModeFactory mode factory type
type ModeFactory func() interface{}

// RegisterModelFactory register a ModeFactory to build modePool of id
func RegisterModelFactory(id int, factory ModeFactory) {
	if factory == nil {
		logus.Panic("factory is nil")
	}
	modelPools[id] = &sync.Pool{New: factory}
}

// RegisterRepository register a Repository instance
func RegisterRepository(repo Repository) {
	if repo == nil {
		logus.Panic("repository is nil")
	}
	repository = repo
}

// Model get a Crud interface instance from pool of id
func Model(id int) interface{} {
	if pool := modelPools[id]; pool != nil {
		return pool.Get()
	}
	return errors.New("not found exists model")
}

// Recycle put a RecyclableCrud of id
func Recycle(id int, model RecyclableCrud) {
	modelPools[id].Put(model)
}

// Create mode by id
func Create(id int, model RecyclableCrud) error {
	return repository.Create(id, model)
}

// Retrieve mode by id
func Retrieve(id int, model RecyclableCrud) error {
	return repository.Retrieve(id, model)
}

// Update model by id
func Update(id int, model RecyclableCrud) error {
	return repository.Update(id, model)
}

// Delete mode by id
func Delete(id int, model RecyclableCrud) error {
	return repository.Delete(id, model)
}
