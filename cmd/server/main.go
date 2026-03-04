package main

import (
	"car-export-go/internal/config"
	"car-export-go/internal/repository"
	"car-export-go/internal/service"
	"car-export-go/internal/storage"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	db, err := storage.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewExportRepository(db)
	svc := service.NewExportService(repo)

	http.HandleFunc("/api/export/rent_histories", func(w http.ResponseWriter, r *http.Request) {
		from, to, page, perPage, ok := parseParams(w, r)
		if !ok {
			return
		}

		rows, total, err := svc.RentHistories(from, to, page, perPage)
		if err != nil {
			http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		writePagedJSON(w, rows, total, from, to, page, perPage)
	})

	http.HandleFunc("/api/export/rental_requests", func(w http.ResponseWriter, r *http.Request) {
		from, to, page, perPage, ok := parseParams(w, r)
		if !ok {
			return
		}

		rows, total, err := svc.RentalRequests(from, to, page, perPage)
		if err != nil {
			http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		writePagedJSON(w, rows, total, from, to, page, perPage)
	})

	log.Println("Go export service listening on :8002")
	log.Fatal(http.ListenAndServe(":8002", nil))
}

func parseParams(w http.ResponseWriter, r *http.Request) (time.Time, time.Time, int, int, bool) {
	q := r.URL.Query()

	fromStr := q.Get("from")
	if fromStr == "" {
		fromStr = q.Get("date_from")
	}
	toStr := q.Get("to")
	if toStr == "" {
		toStr = q.Get("date_to")
	}

	if fromStr == "" || toStr == "" {
		http.Error(w, "from and to are required (YYYY-MM-DD)", http.StatusBadRequest)
		return time.Time{}, time.Time{}, 0, 0, false
	}

	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		http.Error(w, "invalid from (YYYY-MM-DD)", http.StatusBadRequest)
		return time.Time{}, time.Time{}, 0, 0, false
	}
	to, err := time.Parse("2006-01-02", toStr)
	if err != nil {
		http.Error(w, "invalid to (YYYY-MM-DD)", http.StatusBadRequest)
		return time.Time{}, time.Time{}, 0, 0, false
	}

	page := atoiOr(q.Get("page"), 1)
	if page < 1 {
		page = 1
	}
	perPage := atoiOr(q.Get("per_page"), 1000)
	if perPage < 1 {
		perPage = 1000
	}

	return from, to, page, perPage, true
}

func atoiOr(v string, def int) int {
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return i
}

func writePagedJSON(w http.ResponseWriter, rows any, total int64, from, to time.Time, page, perPage int) {
	hasMore := int64(page*perPage) < total

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"data": rows,
		"meta": map[string]any{
			"from":     from.Format("2006-01-02"),
			"to":       to.Format("2006-01-02"),
			"page":     page,
			"per_page": perPage,
			"total":    total,
			"has_more": hasMore,
		},
	})
}