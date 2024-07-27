package repositories

import (
	"encoding/json"
	"html-aiccesible/models"

	"gorm.io/gorm"
)

type ConfigRepository interface {
	UpdateConfig(userID int, configBody *models.UpdateConfigBody) (*models.Configuration, error)
	GetConfig(userID int) (*models.Configuration, error)
}

type configRepository struct {
	DB *gorm.DB
}

func ConfigRepo() ConfigRepository {
	return &configRepository{
		DB: models.GetDB(),
	}
}

func (r *configRepository) UpdateConfig(userID int, configBody *models.UpdateConfigBody) (*models.Configuration, error) {
	var config map[string]interface{}
	configBytes, err := json.Marshal(configBody)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}
	res := r.DB.Model(models.Configuration{}).Where("user_id = ?", userID).Updates(config)

	if res.Error != nil {
		return nil, res.Error
	}
	return r.GetConfig(userID)
}

func (r *configRepository) GetConfig(userID int) (*models.Configuration, error) {
	var config models.Configuration
	res := r.DB.Where("user_id = ?", userID).First(&config)
	if res.Error != nil {
		return nil, res.Error
	}
	return &config, nil
}
