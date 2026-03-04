package dbmodel

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
}

type TagRepository interface {
	Create(tag *Tag) (*Tag, error)
	Find() ([]*Tag, error)
	FindById(id uint) (*Tag, error)
	Update(tag *Tag) (*Tag, error)
	Delete(id uint) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

// Create the tag
func (r *tagRepository) Create(tag *Tag) (*Tag, error) {
	if err := r.db.Create(tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

// Find all tag.
func (r *tagRepository) Find() ([]*Tag, error) {
	var tags []*Tag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

// Find a tag by is id.
func (r *tagRepository) FindById(id uint) (*Tag, error) {
	var tag Tag
	if err := r.db.First(&tag, id).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}

// Update the given tag.
func (r *tagRepository) Update(tag *Tag) (*Tag, error) {
	if err := r.db.Save(tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

// Delete a tag by is id.
func (r *tagRepository) Delete(id uint) error {
	if err := r.db.Delete(Tag{}, id).Error; err != nil {
		return err
	}

	return nil
}
