package dto

type CategoryLite struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AttributeLite struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProductLite struct {
	ID       uint         `json:"id"`
	Name     string       `json:"name"`
	Desc     string       `json:"desc"`
	Price    float64      `json:"price"`
	ImageURL string       `json:"image_url"`
	Category CategoryLite `json:"category"`
}
