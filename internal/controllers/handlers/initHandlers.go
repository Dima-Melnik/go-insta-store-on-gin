package handlers

import (
	c "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers"
	r "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers/repositories"
	s "github.com/Dima-Melnik/go-insta-store-on-gin/internal/controllers/services"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/db"
)

type Handler struct {
	productService c.ProductServices
}

type HandlerConfig struct {
	ProductService c.ProductServices
}

func NewHandler(cfg *HandlerConfig) *Handler {
	return &Handler{
		productService: cfg.ProductService,
	}
}

func InitAllHandlers() *Handler {
	db := db.GetDB()

	productRepo := r.NewProductRepositories(db)

	productService := s.NewProductServices(productRepo)

	h := NewHandler(&HandlerConfig{
		ProductService: productService,
	})

	return h
}
