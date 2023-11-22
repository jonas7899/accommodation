package handlers

import (
	"hotels/src/models"
	"hotels/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoomHandler interface {
	GetRooms(c *gin.Context)
	CreateRoom(c *gin.Context)
	GetRoomById(c *gin.Context)
	UpdateRoom(c *gin.Context)
	DeleteRoom(c *gin.Context)
	GetSpacesWithBadByRoomId(c *gin.Context)
	GetRoomAllDetailsById(c *gin.Context)
}

type roomHandler struct {
	roomService services.RoomService
}

func (h *roomHandler) GetRooms(c *gin.Context) {
	hotelId, err := uuid.Parse(c.Params.ByName("hotelId"))
	if err != nil {
		return
	}
	rooms, err := h.roomService.GetRooms(hotelId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *roomHandler) CreateRoom(c *gin.Context) {
	roomData := models.RoomDTO{}
	err := c.BindJSON(&roomData)
	if err != nil {
		return
	}

	err = h.roomService.CreateRoom(roomData)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func (h *roomHandler) GetRoomById(c *gin.Context) {
	roomId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return
	}

	room, err := h.roomService.GetRoomById(roomId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, room)
}

func (h *roomHandler) UpdateRoom(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *roomHandler) DeleteRoom(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *roomHandler) GetSpacesWithBadByRoomId(c *gin.Context) {
	roomId, err := uuid.Parse(c.Param("roomid"))
	if err != nil {
		return
	}

	room, err := h.roomService.GetSpacesWithBadByRoomId(roomId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, room)
}

func (h *roomHandler) GetRoomAllDetailsById(c *gin.Context) {
	roomId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return
	}

	room, err := h.roomService.GetRoomAllDetailsById(roomId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, room)
}

func CreateRoomHandler(roomService services.RoomService) RoomHandler {
	return &roomHandler{
		roomService: roomService,
	}
}
