package services

import (
	c "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
)

type productService struct {
	repo c.ProductRepositories
}

func NewProductServices(Repo c.ProductRepositories) c.ProductServices {
	return &productService{
		repo: Repo,
	}
}

func (s *productService) GetAll() ([]*models.Product, error) {
	result, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *productService) GetByID(id uint) (*models.Product, error) {
	result, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *productService) GetByName(name string) (*models.Product, error) {
	result, err := s.repo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *productService) Create(data models.Product) (*models.Product, error) {
	result, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *productService) Update(id uint, data models.Product) (*models.Product, error) {
	result, err := s.repo.Update(id, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *productService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
