package dto

import "github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"

type CreateProduct struct {
	Name       string  `json:"name" binding:"required"`
	Desc       string  `json:"desc" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	ImageURL   string  `json:"image_url"`
	CategoryID uint    `json:"category_id"`
}

type UpdateProduct struct {
	Name       *string  `json:"name"`
	Desc       *string  `json:"desc"`
	Price      *float64 `json:"price"`
	ImageURL   *string  `json:"image_url"`
	CategoryID *uint    `json:"category_id"`
}

type ProductResponse struct {
	models.Model
	Name       string          `json:"name"`
	Desc       string          `json:"desc"`
	Price      float64         `json:"price"`
	ImageURL   string          `json:"image_url"`
	Category   CategoryLite    `json:"category"`
	Attributes []AttributeLite `json:"attributes"`
}

func ToProductResponse(p *models.Product) ProductResponse {
	attrs := make([]AttributeLite, len(p.Attributes))
	for i, a := range p.Attributes {
		attrs[i] = AttributeLite{
			ID:    a.ID,
			Name:  a.Name,
			Value: a.Value,
		}
	}

	return ProductResponse{
		Model:    p.Model,
		Name:     p.Name,
		Desc:     p.Desc,
		Price:    p.Price,
		ImageURL: p.ImageURL,
		Category: CategoryLite{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		},
		Attributes: attrs,
	}
}
