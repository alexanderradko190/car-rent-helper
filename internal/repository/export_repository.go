package repository

import (
	"car-export-go/internal/entity"
	"time"

	"gorm.io/gorm"
)

type ExportRepository interface {
	GetRentHistories(from, to time.Time, page, perPage int) ([]entity.RentHistoryExportRow, int64, error)
	GetRentalRequests(from, to time.Time, page, perPage int) ([]entity.RentalRequestExportRow, int64, error)
}

type exportRepository struct {
	db *gorm.DB
}

func NewExportRepository(db *gorm.DB) ExportRepository {
	return &exportRepository{db: db}
}

func (r *exportRepository) GetRentHistories(from, to time.Time, page, perPage int) ([]entity.RentHistoryExportRow, int64, error) {
	var total int64
	countQuery := r.db.Table("rent_histories").
		Where("start_time >= ? AND start_time <= ?", from, to)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	// ВАЖНО: сканим в DTO, где есть vin/full_name/phone...
	var rows []entity.RentHistoryExportRow

	dataQuery := r.db.Table("rent_histories rh").
		Select(`
			rh.id AS id,
			c.vin AS vin,
			c.license_plate AS license_plate,
			cl.full_name AS full_name,
			cl.phone AS phone,
			rh.start_time AS start_time,
			rh.end_time AS end_time,
			rh.total_cost AS total_cost
		`).
		Joins("LEFT JOIN cars c ON c.id = rh.car_id").
		Joins("LEFT JOIN clients cl ON cl.id = rh.client_id").
		Where("rh.start_time >= ? AND rh.start_time <= ?", from, to).
		Order("rh.id ASC").
		Limit(perPage).
		Offset(offset)

	err := dataQuery.Scan(&rows).Error
	return rows, total, err
}

func (r *exportRepository) GetRentalRequests(from, to time.Time, page, perPage int) ([]entity.RentalRequestExportRow, int64, error) {
	var total int64
	countQuery := r.db.Table("rental_requests").
		Where("created_at >= ? AND created_at <= ?", from, to)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	var rows []entity.RentalRequestExportRow

	dataQuery := r.db.Table("rental_requests rr").
		Select(`
			rr.id AS id,
			c.vin AS vin,
			c.license_plate AS license_plate,
			cl.full_name AS full_name,
			cl.phone AS phone,
			rr.start_time AS start_time,
			rr.end_time AS end_time,
			rr.status AS status,
			rr.created_at AS created_at
		`).
		Joins("LEFT JOIN cars c ON c.id = rr.car_id").
		Joins("LEFT JOIN clients cl ON cl.id = rr.client_id").
		Where("rr.created_at >= ? AND rr.created_at <= ?", from, to).
		Order("rr.id ASC").
		Limit(perPage).
		Offset(offset)

	err := dataQuery.Scan(&rows).Error
	return rows, total, err
}