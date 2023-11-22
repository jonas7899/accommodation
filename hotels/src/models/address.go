package models

import (
	"github.com/google/uuid"
)

type AddressDTO struct {
	Id         uuid.UUID `json:"id" db:"id"`
	PlaceId    uuid.UUID `json:"place_id" db:"place_id"`
	Street     string    `json:"street" db:"street"`
	StreetType string    `json:"street_type" db:"street_type"`
	Address    string    `json:"address" db:"address"`
}

type AddressWithPlaceDTO struct {
	Id         uuid.UUID `json:"id" db:"id"`
	PostalCode string    `json:"postalcode" db:"postalcode"`
	PlaceName  string    `json:"place_name" db:"place_name"`
	Street     string    `json:"street" db:"street"`
	StreetType string    `json:"street_type" db:"street_type"`
	Address    string    `json:"address" db:"address"`
}
