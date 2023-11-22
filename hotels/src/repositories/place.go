package repositories

import (
	"hotels/src/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PlaceRepository interface {
	CreatePlace(place models.PlaceDTO) (uuid.UUID, error)
	GetPlaces() ([]models.PlaceDTO, error)
	GetPlaceById(Id uuid.UUID) ([]models.PlaceDTO, error)
}

type placeRepository struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func (r *placeRepository) CreatePlace(place models.PlaceDTO) (uuid.UUID, error) {
	query := `INSERT INTO INSERT INTO hotels.place
	(id, place_name, postalcode) VALUES 
	(:id, :place_name, :postalcode);`

	if place.Id == uuid.Nil {
		place.Id = uuid.New()
	}

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &place)
	} else {
		_, err = r.db.NamedExec(query, &place)
	}

	return place.Id, err
}

func (r *placeRepository) GetPlaces() ([]models.PlaceDTO, error) {
	places := make([]models.PlaceDTO, 0)

	query := `SELECT id, place_name, postalcode 
	FROM hotels.place;`

	var err error
	var rows *sqlx.Rows
	if r.tx != nil {
		rows, err = r.tx.Queryx(query)
	} else {
		rows, err = r.db.Queryx(query)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		place := models.PlaceDTO{}
		err = rows.StructScan(&place)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *placeRepository) GetPlaceById(Id uuid.UUID) ([]models.PlaceDTO, error) {
	places := make([]models.PlaceDTO, 0)

	query := `SELECT id, place_name, postalcode 
	FROM hotels.place 
	WHERE id = $1;`

	var err error
	var rows *sqlx.Rows
	if r.tx != nil {
		rows, err = r.tx.Queryx(query, Id)
	} else {
		rows, err = r.db.Queryx(query, Id)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		place := models.PlaceDTO{}
		err = rows.StructScan(&place)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	return places, nil
}

func CreatePlaceRepository(db *sqlx.DB) PlaceRepository {
	return &placeRepository{
		db: db,
	}
}
