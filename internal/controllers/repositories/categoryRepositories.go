package repositories

import (
	c "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type categoryRepositories struct {
	db *gorm.DB
}

func NewCategoryRepositories(DB *gorm.DB) c.CategoryRepositories {
	return &categoryRepositories{
		db: DB,
	}
}

func (r *categoryRepositories) GetAll() ([]*models.Category, error) {
	var categories []*models.Category

	if err := r.db.Preload("Products").Find(&categories); err != nil {
		return nil, err.Error
	}

	return categories, nil
}

func (r *categoryRepositories) GetByID(id uint) (*models.Category, error) {
	var category *models.Category

	if err := r.db.Preload("Product").First(&category, id); err != nil {
		return nil, err.Error
	}

	return category, nil
}

func (r *categoryRepositories) GetByName(name string) (*models.Category, error) {
	var category *models.Category

	if err := r.db.Preload("Product").First(&category, name); err != nil {
		return nil, err.Error
	}

	return category, nil
}

func (r *categoryRepositories) Create(data models.Category) (*models.Category, error) {
	if err := r.db.Save(&data); err != nil {
		return nil, err.Error
	}

	return &data, nil
}

func (r *categoryRepositories) Update(id uint, data models.Category) (*models.Category, error) {
	if err := r.db.Model(&models.Category{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&data); err != nil {
		return nil, err.Error
	}

	return &data, nil
}

func (r *categoryRepositories) Delete(id uint) error {
	if err := r.db.Delete(&models.Category{}, id); err != nil {
		return err.Error
	}

	return nil
}
