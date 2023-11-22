package services

import (
	"hotels/src/models"
	"hotels/src/repositories"

	"github.com/google/uuid"
)

type AddressService interface {
	CreateAddress(request models.AddressDTO) error
	GetAddresses() ([]models.AddressDTO, error)
	GetAddressesWithPlace() ([]models.AddressWithPlaceDTO, error)
	GetAddressById(Id uuid.UUID) ([]models.AddressDTO, error)
	GetAddressWithPlaceById(Id uuid.UUID) ([]models.AddressWithPlaceDTO, error)
	UpdateAddress(models.AddressDTO) error
	DeleteAddress(Id uuid.UUID) error
}

type addressService struct {
	addressRepo repositories.AddressRepository
}

func CreateAddressService(addressRepo repositories.AddressRepository) AddressService {
	return &addressService{
		addressRepo: addressRepo,
	}
}

func (s *addressService) CreateAddress(address models.AddressDTO) error {
	_, err := s.addressRepo.CreateAddress(address)
	if err != nil {
		return err
	}
	return nil
}

func (s *addressService) GetAddresses() ([]models.AddressDTO, error) {
	return s.addressRepo.GetAddresses()
}

func (s *addressService) GetAddressesWithPlace() ([]models.AddressWithPlaceDTO, error) {
	return s.addressRepo.GetAddressesWithPlace()
}

func (s *addressService) GetAddressById(Id uuid.UUID) ([]models.AddressDTO, error) {
	return s.addressRepo.GetAddressById(Id)
}

func (s *addressService) GetAddressWithPlaceById(Id uuid.UUID) ([]models.AddressWithPlaceDTO, error) {
	return s.addressRepo.GetAddressWithPlaceById(Id)
}

func (s *addressService) UpdateAddress(address models.AddressDTO) error {
	err := s.addressRepo.UpdateAddress(address)
	if err != nil {
		return err
	}
	return nil
}

func (s *addressService) DeleteAddress(Id uuid.UUID) error {
	return s.addressRepo.DeleteAddress(Id)
}
