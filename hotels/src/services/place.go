package services

import (
	"hotels/src/models"
	"hotels/src/repositories"

	"github.com/google/uuid"
)

type PlaceService interface {
	CreatePlace(request models.PlaceDTO) error
	GetPlaces() ([]models.PlaceDTO, error)
	GetPlaceById(id uuid.UUID) ([]models.PlaceDTO, error)
}

type placeService struct {
	placeRepo repositories.PlaceRepository
}

func CreatePlaceService(placeRepo repositories.PlaceRepository) PlaceService {
	return &placeService{
		placeRepo: placeRepo,
	}
}

func (s *placeService) GetPlaces() ([]models.PlaceDTO, error) {
	return s.placeRepo.GetPlaces()
}

func (s *placeService) GetPlaceById(id uuid.UUID) ([]models.PlaceDTO, error) {
	return s.placeRepo.GetPlaceById(id)
}

func (s *placeService) CreatePlace(place models.PlaceDTO) error {
	_, err := s.placeRepo.CreatePlace(place)
	if err != nil {
		return err
	}
	return nil
}
