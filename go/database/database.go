package database

import (
	"github.com/Mikatech/CTF-AI/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewSqliteDB(databaseName string) (*UserService, error) {
	db, err := gorm.Open(sqlite.Open(databaseName))
	if err != nil {
		return nil, err
	}

	svc := &UserService{db: db}
	if err = svc.db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return svc, nil
}

func (svc UserService) GetUsers() ([]model.User, error) {
	var users []model.User

	svc.db.Find(&users)

	return users, nil
}

func (svc UserService) GetUser(userId string) (model.User, error) {
	var user model.User

	err := svc.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (svc UserService) CreateUser(newUser model.User) error {
	if err := svc.db.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func (svc UserService) DeleteUser(userToDelete model.User) error {
	if err := svc.db.Delete(&userToDelete).Error; err != nil {
		return err
	}

	return nil
}
