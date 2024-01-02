package service

import (
	"testing"

	"github.com/simonha9/T2DGen/internal/app/model"
)

func TestNewItineraryService(t *testing.T) {
	tests := []struct {
		name string
		want ItineraryService
	}{
		{"TestNewItineraryService", &ItineraryServiceImpl{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItineraryService(); got == nil {
				t.Errorf("NewItineraryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItineraryServiceImpl_GenerateItinerary(t *testing.T) {
	type fields struct {
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.Itinerary
		wantErr bool
	}{
		{"TestGenerateItinerary", fields{}, []*model.Itinerary{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ItineraryServiceImpl{}
			got, err := i.GenerateItinerary()
			if (err != nil) != tt.wantErr {
				t.Errorf("ItineraryServiceImpl.GenerateItinerary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("ItineraryServiceImpl.GenerateItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}
