package dbmodel

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
}

type CollectionRepository interface {
	Create(collection *Collection) (*Collection, error)
	Find() ([]*Collection, error)
	FindById(id uint) (*Collection, error)
	Update(collection *Collection) (*Collection, error)
	Delete(id uint) error
}

type collectionRepository struct {
	db *gorm.DB
}

func NewCollectionRepository(db *gorm.DB) CollectionRepository {
	return &collectionRepository{db: db}
}

// Create the collection
func (r *collectionRepository) Create(collection *Collection) (*Collection, error) {
	if err := r.db.Create(collection).Error; err != nil {
		return nil, err
	}

	return collection, nil
}

// Find all collection.
func (r *collectionRepository) Find() ([]*Collection, error) {
	var collections []*Collection
	if err := r.db.Find(&collections).Error; err != nil {
		return nil, err
	}

	return collections, nil
}

// Find a collection by is id.
func (r *collectionRepository) FindById(id uint) (*Collection, error) {
	var collection Collection
	if err := r.db.First(&collection, id).Error; err != nil {
		return nil, err
	}

	return &collection, nil
}

// Update the given collection.
func (r *collectionRepository) Update(collection *Collection) (*Collection, error) {
	if err := r.db.Save(collection).Error; err != nil {
		return nil, err
	}

	return collection, nil
}

// Delete a collection by is id.
func (r *collectionRepository) Delete(id uint) error {
	if err := r.db.Delete(Collection{}, id).Error; err != nil {
		return err
	}

	return nil
}
