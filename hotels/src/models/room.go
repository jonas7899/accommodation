package models

import (
	"github.com/google/uuid"
)

type RoomDTO struct {
	Id             uuid.UUID         `json:"id" db:"id"`
	HotelId        uuid.UUID         `json:"hotel_id" db:"hotel_id"`
	RoomName       string            `json:"room_name" db:"room_name"`
	RoomSize       int               `json:"room_size" db:"room_size"`
	AirConditioner bool              `json:"air_conditioner" db:"air_conditioner"`
	Wifi           bool              `json:"wifi" db:"wifi"`
	Bathroom       bool              `json:"bathroom" db:"bathroom"`
	Iron           bool              `json:"iron" db:"iron"`
	SpacessWithBad []SpaceWithBadDTO `json:"spacess_with_bad"`
}

type BadDTO struct {
	BadNum      int    `json:"bad_num" db:"bad_num"`
	BadTypeName string `json:"bad_type_name" db:"bad_type_name"`
	Sleeps      int    `json:"sleeps" db:"sleeps"`
}

type SpaceWithBadDTO struct {
	Id        uuid.UUID `json:"hotel_room_space_id" db:"hotel_room_space_id"`
	SpaceName string    `json:"space_name" db:"space_name"`
	Bads      []BadDTO  `json:"bads"`
}
