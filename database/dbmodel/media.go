package dbmodel

import (
	t "time"

	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	UserId         uint    `json:"user_id"`
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	MediaType      string  `json:"media_type"`
	ImgURL         *string `json:"img_url"`
	Rating         *int    `json:"rating"`
	Notes          *string `json:"notes"`
	Description    *string `json:"description"`
	Genre          *string `json:"genre"`
	StartDate      *t.Time `json:"start_date"`
	CompletionDate *t.Time `json:"completion_date"`

	User   User     `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags   []*Tag   `gorm:"many2many:media_tag;constraint:OnDelete:CASCADE" json:"tags"`
	Fields []*Field `gorm:"many2many:media_fields;constraint:OnDelete:CASCADE" json:"fields"`
}

type MediaRepository interface {
	Create(media *Media) (*Media, error)
	FindAll(userID uint) ([]*Media, error)
	FindById(id uint) (*Media, error)
	FindTagsByMedia(id uint) ([]Tag, error)
	FindFieldsByMedia(id uint) ([]Field, error)
	AddTag(mediaID, tagID uint) error
	RemoveTag(mediaID, tagID uint) error
	Update(media *Media) (*Media, error)
	Delete(id uint) error
}

type mediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{db: db}
}

// Create the media
func (r *mediaRepository) Create(media *Media) (*Media, error) {
	if err := r.db.Create(media).Error; err != nil {
		return nil, err
	}

	return media, nil
}

// Find all media for a specific user.
func (r *mediaRepository) FindAll(userID uint) ([]*Media, error) {
	var medias []*Media
	if err := r.db.Where("user_id = ?", userID).Find(&medias).Error; err != nil {
		return nil, err
	}
	return medias, nil
}

// Find a media by his id.
func (r *mediaRepository) FindById(id uint) (*Media, error) {
	var media Media
	if err := r.db.First(&media, id).Error; err != nil {
		return nil, err
	}

	return &media, nil
}

// Find tags for a media.
func (r *mediaRepository) FindTagsByMedia(id uint) ([]Tag, error) {
	var media Media
	if err := r.db.Preload("Tags").First(&media, id).Error; err != nil {
		return nil, err
	}
	tags := make([]Tag, len(media.Tags))
	for i, u := range media.Tags {
		tags[i] = *u
	}

	return tags, nil
}

// Find fields for a media.
func (r *mediaRepository) FindFieldsByMedia(id uint) ([]Field, error) {
	var media Media
	if err := r.db.Preload("Fields").First(&media, id).Error; err != nil {
		return nil, err
	}
	fields := make([]Field, len(media.Fields))
	for i, u := range media.Fields {
		fields[i] = *u
	}

	return fields, nil
}

// Add a tag to a media.
func (r *mediaRepository) AddTag(mediaID, tagID uint) error {
	return r.db.Model(&Media{Model: gorm.Model{ID: mediaID}}).Association("Tags").Append(&Tag{Model: gorm.Model{ID: tagID}})
}

// Remove a tag from a media.
func (r *mediaRepository) RemoveTag(mediaID, tagID uint) error {
	return r.db.Model(&Media{Model: gorm.Model{ID: mediaID}}).Association("Tags").Delete(&Tag{Model: gorm.Model{ID: tagID}})
}

// Update the given media.
func (r *mediaRepository) Update(media *Media) (*Media, error) {
	if err := r.db.Save(media).Error; err != nil {
		return nil, err
	}

	return media, nil
}

// Delete a media by is id.
func (r *mediaRepository) Delete(id uint) error {
	if err := r.db.Delete(&Media{}, id).Error; err != nil {
		return err
	}

	return nil
}
