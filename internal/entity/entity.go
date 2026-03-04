package entity

import "time"

type Car struct {
	ID           uint   `gorm:"primaryKey;column:id"`
	Vin          string `gorm:"column:vin"`
	LicensePlate string `gorm:"column:license_plate"`
}

func (Car) TableName() string { return "cars" }

type Client struct {
	ID       uint   `gorm:"primaryKey;column:id"`
	FullName string `gorm:"column:full_name"`
	Phone    string `gorm:"column:phone"`
}

func (Client) TableName() string { return "clients" }

type RentHistory struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	ClientID  uint      `gorm:"column:client_id"`
	CarID     uint      `gorm:"column:car_id"`
	StartTime time.Time `gorm:"column:start_time"`
	EndTime   time.Time `gorm:"column:end_time"`
	TotalCost float64   `gorm:"column:total_cost"`
}

func (RentHistory) TableName() string { return "rent_histories" }

type RentalRequest struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	ClientID  uint      `gorm:"column:client_id"`
	CarID     uint      `gorm:"column:car_id"`
	StartTime time.Time `gorm:"column:start_time"`
	EndTime   time.Time `gorm:"column:end_time"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (RentalRequest) TableName() string { return "rental_requests" }