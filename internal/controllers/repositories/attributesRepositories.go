package repositories

import (
	c "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type attributesRepositories struct {
	db *gorm.DB
}

func NewAttrbiteRepositories(DB *gorm.DB) c.AttributesRepositories {
	return &attributesRepositories{
		db: DB,
	}
}

func (r *attributesRepositories) GetAll() ([]*models.Attribute, error) {
	var attributes []*models.Attribute

	if err := r.db.Preload("Products").Find(&attributes); err != nil {
		return nil, err.Error
	}

	return attributes, nil
}

func (r *attributesRepositories) GetByID(id uint) (*models.Attribute, error) {
	var attribute *models.Attribute

	if err := r.db.Preload("Products").First(&attribute, id); err != nil {
		return nil, err.Error
	}

	return attribute, nil
}

func (r *attributesRepositories) GetByName(name string) (*models.Attribute, error) {
	var attribute *models.Attribute

	if err := r.db.Preload("Products").First(&attribute, name); err != nil {
		return nil, err.Error
	}

	return attribute, nil
}

func (r *attributesRepositories) Create(data models.Attribute) (*models.Attribute, error) {
	if err := r.db.Save(&data); err != nil {
		return nil, err.Error
	}

	return &data, nil
}

func (r *attributesRepositories) Update(id uint, data models.Attribute) (*models.Attribute, error) {
	if err := r.db.Model(&models.Attribute{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&data); err != nil {
		return nil, err.Error
	}

	return &data, nil
}

func (r *attributesRepositories) Delete(id uint) error {
	if err := r.db.Delete(&models.Attribute{}, id); err != nil {
		return err.Error
	}

	return nil
}
