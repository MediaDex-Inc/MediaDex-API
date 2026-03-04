package dbmodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
}

type UserRepository interface {
	Create(user *User) (*User, error)
	Find() ([]*User, error)
	FindById(id uint) (*User, error)
	Update(user *User) (*User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create the user
func (r *userRepository) Create(user *User) (*User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Find all user.
func (r *userRepository) Find() ([]*User, error) {
	var users []*User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Find a user by is id.
func (r *userRepository) FindById(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update the given user.
func (r *userRepository) Update(user *User) (*User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Delete a user by is id.
func (r *userRepository) Delete(id uint) error {
	if err := r.db.Delete(User{}, id).Error; err != nil {
		return err
	}

	return nil
}
