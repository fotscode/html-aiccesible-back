package repositories

import (
	"errors"
	"gorm.io/gorm"
	"html-aiccesible/models"
)

type UserRepository interface {
	CreateUser(userBody *models.CreateUserBody) (*models.User, error)
	GetUser(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(userBody *models.UpdateUserBody) (*models.User, error)
	DeleteUser(id int) error
	ListUsers(page, size int) ([]models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func UserRepo(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) CreateUser(userBody *models.CreateUserBody) (*models.User, error) {
	hash, err := models.HashPassword(userBody.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Username: userBody.Username,
		Password: hash,
		Config:   models.FillConfigDefaults(&models.Configuration{}),
	}
	res := r.DB.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (r *userRepository) GetUser(id int) (*models.User, error) {
	var user models.User
	res := r.DB.First(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	res := r.DB.Where("username = ?", username).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(userBody *models.UpdateUserBody) (*models.User, error) {
	hash, err := models.HashPassword(userBody.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Username: userBody.Username,
		Password: hash,
	}
	res := r.DB.Model(&models.User{}).Where("id = ?", userBody.ID).Updates(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return r.GetUser(int(userBody.ID))
}

func (r *userRepository) DeleteUser(id int) error {
	_, err := r.GetUser(id)
	if err != nil {
		return err
	}
	res := r.DB.Delete(&models.User{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *userRepository) ListUsers(page, size int) ([]models.User, error) {
	users := []models.User{}
	res := r.DB.Limit(size).Offset((page - 1) * size).Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}
