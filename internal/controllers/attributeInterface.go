package controllers

import "github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"

type AttributesRepositories interface {
	GetAll() ([]*models.Attribute, error)
	GetByID(id uint) (*models.Attribute, error)
	GetByName(name string) (*models.Attribute, error)
	Create(data models.Attribute) (*models.Attribute, error)
	Update(id uint, data models.Attribute) (*models.Attribute, error)
	Delete(id uint) error
}

type AttributesServices interface {
	GetAll() ([]*models.Attribute, error)
	GetByID(id uint) (*models.Attribute, error)
	GetByName(name string) (*models.Attribute, error)
	Create(data models.Attribute) (*models.Attribute, error)
	Update(id uint, data models.Attribute) (*models.Attribute, error)
	Delete(id uint) error
}
