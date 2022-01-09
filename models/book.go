package models

type Book struct {
	Name            string  `json:"name"`
	Author          string  `json:"author"`
	Rating          float64 `json:"rating"`
	Report          string  `json:"report"`
	ImageUrl        string  `json:"image_url"`
	BackgroundColor string  `json:"background_color"`
}
