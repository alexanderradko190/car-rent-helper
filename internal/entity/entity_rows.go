package entity

import "time"

type RentHistoryExportRow struct {
	ID           uint      `json:"id" gorm:"column:id"`
	Vin          string    `json:"vin" gorm:"column:vin"`
	LicensePlate string    `json:"license_plate" gorm:"column:license_plate"`
	FullName     string    `json:"full_name" gorm:"column:full_name"`
	Phone        string    `json:"phone" gorm:"column:phone"`
	StartTime    time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime      time.Time `json:"end_time" gorm:"column:end_time"`
	TotalCost    float64   `json:"total_cost" gorm:"column:total_cost"`
}

type RentalRequestExportRow struct {
	ID           uint      `json:"id" gorm:"column:id"`
	Vin          string    `json:"vin" gorm:"column:vin"`
	LicensePlate string    `json:"license_plate" gorm:"column:license_plate"`
	FullName     string    `json:"full_name" gorm:"column:full_name"`
	Phone        string    `json:"phone" gorm:"column:phone"`
	StartTime    time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime      time.Time `json:"end_time" gorm:"column:end_time"`
	Status       string    `json:"status" gorm:"column:status"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
}