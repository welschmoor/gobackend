package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/welschmoor/gobackend/internal/models"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) AllListings() ([]*models.Listing, error) {
	var listings []*models.Listing

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, title, description, created_at
		FROM movies ORDER BY title;
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var listing models.Listing
		err := rows.Scan(
			&listing.ID,
			&listing.Title,
			&listing.Description,
			&listing.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		listings = append(listings, &listing)
	}

	return listings, nil
}

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}
