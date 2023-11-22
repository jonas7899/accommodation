package models

import (
	"github.com/google/uuid"
)

type PlaceDTO struct {
	Id         uuid.UUID `json:"id" db:"id"`
	PlaceName  string    `json:"place_name" db:"place_name"`
	PostalCode string    `json:"postalcode" db:"postalcode"`
}
