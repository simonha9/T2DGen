package service

import (
	"github.com/simonha9/T2DGen/internal/app/model"
)

// ItineraryService is an interface for the ItineraryService
type ItineraryService interface {
	GenerateItinerary() ([]*model.Itinerary, error)
}

type ItineraryServiceImpl struct{}

// NewItineraryService creates a new ItineraryService
func NewItineraryService() ItineraryService {
	return &ItineraryServiceImpl{}
}

/*
 * GenerateItinerary generates an itinerary by making yelp API call to get top 5 in area
 * @return Itinerary
 */
func (i *ItineraryServiceImpl) GenerateItinerary() ([]*model.Itinerary, error) {
	res := make([]*model.Itinerary, 0)
	return res, nil
}
