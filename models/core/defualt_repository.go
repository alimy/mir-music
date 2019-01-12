package core

// DefaultRepository default Repository interface implement
type DefaultRepository struct {
	OnRepository
}

// Create a mode
func (r *DefaultRepository) Create(id int, model RecyclableCrud) error {
	err := r.OnCreate(model)
	model.Reset()
	Recycle(id, model)
	return err
}

// Retrieve a mode
func (r *DefaultRepository) Retrieve(id int, model RecyclableCrud) error {
	err := r.OnRetrieve(model)
	return err
}

// Update a Model
func (r *DefaultRepository) Update(id int, model RecyclableCrud) error {
	err := r.OnUpdate(model)
	model.Reset()
	Recycle(id, model)
	return err
}

// Delete remove model
func (r *DefaultRepository) Delete(id int, model RecyclableCrud) error {
	err := r.OnDelete(model)
	model.Reset()
	Recycle(id, model)
	return err
}
