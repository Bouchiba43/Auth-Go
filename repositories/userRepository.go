package repositories

import (
	"github.com/Bouchiba43/Auth-Go/models"
	"gorm.io/gorm"
)

type userRepository interface {
	FindAll() ([]models.User, error)
	FindById(id int) (models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(data models.User) (models.User, error)
	UpdateById(id int, data models.User) (models.User, error)
	DeleteById(id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindById(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(data models.User) (models.User, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *UserRepository) UpdateById(id int, data models.User) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	err = r.db.Model(&user).Updates(data).Error
	return user, err
}

func (r *UserRepository) DeleteById(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}
