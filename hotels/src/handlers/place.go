package handlers

import (
	"hotels/src/models"
	"hotels/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PlaceHandler interface {
	GetPlaces(c *gin.Context)
	CreatePlace(c *gin.Context)
	GetPlaceById(c *gin.Context)
	UpdatePlace(c *gin.Context)
	DeletePlace(c *gin.Context)
}

type placeHandler struct {
	placeService services.PlaceService
}

func (h *placeHandler) GetPlaces(c *gin.Context) {
	places, err := h.placeService.GetPlaces()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, places)
}

func (h *placeHandler) CreatePlace(c *gin.Context) {
	placeData := models.PlaceDTO{}
	err := c.BindJSON(&placeData)
	if err != nil {
		return
	}

	err = h.placeService.CreatePlace(placeData)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func (h *placeHandler) GetPlaceById(c *gin.Context) {
	placeId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return
	}

	place, err := h.placeService.GetPlaceById(placeId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, place)
}

func (h *placeHandler) UpdatePlace(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *placeHandler) DeletePlace(c *gin.Context) {
	c.Status(http.StatusOK)
}

func CreatePlaceHandler(placeService services.PlaceService) PlaceHandler {
	return &placeHandler{
		placeService: placeService,
	}
}
