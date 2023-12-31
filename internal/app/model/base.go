package model

type BaseModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Location    string `json:"location"`
}
