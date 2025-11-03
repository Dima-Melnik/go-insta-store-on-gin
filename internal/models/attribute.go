package models

type Attribute struct {
	Model
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Products []Product `gorm:"many2many:product_attributes;"`
}
