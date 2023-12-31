package model

type Food struct {
	base    BaseModel
	Cuisine string `json:"cuisine"`
}
