package controllers

import "github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"

type CategoryRepositories interface {
	GetAll() ([]*models.Category, error)
	GetByID(id uint) (*models.Category, error)
	GetByName(name string) (*models.Category, error)
	Create(data models.Category) (*models.Category, error)
	Update(id uint, data models.Category) (*models.Category, error)
	Delete(id uint) error
}

type CategoryServices interface {
	GetAll() ([]*models.Category, error)
	GetByID(id uint) (*models.Category, error)
	GetByName(name string) (*models.Category, error)
	Create(data models.Category) (*models.Category, error)
	Update(id uint, data models.Category) (*models.Category, error)
	Delete(id uint) error
}
