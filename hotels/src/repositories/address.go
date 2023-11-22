package repositories

import (
	"hotels/src/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AddressRepository interface {
	CreateAddress(address models.AddressDTO) (uuid.UUID, error)
	GetAddresses() ([]models.AddressDTO, error)
	GetAddressesWithPlace() ([]models.AddressWithPlaceDTO, error)
	GetAddressById(Id uuid.UUID) ([]models.AddressDTO, error)
	GetAddressWithPlaceById(Id uuid.UUID) ([]models.AddressWithPlaceDTO, error)
	UpdateAddress(address models.AddressDTO) error
	DeleteAddress(Id uuid.UUID) error
}

type addressRepository struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func (r *addressRepository) CreateAddress(address models.AddressDTO) (uuid.UUID, error) {
	query := `INSERT INTO hotels.address
	(id, place_id, street, street_type, address) VALUES 
	(:id, :place_id, :street, :street_type, :address);`

	if address.Id == uuid.Nil {
		address.Id = uuid.New()
	}

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &address)
	} else {
		_, err = r.db.NamedExec(query, &address)
	}

	return address.Id, err
}

func (r *addressRepository) GetAddresses() ([]models.AddressDTO, error) {
	addresses := make([]models.AddressDTO, 0)

	query := `SELECT id, place_id, street, street_type, address 
	FROM hotels.address;`

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
		address := models.AddressDTO{}
		err = rows.StructScan(&address)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}
func (r *addressRepository) GetAddressesWithPlace() ([]models.AddressWithPlaceDTO, error) {
	addresses := make([]models.AddressWithPlaceDTO, 0)

	query := `SELECT a.id, p.postalcode as postalcode, p.place_name as place_name, street, street_type, address 
	FROM hotels.address a
	INNER JOIN hotels.place p ON a.place_id = p.id;`

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
		address := models.AddressWithPlaceDTO{}
		err = rows.StructScan(&address)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func (r *addressRepository) GetAddressById(Id uuid.UUID) ([]models.AddressDTO, error) {
	addresses := make([]models.AddressDTO, 0)

	query := `SELECT id, place_id, street, street_type, address 
	FROM hotels.address 
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
		address := models.AddressDTO{}
		err = rows.StructScan(&address)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}
func (r *addressRepository) GetAddressWithPlaceById(Id uuid.UUID) ([]models.AddressWithPlaceDTO, error) {
	addresses := make([]models.AddressWithPlaceDTO, 0)

	query := `SELECT a.id as id, p.postalcode as postalcode, p.place_name as place_name, street, street_type, address 
	FROM hotels.address a
	INNER JOIN hotels.place p ON a.place_id = p.id 
	WHERE a.id = $1;`

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
		address := models.AddressWithPlaceDTO{}
		err = rows.StructScan(&address)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func (r *addressRepository) UpdateAddress(address models.AddressDTO) error {
	query := `UPDATE hotels.address SET 
	place_id=:place_id,
	street=:street,
	street_type=:street_type,
	address=:address;
	WHERE id=:id;`

	if address.Id == uuid.Nil {
		address.Id = uuid.New()
	}

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &address)
	} else {
		_, err = r.db.NamedExec(query, &address)
	}

	return err
}

func (r *addressRepository) DeleteAddress(Id uuid.UUID) error {
	query := `DELETE FROM hotels.address 
	WHERE id=$1;`

	var err error
	if r.tx != nil {
		_, err = r.tx.NamedExec(query, &Id)
	} else {
		_, err = r.db.NamedExec(query, &Id)
	}

	return err
}

func CreateAddressRepository(db *sqlx.DB) AddressRepository {
	return &addressRepository{
		db: db,
	}
}
