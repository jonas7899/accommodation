package handlers

import (
	"hotels/src/models"
	"hotels/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HotelHandler interface {
	GetHotels(c *gin.Context)
	GetHotelsWithType(c *gin.Context)
	CreateHotel(c *gin.Context)
	GetHotelById(c *gin.Context)
	GetHotelWithTypeById(c *gin.Context)
	UpdateHotel(c *gin.Context)
	DeleteHotel(c *gin.Context)
}

type hotelHandler struct {
	hotelService services.HotelService
}

func (h *hotelHandler) GetHotels(c *gin.Context) {
	hotels, err := h.hotelService.GetHotels()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, hotels)
}

func (h *hotelHandler) GetHotelsWithType(c *gin.Context) {
	hotels, err := h.hotelService.GetHotelsWithType()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, hotels)
}

func (h *hotelHandler) CreateHotel(c *gin.Context) {
	hotelData := models.HotelDTO{}
	err := c.BindJSON(&hotelData)
	if err != nil {
		return
	}

	err = h.hotelService.CreateHotel(hotelData)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func (h *hotelHandler) GetHotelById(c *gin.Context) {
	hotelId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return
	}

	place, err := h.hotelService.GetHotelById(hotelId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, place)
}

func (h *hotelHandler) GetHotelWithTypeById(c *gin.Context) {
	hotelId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return
	}

	place, err := h.hotelService.GetHotelWithTypeById(hotelId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, place)
}

func (h *hotelHandler) UpdateHotel(c *gin.Context) {
	updateData := models.HotelDTO{}
	err := c.BindJSON(&updateData)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.hotelService.UpdateHotel(updateData)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *hotelHandler) DeleteHotel(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.hotelService.DeleteHotel(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

func CreateHotelHandler(hotelService services.HotelService) HotelHandler {
	return &hotelHandler{
		hotelService: hotelService,
	}
}
