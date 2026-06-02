package dbmodel

import "gorm.io/gorm"

type MediaField struct {
	FieldID uint   `gorm:"primaryKey;column:fieldId"`
	MediaID uint   `gorm:"primaryKey;column:mediaId"`
	Value   string `gorm:"column:value"`

	Field Field `gorm:"foreignKey:FieldID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Media Media `gorm:"foreignKey:MediaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type MediaFieldRepository interface {
	Create(mediaField *MediaField) (*MediaField, error)
	FindAll() ([]*MediaField, error)
	FindById(fieldID uint, mediaID uint) (*MediaField, error)
	Update(mediaField *MediaField) (*MediaField, error)
	Delete(id uint) error
}

type mediaFieldRepository struct {
	db *gorm.DB
}

func NewMediaFieldRepository(db *gorm.DB) MediaFieldRepository {
	return &mediaFieldRepository{db: db}
}

// Create the mediaField
func (r *mediaFieldRepository) Create(mediaField *MediaField) (*MediaField, error) {
	if err := r.db.Create(mediaField).Error; err != nil {
		return nil, err
	}

	return mediaField, nil
}

// Find all mediaField.
func (r *mediaFieldRepository) FindAll() ([]*MediaField, error) {
	var fields []*MediaField
	if err := r.db.Find(&fields).Error; err != nil {
		return nil, err
	}

	return fields, nil
}

// Find a mediaField by is fieldID and mediaID.
func (r *mediaFieldRepository) FindById(fieldID uint, mediaID uint) (*MediaField, error) {
	var mediaField MediaField
	if err := r.db.First(&mediaField, fieldID, mediaID).Error; err != nil {
		return nil, err
	}

	return &mediaField, nil
}

// Update the given mediaField.
func (r *mediaFieldRepository) Update(mediaField *MediaField) (*MediaField, error) {
	if err := r.db.Save(mediaField).Error; err != nil {
		return nil, err
	}

	return mediaField, nil
}

// Delete a mediaField by is id.
func (r *mediaFieldRepository) Delete(id uint) error {
	if err := r.db.Delete(MediaField{}, id).Error; err != nil {
		return err
	}

	return nil
}
