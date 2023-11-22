package handlers

import (
	"hotels/src/models"
	"hotels/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddressHandler interface {
	GetAddresses(c *gin.Context)
	CreateAddress(c *gin.Context)
	GetAddressById(c *gin.Context)
	UpdateAddress(c *gin.Context)
	DeleteAddress(c *gin.Context)
}

type addressHandler struct {
	addressService services.AddressService
}

func (h *addressHandler) GetAddresses(c *gin.Context) {
	addresses, err := h.addressService.GetAddresses()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, addresses)
}

func (h *addressHandler) CreateAddress(c *gin.Context) {
	addressData := models.AddressDTO{}
	err := c.BindJSON(&addressData)
	if err != nil {
		return
	}

	err = h.addressService.CreateAddress(addressData)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Writer.WriteHeaderNow()
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func (h *addressHandler) GetAddressById(c *gin.Context) {
	addressId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return
	}

	place, err := h.addressService.GetAddressById(addressId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, place)
}

func (h *addressHandler) UpdateAddress(c *gin.Context) {
	updateData := models.AddressDTO{}
	err := c.BindJSON(&updateData)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.addressService.UpdateAddress(updateData)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *addressHandler) DeleteAddress(c *gin.Context) {
	addressId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.addressService.DeleteAddress(addressId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

func CreateAddressHandler(addressService services.AddressService) AddressHandler {
	return &addressHandler{
		addressService: addressService,
	}
}
