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

func (r *Repository) CreateConnection(conn models.Connection) (string, error) {
	query := `
		INSERT INTO connections (name)
		VALUES ($1)
		RETURNING name
	`

	var newConn string
	err := r.db.QueryRow(query, conn.Name).Scan(&newConn)
	if err != nil {
		return "", err
	}

	return newConn, nil
}
