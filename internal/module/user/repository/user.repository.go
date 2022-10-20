package repository

import (
	"github.com/HiBang15/sample-gorm.git/internal/database"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.Connection}
}
func (userRepo *UserRepository) Create(user *entities.User) error {
	result := userRepo.db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (userRepo *UserRepository) GetUsers() ([]entities.User, error) {
	var users []entities.User
	result := userRepo.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
	//panic("pls install func")
}
func (userRepo *UserRepository) GetUser(id string) (*entities.User, error) {
	var user *entities.User
	result := userRepo.db.Where("id=?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (userRepo *UserRepository) DeleteUser(id string) error {
	var user *entities.User
	result := userRepo.db.Where("id=?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (userRepo *UserRepository) UpdateUser(id string, user *entities.User) error {
	/*
		Updates supports updating with struct or map[string]interface{},
		when updating with struct it will only update non-zero fields by default
	*/
	var ogUser *entities.User

	result := userRepo.db.Where("id=?", id).First(&ogUser)
	if result.Error != nil {
		return result.Error
	}

	if err := userRepo.db.Model(&ogUser).Updates(user); err != nil {
		return err.Error
	}
	return nil
}
