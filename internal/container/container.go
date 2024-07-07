package container

import (
	"database/sql"

	"github.com/LiddleChild/slingshot/internal/core/conn"
)

type Container struct {
	ConnService *conn.Service
}

func NewContainer(db *sql.DB) *Container {
	container := &Container{}

	connRepo := conn.NewRepository(db)
	connSvc := conn.NewService(connRepo)
	container.ConnService = connSvc

	return container
}
