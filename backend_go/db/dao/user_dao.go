package dao

import (
	"backend_go/db/models"
	"errors"

	"gorm.io/gorm"
)

// UserDAO is a struct that holds the GORM database connection.
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO initializes and returns a new UserDAO with a GORM DB connection.
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// CRUD methods

// Create inserts a new user into the database.
func (dao *UserDAO) Create(name, email, password string) (*models.User, error) {
	user := &models.User{Name: name, Email: email, Password: password}
	if err := dao.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll retrieves all users from the database.
func (dao *UserDAO) GetAll() ([]models.User, error) {
	var users []models.User
	if err := dao.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Update updates the user's details, including the password if provided.
func (dao *UserDAO) Update(id int, name, email, password string) error {
	user := &models.User{ID: id, Name: name, Email: email, Password: password}
	if err := dao.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a user from the database.
func (dao *UserDAO) Delete(id int) error {
	if err := dao.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetByEmail checks if a user with the given email exists.
func (dao *UserDAO) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := dao.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) AuthenticateUser(identifier, password string) (*models.User, error) {
	var user models.User
	if err := dao.db.Where("email = ? AND password = ?", identifier, password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No user found
		}
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) GetByIDOrEmail(identifier string) (*models.User, error) {
	var user models.User

	// Try by ID
	if err := dao.db.First(&user, identifier).Error; err == nil {
		return &user, nil
	}

	// Fallback: Try by Email
	if err := dao.db.Where("email = ?", identifier).First(&user).Error; err == nil {
		return &user, nil
	}

	return nil, gorm.ErrRecordNotFound
}
