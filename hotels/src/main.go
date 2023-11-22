package main

import (
	"fmt"
	"hotels/src/api"
	"hotels/src/handlers"
	"hotels/src/repositories"
	"hotels/src/services"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func CreateDBConnection(host string, port int, user string, passw string, database string) (*sqlx.DB, error) {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode='disable'", host, port, user, passw, database)

	db, err := sqlx.Connect("postgres", url)
	if errors.Cause(err) != nil {
		return nil, err
	} else {
		err = db.Ping()
		if errors.Cause(err) != nil {
			return nil, err
		}
	}
	return db, nil
}

func main() {
	fmt.Println("Hello Hotels")

	db, err := CreateDBConnection("localhost", 5432, "postgres", "postgres", "postgres")

	if err != nil {
		err := errors.Wrap(err, err.Error())

		type stackTracer interface {
			StackTrace() errors.StackTrace
		}
		if errors.Cause(err) != nil {
			if err, ok := err.(stackTracer); ok {
				for _, f := range err.StackTrace() {
					fmt.Printf("%+s:%d\n", f, f)
				}
			}

			//log.Panic().Err(errs).Msgf("%s\n%s", "Failed to read configuration", errors.StackTrace(err))
			//fmt.Printf("%s", err.Error())
			panic(err)
		}

	}

	placeRepo := repositories.CreatePlaceRepository(db)
	placeService := services.CreatePlaceService(placeRepo)
	placeHandler := handlers.CreatePlaceHandler(placeService)

	addressRepo := repositories.CreateAddressRepository(db)
	addressService := services.CreateAddressService(addressRepo)
	addressHandler := handlers.CreateAddressHandler(addressService)

	hotelRepo := repositories.CreateHotelRepository(db)
	hotelService := services.CreateHotelService(hotelRepo)
	hotelHandler := handlers.CreateHotelHandler(hotelService)

	roomRepo := repositories.CreateRoomRepository(db)
	roomService := services.CreateRoomService(roomRepo)
	roomHandler := handlers.CreateRoomHandler(roomService)

	reservationRepo := repositories.CreateReservationRepository(db)
	rreservationService := services.CreateReservationService(reservationRepo)
	reservationHandler := handlers.CreateReservationHandler(rreservationService)

	api := api.CreateApi(
		placeHandler,
		addressHandler,
		hotelHandler,
		roomHandler,
		reservationHandler,
	)

	api.RunAPI()
	db.Close()
}
