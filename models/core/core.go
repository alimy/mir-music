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
	Create(int, RecyclableCreate) error
	Retrieve(int, RecyclableRetrieve) error
	Update(int, RecyclableUpdate) error
	Delete(int, RecyclableDelete) error
}

// OnRepository data operator
type OnRepository interface {
	OnCreate(RecyclableCreate) error
	OnRetrieve(RecyclableRetrieve) error
	OnUpdate(RecyclableUpdate) error
	OnDelete(RecyclableDelete) error
}

// Template mode data template operator
type Template interface {
	Create(RecyclableCreate) error
	Retrieve(RecyclableRetrieve) error
	Update(RecyclableUpdate) error
	Delete(RecyclableDelete) error
}

// CrudCreate create mode data operator
type CrudCreate interface {
	Create() error
}

// CrudRetrieve retrieve mode data operator
type CrudRetrieve interface {
	Retrieve() error
}

// CrudUpdate update mode data operator
type CrudUpdate interface {
	Update() error
}

// CrudDelete delete mode data operator
type CrudDelete interface {
	Delete() error
}

// Crud mode data operator
type Crud interface {
	CrudCreate
	CrudRetrieve
	CrudUpdate
	CrudDelete
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

// RecyclableCreate interface
type RecyclableCreate interface {
	CrudCreate
	Recycler
}

// RecyclableRetrieve interface
type RecyclableRetrieve interface {
	CrudRetrieve
	Recycler
}

// RecyclableUpdate interface
type RecyclableUpdate interface {
	CrudUpdate
	Recycler
}

// RecyclableDelete interface
type RecyclableDelete interface {
	CrudDelete
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
func Recycle(id int, model Recycler) {
	modelPools[id].Put(model)
}

// Create mode by id
func Create(id int, model RecyclableCreate) error {
	return repository.Create(id, model)
}

// Retrieve mode by id
func Retrieve(id int, model RecyclableRetrieve) error {
	return repository.Retrieve(id, model)
}

// Update model by id
func Update(id int, model RecyclableUpdate) error {
	return repository.Update(id, model)
}

// Delete mode by id
func Delete(id int, model RecyclableDelete) error {
	return repository.Delete(id, model)
}
