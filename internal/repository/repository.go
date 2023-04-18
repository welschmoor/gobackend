package repository

import (
	"database/sql"

	"github.com/welschmoor/gobackend/internal/models"
)

type DatabaseRepo interface {
	AllListings() ([]*models.Listing, error)
	Connection() *sql.DB
}
