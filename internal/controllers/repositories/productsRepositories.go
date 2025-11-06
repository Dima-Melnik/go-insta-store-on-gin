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

	if err := r.db.Preload("Attributes").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepositorues) GetByID(id uint) (*models.Product, error) {
	var product *models.Product

	if err := r.db.Preload("Attributes").First(&product, id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepositorues) GetByName(name string) (*models.Product, error) {
	var product *models.Product

	if err := r.db.Preload("Attributes").Where("name ILIKE ?", name).First(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepositorues) Create(data models.Product) (*models.Product, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *productRepositorues) Update(id uint, data models.Product) (*models.Product, error) {
	if err := r.db.Model(&models.Product{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *productRepositorues) Delete(id uint) error {
	if err := r.db.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}
