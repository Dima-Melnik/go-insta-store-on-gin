package repositories

import (
	c "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepositorues struct {
	db *gorm.DB
}

func NewProductRepositories(DB *gorm.DB) c.ProductRepositories {
	return &productRepositorues{
		db: DB,
	}
}

func (r *productRepositorues) GetAll() ([]*models.Product, error) {
	var products []*models.Product

	if err := r.db.Preload("Attributes").Find(&products); err != nil {
		return nil, err.Error
	}

	return products, nil
}

func (r *productRepositorues) GetByID(id uint) (*models.Product, error) {
	var product *models.Product

	if err := r.db.Preload("Attributes").First(&product, id); err != nil {
		return nil, err.Error
	}

	return product, nil
}

func (r *productRepositorues) GetByName(name string) (*models.Product, error) {
	var product *models.Product

	if err := r.db.Preload("Attributes").First(&product, name); err != nil {
		return nil, err.Error
	}

	return product, nil
}

func (r *productRepositorues) Create(data models.Product) (*models.Product, error) {
	if err := r.db.Create(&data); err != nil {
		return nil, err.Error
	}

	return &data, nil
}

func (r *productRepositorues) Update(id uint, data models.Product) (*models.Product, error) {
	if err := r.db.Model(&models.Product{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&data); err != nil {
		return nil, err.Error
	}

	return &data, nil
}

func (r *productRepositorues) Delete(id uint) error {
	if err := r.db.Delete(&models.Product{}, id); err != nil {
		return err.Error
	}

	return nil
}
