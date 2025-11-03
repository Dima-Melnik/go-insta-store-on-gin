package dto

import "github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"

type CreateCategory struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategory struct {
	Name *string `json:"name"`
}

type CategoryResponse struct {
	models.Model
	Name     string        `json:"name"`
	Products []ProductLite `json:"products"`
}

func ToCategoryResponse(c *models.Category) CategoryResponse {
	products := make([]ProductLite, len(c.Products))
	for i, p := range c.Products {
		products[i] = ProductLite{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			ImageURL: p.ImageURL,
		}
	}

	return CategoryResponse{
		Model:    c.Model,
		Name:     c.Name,
		Products: products,
	}
}
