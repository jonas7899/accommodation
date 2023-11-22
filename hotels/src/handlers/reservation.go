package handlers

import (
	"hotels/src/models"
	"hotels/src/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReservationHandler interface {
	GetReservations(c *gin.Context)
	CreateReservation(c *gin.Context)
	GetReservationById(c *gin.Context)
	DeleteReservation(c *gin.Context)
	GetAllReservationsByRoomIdWithIntrval(c *gin.Context)
	GetAllReservationsByHotelIdWithIntrval(c *gin.Context)
	CheckFreeByRoomIdWithIntrval(c *gin.Context)
}

type reservationHandler struct {
	reservationService services.ReservationService
}

func (h *reservationHandler) GetReservations(c *gin.Context) {
	reservations, err := h.reservationService.GetReservations()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, reservations)
}

func (h *reservationHandler) CreateReservation(c *gin.Context) {
	reservationData := models.ReservationDTO{}
	err := c.BindJSON(&reservationData)
	if err != nil {
		return
	}

	err = h.reservationService.CreateReservation(reservationData)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func (h *reservationHandler) GetReservationById(c *gin.Context) {
	reservationId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return
	}

	reservation, err := h.reservationService.GetReservationById(reservationId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, reservation)
}

func (h *reservationHandler) DeleteReservation(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.reservationService.DeleteReservation(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *reservationHandler) GetAllReservationsByRoomIdWithIntrval(c *gin.Context) {

	roomId, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		return
	}

	rbegin, err := time.Parse(time.DateOnly, c.Params.ByName("begin"))
	if err != nil {
		return
	}
	rend, err := time.Parse(time.DateOnly, c.Params.ByName("end"))
	if err != nil {
		return
	}

	reservations, err := h.reservationService.GetAllReservationsByRoomIdWithIntrval(roomId, rbegin, rend)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, reservations)
}

func (h *reservationHandler) GetAllReservationsByHotelIdWithIntrval(c *gin.Context) {
	hotelId, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		return
	}

	rbegin, err := time.Parse(time.DateOnly, c.Params.ByName("begin"))
	if err != nil {
		return
	}
	rend, err := time.Parse(time.DateOnly, c.Params.ByName("end"))
	if err != nil {
		return
	}

	reservations, err := h.reservationService.GetAllReservationsByHotelIdWithIntrval(hotelId, rbegin, rend)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, reservations)
}

func (h *reservationHandler) CheckFreeByRoomIdWithIntrval(c *gin.Context) {
	roomId, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		return
	}

	rbegin, err := time.Parse(time.DateOnly, c.Params.ByName("begin"))
	if err != nil {
		return
	}
	rend, err := time.Parse(time.DateOnly, c.Params.ByName("end"))
	if err != nil {
		return
	}

	isFree, err := h.reservationService.CheckFreeByRoomIdWithIntrval(roomId, rbegin, rend)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, isFree)
}

func CreateReservationHandler(reservationService services.ReservationService) ReservationHandler {
	return &reservationHandler{
		reservationService: reservationService,
	}
}
