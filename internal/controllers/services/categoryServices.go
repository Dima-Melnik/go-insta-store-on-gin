package services

import (
	c "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
)

type categoryServices struct {
	repo c.CategoryRepositories
}

func NewCategoryServices(Repo c.CategoryRepositories) c.CategoryServices {
	return &categoryServices{
		repo: Repo,
	}
}

func (s *categoryServices) GetAll() ([]*models.Category, error) {
	result, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *categoryServices) GetByID(id uint) (*models.Category, error) {
	result, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *categoryServices) GetByName(name string) (*models.Category, error) {
	result, err := s.repo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *categoryServices) Create(data models.Category) (*models.Category, error) {
	result, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *categoryServices) Update(id uint, data models.Category) (*models.Category, error) {
	result, err := s.repo.Update(id, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *categoryServices) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
