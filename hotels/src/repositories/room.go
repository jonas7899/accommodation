package repositories

import (
	"hotels/src/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RoomRepository interface {
	CreateRoom(room models.RoomDTO) (uuid.UUID, error)
	GetRooms(HotelId uuid.UUID) ([]models.RoomDTO, error)
	GetRoomById(Id uuid.UUID) ([]models.RoomDTO, error)
	GetSpacesWithBadByRoomId(Id uuid.UUID) ([]models.SpaceWithBadDTO, error)
	GetRoomAllDetailsById(Id uuid.UUID) ([]models.RoomDTO, error)
}

type roomRepository struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func (r *roomRepository) CreateRoom(room models.RoomDTO) (uuid.UUID, error) {
	query := `INSERT INTO INSERT INTO hotels.room
	(id, room_name, postalcode) VALUES 
	(:id, :room_name, :postalcode);`

	if room.Id == uuid.Nil {
		room.Id = uuid.New()
	}

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &room)
	} else {
		_, err = r.db.NamedExec(query, &room)
	}

	return room.Id, err
}

func (r *roomRepository) GetRooms(HotelId uuid.UUID) ([]models.RoomDTO, error) {
	rooms := make([]models.RoomDTO, 0)

	query := `SELECT id, hotel_id, room_name, room_size, air_conditioner, wifi, bathroom, iron 
	FROM hotels.hotel_room 
	WHERE hotel_id = $1;`

	var err error
	var rows *sqlx.Rows
	if r.tx != nil {
		rows, err = r.tx.Queryx(query, HotelId)
	} else {
		rows, err = r.db.Queryx(query, HotelId)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		room := models.RoomDTO{}
		err = rows.StructScan(&room)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (r *roomRepository) GetRoomById(Id uuid.UUID) ([]models.RoomDTO, error) {
	rooms := make([]models.RoomDTO, 0)

	query := `SELECT id, hotel_id, room_name, room_size, air_conditioner, wifi, bathroom, iron
	FROM hotels.hotel_room 
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
		room := models.RoomDTO{}
		err = rows.StructScan(&room)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (r *roomRepository) GetRoomAllDetailsById(Id uuid.UUID) ([]models.RoomDTO, error) {
	rooms := make([]models.RoomDTO, 0)

	query := `SELECT id, hotel_id, room_name, room_size, air_conditioner, wifi, bathroom, iron
	FROM hotels.hotel_room 
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
		room := models.RoomDTO{}
		err = rows.StructScan(&room)
		if err != nil {
			return nil, err
		}
		room.SpacessWithBad, err = r.GetSpacesWithBadByRoomId(room.Id)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (r *roomRepository) GetSpacesWithBadByRoomId(Id uuid.UUID) ([]models.SpaceWithBadDTO, error) {
	spaces := make([]models.SpaceWithBadDTO, 0)

	query := `SELECT distinct 
	hrs.space_name, hrb.hotel_room_space_id
	FROM hotels.hotel_room_space hrs
	INNER JOIN hotels.hotel_room_bad hrb ON hrb.hotel_room_space_id = hrs.id 
	WHERE hrs.hotel_room_id = $1;`

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
		space := models.SpaceWithBadDTO{}
		err = rows.StructScan(&space)
		if err != nil {
			return nil, err
		}
		space.Bads, err = r.GetBadsBySpaceId(space.Id)
		if err != nil {
			return nil, err
		}
		spaces = append(spaces, space)

	}
	return spaces, nil
}

func (r *roomRepository) GetBadsBySpaceId(Id uuid.UUID) ([]models.BadDTO, error) {
	bads := make([]models.BadDTO, 0)

	query := `SELECT
	hrb.bad_num, bt.type_name as bad_type_name, bt.sleeps  
	FROM hotels.hotel_room_bad hrb 
	INNER JOIN hotels.bad_type bt ON bt.id = hrb.bad_type_id 
	WHERE hrb.hotel_room_space_id = $1;`

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
		bad := models.BadDTO{}
		err = rows.StructScan(&bad)
		if err != nil {
			return nil, err
		}
		bads = append(bads, bad)
	}
	return bads, nil
}

func CreateRoomRepository(db *sqlx.DB) RoomRepository {
	return &roomRepository{
		db: db,
	}
}
