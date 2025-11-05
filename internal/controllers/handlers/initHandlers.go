package handlers

import (
	c "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers"
	r "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers/repositories"
	s "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers/services"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/db"
)

type Handler struct {
	productService  c.ProductServices
	categoryService c.CategoryServices
}

type HandlerConfig struct {
	ProductService  c.ProductServices
	CategoryService c.CategoryServices
}

func NewHandler(cfg *HandlerConfig) *Handler {
	return &Handler{
		productService:  cfg.ProductService,
		categoryService: cfg.CategoryService,
	}
}

func InitAllHandlers() *Handler {
	db := db.GetDB()

	productRepo := r.NewProductRepositories(db)
	categoryRepo := r.NewCategoryRepositories(db)

	productService := s.NewProductServices(productRepo)
	categoryService := s.NewCategoryServices(categoryRepo)

	h := NewHandler(&HandlerConfig{
		ProductService:  productService,
		CategoryService: categoryService,
	})

	return h
}
