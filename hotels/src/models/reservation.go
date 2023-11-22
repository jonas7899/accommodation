package models

import (
	"time"

	"github.com/google/uuid"
)

type ReservationDTO struct {
	Id               uuid.UUID `json:"id" db:"id"`
	HotelRoomId      uuid.UUID `json:"hotel_room_id" db:"hotel_room_id"`
	ReservationBegin time.Time `json:"reservation_begin" db:"reservation_begin"`
	ReservationEnd   time.Time `json:"reservation_end" db:"reservation_end"`
}

type ReservedRoomID struct {
	RoomId  uuid.UUID `json:"room_id" db:"room_id"`
	HotelId uuid.UUID `json:"hotel_id" db:"hotel_id"`
}
