package services

import (
	c "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
)

type attributesServices struct {
	repo c.AttributesRepositories
}

func NewAttributesServices(Repo c.AttributesRepositories) c.AttrbiutesServices {
	return &attributesServices{
		repo: Repo,
	}
}

func (s *attributesServices) GetAll() ([]*models.Attribute, error) {
	result, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *attributesServices) GetByID(id uint) (*models.Attribute, error) {
	result, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *attributesServices) GetByName(name string) (*models.Attribute, error) {
	result, err := s.repo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *attributesServices) Create(data models.Attribute) (*models.Attribute, error) {
	result, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *attributesServices) Update(id uint, data models.Attribute) (*models.Attribute, error) {
	result, err := s.repo.Update(id, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *attributesServices) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
