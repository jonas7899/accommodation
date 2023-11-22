package services

import (
	"hotels/src/models"
	"hotels/src/repositories"
	"time"

	"github.com/google/uuid"
)

type ReservationService interface {
	CreateReservation(request models.ReservationDTO) error
	GetReservations() ([]models.ReservationDTO, error)
	GetReservationById(id uuid.UUID) ([]models.ReservationDTO, error)
	DeleteReservation(id uuid.UUID) error
	GetAllReservationsByRoomIdWithIntrval(roomId uuid.UUID, begin, end time.Time) ([]uuid.UUID, error)
	GetAllReservationsByHotelIdWithIntrval(hotelId uuid.UUID, begin, end time.Time) ([]models.ReservedRoomID, error)
	CheckFreeByRoomIdWithIntrval(roomId uuid.UUID, begin, end time.Time) (bool, error)
}

type reservationService struct {
	reservationRepo repositories.ReservationRepository
}

func CreateReservationService(reservationRepo repositories.ReservationRepository) ReservationService {
	return &reservationService{
		reservationRepo: reservationRepo,
	}
}

func (s *reservationService) GetReservations() ([]models.ReservationDTO, error) {
	return s.reservationRepo.GetReservations()
}

func (s *reservationService) GetReservationById(id uuid.UUID) ([]models.ReservationDTO, error) {
	return s.reservationRepo.GetReservationById(id)
}

func (s *reservationService) CreateReservation(reservation models.ReservationDTO) error {
	_, err := s.reservationRepo.CreateReservation(reservation)
	if err != nil {
		return err
	}
	return nil
}

func (s *reservationService) DeleteReservation(id uuid.UUID) error {
	return s.reservationRepo.DeleteReservation(id)
}

func (s *reservationService) GetAllReservationsByRoomIdWithIntrval(roomId uuid.UUID, begin, end time.Time) ([]uuid.UUID, error) {
	return s.reservationRepo.GetAllReservationsByRoomIdWithIntrval(roomId, begin, end)
}

func (s *reservationService) GetAllReservationsByHotelIdWithIntrval(hotelId uuid.UUID, begin, end time.Time) ([]models.ReservedRoomID, error) {
	return s.reservationRepo.GetAllReservationsByHotelIdWithIntrval(hotelId, begin, end)
}

func (s *reservationService) CheckFreeByRoomIdWithIntrval(roomId uuid.UUID, begin, end time.Time) (bool, error) {
	return s.reservationRepo.CheckFreeByRoomIdWithIntrval(roomId, begin, end)
}
