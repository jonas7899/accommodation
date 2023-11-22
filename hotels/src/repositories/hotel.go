package repositories

import (
	"hotels/src/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type HotelRepository interface {
	CreateHotel(hotel models.HotelDTO) (uuid.UUID, error)
	GetHotels() ([]models.HotelDTO, error)
	GetHotelsWithType() ([]models.HotelWithTypeDTO, error)
	GetHotelById(Id uuid.UUID) ([]models.HotelDTO, error)
	GetHotelWithTypeById(Id uuid.UUID) ([]models.HotelWithTypeDTO, error)
	UpdateHotel(hotel models.HotelDTO) error
	DeleteHotel(Id uuid.UUID) error
}

type hotelRepository struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func (r *hotelRepository) CreateHotel(hotel models.HotelDTO) (uuid.UUID, error) {
	query := `INSERT INTO hotels.hotel
	(id, hotel_name, type_id, stars, wellness, carpark) VALUES 
	(:id, :hotel_name, :type_id, :stars, :wellness, :carpark);`

	if hotel.Id == uuid.Nil {
		hotel.Id = uuid.New()
	}

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &hotel)
	} else {
		_, err = r.db.NamedExec(query, &hotel)
	}

	return hotel.Id, err
}

func (r *hotelRepository) GetHotelTypes() ([]models.HotelTypeDTO, error) {
	hotelTypes := make([]models.HotelTypeDTO, 0)

	query := `SELECT id, type_name 
	FROM hotels.hotel_type;`

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
		hotelType := models.HotelTypeDTO{}
		err = rows.StructScan(&hotelType)
		if err != nil {
			return nil, err
		}
		hotelTypes = append(hotelTypes, hotelType)
	}
	return hotelTypes, nil
}

func (r *hotelRepository) GetHotels() ([]models.HotelDTO, error) {
	hotels := make([]models.HotelDTO, 0)

	query := `SELECT id, hotel_name, type_id, stars, wellness, carpark 
	FROM hotels.hotel;`

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
		hotel := models.HotelDTO{}
		err = rows.StructScan(&hotel)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (r *hotelRepository) GetHotelsWithType() ([]models.HotelWithTypeDTO, error) {
	hotels := make([]models.HotelWithTypeDTO, 0)

	query := `SELECT h.id as id, hotel_name, t.type_name as type_name, stars, wellness, carpark
	FROM hotels.hotel h
	INNER JOIN hotels.hotel_type t ON t.id = h.type_id;`

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
		hotel := models.HotelWithTypeDTO{}
		err = rows.StructScan(&hotel)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (r *hotelRepository) GetHotelById(Id uuid.UUID) ([]models.HotelDTO, error) {
	hotels := make([]models.HotelDTO, 0)

	query := `SELECT id, hotel_name, type_id, stars, wellness, carpark
	FROM hotels.hotel 
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
		hotel := models.HotelDTO{}
		err = rows.StructScan(&hotel)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (r *hotelRepository) GetHotelWithTypeById(Id uuid.UUID) ([]models.HotelWithTypeDTO, error) {
	hotels := make([]models.HotelWithTypeDTO, 0)

	query := `SELECT h.id as id, hotel_name, t.type_name as type_name, stars, wellness, carpark
	FROM hotels.hotel h
	INNER JOIN hotels.hotel_type t ON t.id = h.type_id
	WHERE h.id = $1;`

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
		hotel := models.HotelWithTypeDTO{}
		err = rows.StructScan(&hotel)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (r *hotelRepository) UpdateHotel(hotel models.HotelDTO) error {
	query := `UPDATE hotels.hotel SET 
	hotel_name=hotel_name, 
	type_id=:type_id, 
	stars=:stars, 
	wellness=:wellness, 
	carpark=:carpark
	WHERE id=:id;`

	if hotel.Id == uuid.Nil {
		hotel.Id = uuid.New()
	}

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &hotel)
	} else {
		_, err = r.db.NamedExec(query, &hotel)
	}

	return err
}

func (r *hotelRepository) DeleteHotel(Id uuid.UUID) error {
	query := `DELETE 
	FROM hotels.hotel 
	WHERE id=$1;`

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &Id)
	} else {
		_, err = r.db.NamedExec(query, &Id)
	}

	return err
}

func (r *hotelRepository) GetHotelAllRoomsById(Id uuid.UUID) ([]models.HotelAllRoomsDTO, error) {
	rooms := make([]models.HotelAllRoomsDTO, 0)

	query := `SELECT h.id as hotel_id, h.hotel_name, t.type_name as type_name, h.stars, h.wellness, h.carpark, 
	r.id as room_id, r.room_name, r.room_size, r.air_conditioner, r.wifi,r.bathroom, r.iron
	FROM hotels.hotel h
	INNER JOIN hotels.hotel_type t ON t.id = h.type_id
	INNER JOIN hotels.hotel_room r ON r.hotel_id = h.id 
	WHERE h.id = $1;`

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
		room := models.HotelAllRoomsDTO{}
		err = rows.StructScan(&room)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func CreateHotelRepository(db *sqlx.DB) HotelRepository {
	return &hotelRepository{
		db: db,
	}
}
