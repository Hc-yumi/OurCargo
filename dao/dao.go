package dao

import (
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	ID        int
	LineID    string
	Name      string
	CompanyID int
	Company   Company
	Role      string
	Email     string
	Tel       string
}

type Company struct {
	gorm.Model
	ID              int
	CompanyName     string
	CompanyLocation string
}

// type TruckType struct {
// 	gorm.Model
// 	ID          int
// 	UserID      int
// 	User        User
// 	TruckSizeID int
// 	TruckSize   TruckSize
// 	Price       string
// }

type TruckSize struct {
	gorm.Model
	ID        int
	TruckSize string
}

type Order struct {
	gorm.Model
	ID     int
	UserID int
	User   User
	// TruckTypeID int
	// TruckType       TruckType
	TruckSizeID     int
	TruckSize       TruckSize
	Price           string
	PickupLocation  string
	ArrivalLocation string
	PickupDatetime  time.Time
	ArrivalDatetime time.Time
	Mileage         int
	OrderDatetime   time.Time
}
