package models

import (
	"github.com/google/uuid"
)

type HotelDTO struct {
	Id        uuid.UUID `json:"id" db:"id"`
	HotelName string    `json:"hotel_name" db:"hotel_name"`
	TypeId    uuid.UUID `json:"type_id" db:"type_id"`
	Stars     int       `json:"stars" db:"stars"`
	Wellness  bool      `json:"wellness" db:"wellness"`
	Carpark   bool      `json:"carpark" db:"carpark"`
}

type HotelWithTypeDTO struct {
	Id        uuid.UUID `json:"id" db:"id"`
	HotelName string    `json:"hotel_name" db:"hotel_name"`
	TypeName  string    `json:"type_name" db:"type_name"`
	Stars     int       `json:"stars" db:"stars"`
	Wellness  bool      `json:"wellness" db:"wellness"`
	Carpark   bool      `json:"carpark" db:"carpark"`
}

type HotelTypeDTO struct {
	Id       uuid.UUID `json:"id" db:"id"`
	TypeName string    `json:"type_name" db:"type_name"`
}

type HotelAllRoomsDTO struct {
	HotelId        uuid.UUID `json:"hotel_id" db:"hotel_id"`
	HotelName      string    `json:"hotel_name" db:"hotel_name"`
	TypeName       string    `json:"type_name" db:"type_name"`
	Stars          int       `json:"stars" db:"stars"`
	Wellness       bool      `json:"wellness" db:"wellness"`
	Carpark        bool      `json:"carpark" db:"carpark"`
	RoomId         uuid.UUID `json:"room_id" db:"room_id"`
	RoomName       string    `json:"room_name" db:"room_name"`
	RoomSize       int       `json:"room_size" db:"room_size"`
	AirConditioner bool      `json:"air_conditioner" db:"air_conditioner"`
	Wifi           bool      `json:"wifi" db:"wifi"`
	Bathroom       bool      `json:"bathroom" db:"bathroom"`
	Iron           bool      `json:"iron" db:"iron"`
}
