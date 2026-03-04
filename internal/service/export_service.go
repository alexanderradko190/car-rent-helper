package service

import (
	"car-export-go/internal/entity"
	"car-export-go/internal/repository"
	"time"
)

type ExportService struct {
	repo repository.ExportRepository
}

func NewExportService(repo repository.ExportRepository) *ExportService {
	return &ExportService{repo: repo}
}

func (s *ExportService) RentHistories(from, to time.Time, page, perPage int) ([]entity.RentHistoryExportRow, int64, error) {
	return s.repo.GetRentHistories(from, to, page, perPage)
}

func (s *ExportService) RentalRequests(from, to time.Time, page, perPage int) ([]entity.RentalRequestExportRow, int64, error) {
	return s.repo.GetRentalRequests(from, to, page, perPage)
}