package controllers

import "github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"

type ProductRepositories interface {
	GetAll() ([]*models.Product, error)
	GetByID(id uint) (*models.Product, error)
	GetByName(name string) (*models.Product, error)
	Create(data models.Product) (*models.Product, error)
	Update(id uint, data models.Product) (*models.Product, error)
	Delete(id uint) error
}

type ProductServices interface {
	GetAll() ([]*models.Product, error)
	GetByID(id uint) (*models.Product, error)
	GetByName(name string) (*models.Product, error)
	Create(data models.Product) (*models.Product, error)
	Update(id uint, data models.Product) (*models.Product, error)
	Delete(id uint) error
}
