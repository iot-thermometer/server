package service

import (
	"github.com/iot-thermometer/server/internal/model"
	"github.com/iot-thermometer/server/internal/repository"
)

type Ownership interface {
	ListMembers(requesterID, deviceID uint) ([]*model.Ownership, error)
	AddMember(requesterID, deviceID uint, email string, role model.OwnershipRole) error
	RemoveMember(requesterID, deviceID uint, userID uint) error
}

type ownership struct {
	ownershipRepository repository.Ownership
	userRepository      repository.User
}

func newOwnershipService(
	ownershipRepository repository.Ownership,
	userRepository repository.User,
) Ownership {
	return &ownership{ownershipRepository, userRepository}
}

func (o *ownership) AddMember(requesterID, deviceID uint, email string, role model.OwnershipRole) error {
	user, err := o.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}
	userID := user.ID

	exists, err := o.ownershipRepository.Exists(userID, deviceID)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	_, err = o.ownershipRepository.Create(userID, deviceID, role)
	return err
}

func (o *ownership) ListMembers(requesterID, deviceID uint) ([]*model.Ownership, error) {
	ownerships, err := o.ownershipRepository.FindByDeviceID(deviceID)
	if err != nil {
		return nil, err
	}
	for _, ownership := range ownerships {
		user, err := o.userRepository.FindByID(ownership.UserID)
		if err != nil {
			return nil, err
		}
		ownership.User = &user
	}
	return ownerships, nil
}

func (o *ownership) RemoveMember(requesterID, deviceID uint, userID uint) error {
	exists, err := o.ownershipRepository.Exists(userID, deviceID)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	return o.ownershipRepository.Delete(userID, deviceID)
}
