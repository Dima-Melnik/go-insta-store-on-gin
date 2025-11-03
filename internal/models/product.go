package models

type Product struct {
	Model
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float64 `json:"price"`
	ImageURL   string  `json:"image_url"`
	CategoryID uint    `json:"category_id"`
	Category   Category
	Attributes []Attribute
}
