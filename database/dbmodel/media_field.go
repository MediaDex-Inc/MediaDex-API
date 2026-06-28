package dbmodel

import "gorm.io/gorm"

type MediaField struct {
	FieldID uint   `gorm:"not null;primaryKey;constraint:OnDelete:CASCADE"`
	MediaID uint   `gorm:"not null;primaryKey;constraint:OnDelete:CASCADE"`
	Value   string `gorm:"not null;column:value"`
}

type MediaFieldRepository interface {
	Create(mediaField *MediaField) (*MediaField, error)
	FindAll(userID uint) ([]*MediaField, error)
	FindById(fieldID uint, mediaID uint) (*MediaField, error)
	FindByMediaId(mediaID uint) ([]*MediaField, error)
	Update(mediaField *MediaField) (*MediaField, error)
	Delete(fieldID uint, mediaID uint) error
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

// Find all mediaFields for a specific user (via media ownership).
func (r *mediaFieldRepository) FindAll(userID uint) ([]*MediaField, error) {
	var fields []*MediaField
	if err := r.db.Joins("JOIN media ON media.id = media_fields.media_id AND media.deleted_at IS NULL").
		Where("media.user_id = ?", userID).
		Find(&fields).Error; err != nil {
		return nil, err
	}

	return fields, nil
}

// Find all mediaFields for a specific media.
func (r *mediaFieldRepository) FindByMediaId(mediaID uint) ([]*MediaField, error) {
	var mediaFields []*MediaField
	if err := r.db.Where("media_id = ?", mediaID).Find(&mediaFields).Error; err != nil {
		return nil, err
	}

	return mediaFields, nil
}

// Find a mediaField by his fieldID and mediaID.
func (r *mediaFieldRepository) FindById(fieldID uint, mediaID uint) (*MediaField, error) {
	var mediaField MediaField
	if err := r.db.Where("field_id = ? AND media_id = ?", fieldID, mediaID).First(&mediaField).Error; err != nil {
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

// Delete a mediaField by his composite id.
func (r *mediaFieldRepository) Delete(fieldID uint, mediaID uint) error {
	if err := r.db.Where("field_id = ? AND media_id = ?", fieldID, mediaID).Delete(&MediaField{}).Error; err != nil {
		return err
	}

	return nil
}
