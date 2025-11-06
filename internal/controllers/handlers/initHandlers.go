package handlers

import (
	c "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers"
	r "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers/repositories"
	s "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers/services"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/db"
)

type Handler struct {
	productService   c.ProductServices
	categoryService  c.CategoryServices
	attributeService c.AttributesServices
}

type HandlerConfig struct {
	ProductService  c.ProductServices
	CategoryService c.CategoryServices
	AttrbiteService c.AttributesServices
}

func NewHandler(cfg *HandlerConfig) *Handler {
	return &Handler{
		productService:   cfg.ProductService,
		categoryService:  cfg.CategoryService,
		attributeService: cfg.AttrbiteService,
	}
}

func InitAllHandlers() *Handler {
	db := db.GetDB()

	productRepo := r.NewProductRepositories(db)
	categoryRepo := r.NewCategoryRepositories(db)
	attributeRepo := r.NewAttrbiteRepositories(db)

	productService := s.NewProductServices(productRepo)
	categoryService := s.NewCategoryServices(categoryRepo)
	attributeService := s.NewAttributesServices(attributeRepo)

	h := NewHandler(&HandlerConfig{
		ProductService:  productService,
		CategoryService: categoryService,
		AttrbiteService: attributeService,
	})

	return h
}
