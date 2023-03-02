package dao

import (
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB


type User struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	LineID		string
	Name        string
	CompanyID   int
	Company     Company
	Role        string
	Email       string
	Tel         int
	RegisterDay time.Time
	UpdateDay   time.Time
}

type Company struct {
	gorm.Model
	ID              int `gorm:"primary_key"`
	CompanyName     string
	CompanyLocation string
	RegisterDay     time.Time
	UpdateDay       time.Time
}

type TruckType struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	UserID      int
	User        User
	TruckSizeID int
	TruckSize   TruckSize
	Price       float64
	RegisterDay time.Time
	UpdateDay   time.Time
}

type TruckSize struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	TruckSize   string
	RegisterDay time.Time
	UpdateDay   time.Time
}

type Order struct {
	gorm.Model
	ID              uint `gorm:"primary_key"`
	UserID          uint
	User            User
	TruckTypeID     uint
	TruckType       TruckType
	TruckSizeID     uint
	TruckSize       TruckSize
	PickupLocation  string
	ArrivalLocation string
	PickupDatetime  time.Time
	ArrivalDatetime time.Time
	Mileage         int
	OrderDatetime   time.Time
	RegisterDay     time.Time
	UpdateDay       time.Time
}
