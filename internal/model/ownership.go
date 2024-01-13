package model

import "gorm.io/gorm"

type Ownership struct {
	gorm.Model
	UserID   uint
	DeviceID uint
	Role     OwnershipRole
	User     *User `gorm:"-"`
}

type OwnershipRole int

const (
	OwnershipRoleOwner OwnershipRole = iota
	OwnershipRoleMember
)
