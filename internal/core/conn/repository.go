package conn

import (
	"database/sql"

	"github.com/LiddleChild/slingshot/internal/core/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllConnections() ([]models.Connection, error) {
	connections := []models.Connection{}

	query := `SELECT * FROM connections`
	rows, err := r.db.Query(query)
	if err != nil {
		return []models.Connection{}, err
	}

	defer rows.Close()

	for rows.Next() {
		conn := models.Connection{}
		err := rows.Scan(&conn.Name)
		if err != nil {
			return []models.Connection{}, err
		}

		connections = append(connections, conn)
	}

	return connections, err
}
