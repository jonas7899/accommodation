package api

import (
	"hotels/src/handlers"

	"github.com/gin-gonic/gin"
)

type Api interface {
	RunAPI()
}

type api struct {
	placeHandler       handlers.PlaceHandler
	addressHandler     handlers.AddressHandler
	hotelHandler       handlers.HotelHandler
	roomHandler        handlers.RoomHandler
	reservationHandler handlers.ReservationHandler
}

func (api *api) RunAPI() {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(gin.Recovery())

	root := engine.Group("")

	v1Group := root.Group("v1")
	{
		placeGroup := v1Group.Group("/place")
		{
			placeGroup.GET("", api.placeHandler.GetPlaces)
			placeGroup.POST("", api.placeHandler.CreatePlace)
			placeGroup.GET("/:id", api.placeHandler.GetPlaceById)
			placeGroup.PUT("/:id", api.placeHandler.UpdatePlace)
			placeGroup.DELETE("/:id", api.placeHandler.DeletePlace)
		}

		addressGroup := v1Group.Group("/address")
		{
			addressGroup.GET("", api.addressHandler.GetAddresses)
			addressGroup.POST("", api.addressHandler.CreateAddress)
			addressGroup.GET("/:id", api.addressHandler.GetAddressById)
			addressGroup.PUT("/:id", api.addressHandler.UpdateAddress)
			addressGroup.DELETE("/:id", api.addressHandler.DeleteAddress)
		}

		hotelGroup := v1Group.Group("/hotel")
		{
			hotelGroup.GET("", api.hotelHandler.GetHotels)
			hotelGroup.POST("", api.hotelHandler.CreateHotel)
			hotelGroup.GET("/:id", api.hotelHandler.GetHotelById)
			hotelGroup.PUT("/:id", api.hotelHandler.UpdateHotel)
			hotelGroup.DELETE("/:id", api.hotelHandler.DeleteHotel)
		}

		roomGroup := v1Group.Group("/room")
		{
			roomGroup.GET("", api.roomHandler.GetRooms)
			roomGroup.POST("", api.roomHandler.CreateRoom)
			roomGroup.GET("/:id", api.roomHandler.GetRoomById)
			roomGroup.PUT("/:id", api.roomHandler.UpdateRoom)
			roomGroup.DELETE("/:id", api.roomHandler.DeleteRoom)
			roomGroup.GET("/screen/:id", api.roomHandler.GetSpacesWithBadByRoomId)
			roomGroup.GET("/details/:id", api.roomHandler.GetRoomAllDetailsById)
		}

		reservationGroup := v1Group.Group("/reservation")
		{
			reservationGroup.GET("", api.reservationHandler.GetReservations)
			reservationGroup.POST("", api.reservationHandler.CreateReservation)
			reservationGroup.GET("/:id", api.reservationHandler.GetReservationById)
			reservationGroup.DELETE("/:id", api.reservationHandler.DeleteReservation)
			reservationGroup.GET("/check/:id/:begin/:end", api.reservationHandler.CheckFreeByRoomIdWithIntrval)
			reservationGroup.GET("/hotel/:id/:begin/:end", api.reservationHandler.GetAllReservationsByHotelIdWithIntrval)
			reservationGroup.GET("/room/:id/:begin/:end", api.reservationHandler.GetAllReservationsByRoomIdWithIntrval)
		}

	}

	err := engine.Run(":8080")
	if err != nil {
		return
	}

}

func CreateApi(
	placeHandler handlers.PlaceHandler,
	addressHandler handlers.AddressHandler,
	hotelHandler handlers.HotelHandler,
	roomHandler handlers.RoomHandler,
	reservationHandler handlers.ReservationHandler,
) Api {
	return &api{
		placeHandler:       placeHandler,
		addressHandler:     addressHandler,
		hotelHandler:       hotelHandler,
		roomHandler:        roomHandler,
		reservationHandler: reservationHandler,
	}
}
