package repositories

import (
	"html-aiccesible/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(userBody *models.CreateUserBody) (*models.User, error)
	GetUser(id int) (*models.User, error)
	UpdateUser(userBody *models.UpdateUserBody) (*models.User, error)
	DeleteUser(id int) error
	ListUsers(page, size int) ([]models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func UserRepo() Repository {
	return &userRepository{
		DB: models.GetDB(),
	}
}

func (r *userRepository) CreateUser(userBody *models.CreateUserBody) (*models.User, error) {
	user := &models.User{
		Username: userBody.Username,
		Password: userBody.Password,
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

func (r *userRepository) UpdateUser(userBody *models.UpdateUserBody) (*models.User, error) {
	user := &models.User{
		Model: gorm.Model{
			ID: userBody.ID,
		},
		Username: userBody.Username,
		Password: userBody.Password,
	}
	res := r.DB.Save(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id int) error {
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
