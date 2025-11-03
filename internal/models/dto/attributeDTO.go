package dto

import "github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"

type CreateAttribute struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type UpdateAttribute struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

type AttributeResponse struct {
	models.Model
	Name     string        `json:"name"`
	Value    string        `json:"value"`
	Products []ProductLite `json:"products"`
}

func ToAttributeResponse(a *models.Attribute) AttributeResponse {
	products := make([]ProductLite, len(a.Products))
	for i, p := range a.Products {
		products[i] = ProductLite{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			ImageURL: p.ImageURL,
			Category: CategoryLite{
				ID:   p.Category.ID,
				Name: p.Category.Name,
			},
		}
	}

	return AttributeResponse{
		Model:    a.Model,
		Name:     a.Name,
		Value:    a.Value,
		Products: products,
	}
}
