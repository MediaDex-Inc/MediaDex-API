package dbmodel

import (
	t "time"

	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	userId         int     `json:"user_id"`
	name           string  `json:"name"`
	status         string  `json:"status"`
	mediaType      string  `json:"media_type"`
	imgURL         *string `json:"img_url"`
	rating         *int    `json:"rating"`
	notes          *string `json:"notes"`
	description    *string `json:"description"`
	genre          *string `json:"genre"`
	startDate      *t.Time `json:"start_date"`
	completionDate *t.Time `json:"completion_date"`
}

type MediaRepository interface {
	Create(media *Media) (*Media, error)
	Find() ([]*Media, error)
	FindById(id uint) (*Media, error)
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

// Find all media.
func (r *mediaRepository) Find() ([]*Media, error) {
	var medias []*Media
	if err := r.db.Find(&medias).Error; err != nil {
		return nil, err
	}

	return medias, nil
}

// Find a media by is id.
func (r *mediaRepository) FindById(id uint) (*Media, error) {
	var media Media
	if err := r.db.First(&media, id).Error; err != nil {
		return nil, err
	}

	return &media, nil
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
	if err := r.db.Delete(Media{}, id).Error; err != nil {
		return err
	}

	return nil
}
