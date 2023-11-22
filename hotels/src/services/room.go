package services

import (
	"hotels/src/models"
	"hotels/src/repositories"

	"github.com/google/uuid"
)

type RoomService interface {
	CreateRoom(request models.RoomDTO) error
	GetRooms(hotelId uuid.UUID) ([]models.RoomDTO, error)
	GetRoomById(id uuid.UUID) ([]models.RoomDTO, error)
	GetSpacesWithBadByRoomId(Id uuid.UUID) ([]models.SpaceWithBadDTO, error)
	GetRoomAllDetailsById(Id uuid.UUID) ([]models.RoomDTO, error)
}

type roomService struct {
	roomRepo repositories.RoomRepository
}

func (s *roomService) GetRoomAllDetailsById(Id uuid.UUID) ([]models.RoomDTO, error) {
	return s.roomRepo.GetRoomAllDetailsById(Id)
}

func (s *roomService) GetSpacesWithBadByRoomId(Id uuid.UUID) ([]models.SpaceWithBadDTO, error) {
	return s.roomRepo.GetSpacesWithBadByRoomId(Id)
}

func (s *roomService) GetRooms(hotelId uuid.UUID) ([]models.RoomDTO, error) {
	return s.roomRepo.GetRooms(hotelId)
}

func (s *roomService) GetRoomById(id uuid.UUID) ([]models.RoomDTO, error) {
	return s.roomRepo.GetRoomById(id)
}

func (s *roomService) CreateRoom(room models.RoomDTO) error {
	_, err := s.roomRepo.CreateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func CreateRoomService(roomRepo repositories.RoomRepository) RoomService {
	return &roomService{
		roomRepo: roomRepo,
	}
}
