package model

type Itinerary struct {
	Activities []Activity `json:"activities"`
	Food       []Food     `json:"food"`
}
