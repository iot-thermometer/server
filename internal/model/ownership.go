package model

import (
	"time"
)

type Ownership struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	DeviceID  uint
	Role      OwnershipRole
}

type OwnershipRole int

const (
	OwnershipRoleOwner OwnershipRole = iota
	OwnershipRoleMember
)
