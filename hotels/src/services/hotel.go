package services

import (
	"hotels/src/models"
	"hotels/src/repositories"

	"github.com/google/uuid"
)

type HotelService interface {
	CreateHotel(request models.HotelDTO) error
	GetHotels() ([]models.HotelDTO, error)
	GetHotelsWithType() ([]models.HotelWithTypeDTO, error)
	GetHotelById(Id uuid.UUID) ([]models.HotelDTO, error)
	GetHotelWithTypeById(Id uuid.UUID) ([]models.HotelWithTypeDTO, error)
	UpdateHotel(models.HotelDTO) error
	DeleteHotel(Id uuid.UUID) error
}

type hotelService struct {
	hotelRepo repositories.HotelRepository
}

func CreateHotelService(hotelRepo repositories.HotelRepository) HotelService {
	return &hotelService{
		hotelRepo: hotelRepo,
	}
}

func (s *hotelService) CreateHotel(hotel models.HotelDTO) error {
	_, err := s.hotelRepo.CreateHotel(hotel)
	if err != nil {
		return err
	}
	return nil
}

func (s *hotelService) GetHotels() ([]models.HotelDTO, error) {
	return s.hotelRepo.GetHotels()
}

func (s *hotelService) GetHotelsWithType() ([]models.HotelWithTypeDTO, error) {
	return s.hotelRepo.GetHotelsWithType()
}

func (s *hotelService) GetHotelById(Id uuid.UUID) ([]models.HotelDTO, error) {
	return s.hotelRepo.GetHotelById(Id)
}

func (s *hotelService) GetHotelWithTypeById(Id uuid.UUID) ([]models.HotelWithTypeDTO, error) {
	return s.hotelRepo.GetHotelWithTypeById(Id)
}

func (s *hotelService) UpdateHotel(hotel models.HotelDTO) error {
	err := s.hotelRepo.UpdateHotel(hotel)
	if err != nil {
		return err
	}
	return nil
}

func (s *hotelService) DeleteHotel(Id uuid.UUID) error {
	return s.hotelRepo.DeleteHotel(Id)
}
