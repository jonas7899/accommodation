package repositories

import (
	"fmt"
	"hotels/src/models"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ReservationRepository interface {
	CreateReservation(reservation models.ReservationDTO) (uuid.UUID, error)
	GetReservations() ([]models.ReservationDTO, error)
	GetReservationById(Id uuid.UUID) ([]models.ReservationDTO, error)
	DeleteReservation(Id uuid.UUID) error
	GetAllReservationsByRoomIdWithIntrval(roomId uuid.UUID, begin, end time.Time) ([]uuid.UUID, error)
	GetAllReservationsByHotelIdWithIntrval(hotelId uuid.UUID, begin, end time.Time) ([]models.ReservedRoomID, error)
	CheckFreeByRoomIdWithIntrval(roomId uuid.UUID, begin, end time.Time) (bool, error)
}

type reservationRepository struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func (r *reservationRepository) CreateReservation(reservation models.ReservationDTO) (uuid.UUID, error) {
	query := `INSERT INTO hotels.reservation
	(id, hotel_room_id, reservation_begin, reservation_end)
	VALUES(:id, :hotel_room_id, :reservation_begin, :reservation_end);`

	if reservation.Id == uuid.Nil {
		reservation.Id = uuid.New()
	}

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &reservation)
	} else {
		_, err = r.db.NamedExec(query, &reservation)
	}

	return reservation.Id, err
}

func (r *reservationRepository) GetReservations() ([]models.ReservationDTO, error) {
	reservations := make([]models.ReservationDTO, 0)

	query := `SELECT id, hotel_room_id, reservation_begin, reservation_end 
	FROM hotels.reservation;`

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
		reservation := models.ReservationDTO{}
		err = rows.StructScan(&reservation)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

func (r *reservationRepository) GetReservationById(Id uuid.UUID) ([]models.ReservationDTO, error) {
	reservations := make([]models.ReservationDTO, 0)

	query := `SELECT id, hotel_room_id, reservation_begin, reservation_end 
	FROM hotels.reservation 
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
		reservation := models.ReservationDTO{}
		err = rows.StructScan(&reservation)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

func (r *reservationRepository) DeleteReservation(Id uuid.UUID) error {
	query := `DELETE 
	FROM hotels.reservation 
	WHERE id=$1;`

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &Id)
	} else {
		_, err = r.db.NamedExec(query, &Id)
	}

	return err
}

func (r *reservationRepository) GetAllReservationsByRoomIdWithIntrval(roomId uuid.UUID, rbegin time.Time, rend time.Time) ([]uuid.UUID, error) {
	if !rend.After(rbegin) {
		return nil, fmt.Errorf("time intervall error 'end' (%v) must be gater then 'begin' (%v)", rend, rbegin)
	}

	reservedRoomIds := make([]uuid.UUID, 0)

	query := `SELECT hotel_room_id 
	FROM hotels.reservation 
	WHERE
		($2 BETWEEN reservation_begin AND reservation_end OR 
			$3 BETWEEN reservation_begin AND reservation_end OR 
			reservation_begin between $2 and $3 OR 
			reservation_end between $2 and $3) and
		hotel_room_id = $1;`

	var err error
	var rows *sqlx.Rows
	if r.tx != nil {
		rows, err = r.tx.Queryx(query, roomId, rbegin, rend)
	} else {
		rows, err = r.db.Queryx(query, roomId, rbegin, rend)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		roomId := ""
		err = rows.Scan(&roomId)
		if err != nil {
			return nil, err
		}
		id, err := uuid.Parse(roomId)
		if err != nil {
			return nil, err
		}
		reservedRoomIds = append(reservedRoomIds, id)
	}
	return reservedRoomIds, nil
}

func (r *reservationRepository) GetAllReservationsByHotelIdWithIntrval(hotelId uuid.UUID, rbegin time.Time, rend time.Time) ([]models.ReservedRoomID, error) {

	if !rend.After(rbegin) {
		return nil, fmt.Errorf("time intervall error 'end' (%v) must be gater then 'begin' (%v)", rend, rbegin)
	}
	reservedRooms := make([]models.ReservedRoomID, 0)

	query := `SELECT re.hotel_room_id as room_id, ro.hotel_id as hotel_id 
	FROM hotels.reservation re
	INNER JOIN hotels.hotel_room ro ON ro.id = re.hotel_room_id 
	WHERE
		($2 BETWEEN reservation_begin AND reservation_end OR 
			$3 BETWEEN reservation_begin AND reservation_end OR 
			reservation_begin between $2 and $3 OR 
			reservation_end between $2 and $3) and
			ro.hotel_id = $1;`

	var err error
	var rows *sqlx.Rows
	if r.tx != nil {
		rows, err = r.tx.Queryx(query, hotelId, rbegin, rend)
	} else {
		rows, err = r.db.Queryx(query, hotelId, rbegin, rend)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var roomId string
		var hotelId string

		err := rows.Scan(&roomId, &hotelId)
		if err != nil {
			return nil, err
		}

		reservedRoom := models.ReservedRoomID{}
		reservedRoom.RoomId, err = uuid.Parse(roomId)
		if err != nil {
			return nil, err
		}
		reservedRoom.HotelId, err = uuid.Parse(hotelId)
		if err != nil {
			return nil, err
		}

		//err = rows.StructScan(&reservedRoom)

		reservedRooms = append(reservedRooms, reservedRoom)
	}
	return reservedRooms, nil
}

func (r *reservationRepository) CheckFreeByRoomIdWithIntrval(roomId uuid.UUID, begin, end time.Time) (bool, error) {
	reservations, err := r.GetAllReservationsByRoomIdWithIntrval(roomId, begin, end)
	if err != nil {
		return false, err
	}
	if len(reservations) == 0 {
		return true, err
	}
	return false, err
}

func CreateReservationRepository(db *sqlx.DB) ReservationRepository {
	return &reservationRepository{
		db: db,
	}
}
