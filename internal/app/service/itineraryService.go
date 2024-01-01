package service

import (
	"github.com/simonha9/T2DGen/internal/app/model"
)

// ItineraryService is an interface for the ItineraryService
type ItineraryService interface {
	GenerateItinerary() (*model.Itinerary, error)
}
