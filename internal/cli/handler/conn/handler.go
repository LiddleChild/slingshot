package conn

import (
	"errors"
	"fmt"

	"github.com/LiddleChild/slingshot/internal/cli/parser"
	"github.com/LiddleChild/slingshot/internal/core/conn"
	"github.com/LiddleChild/slingshot/internal/core/models"
	"github.com/LiddleChild/slingshot/internal/util/logger"
)

type Handler struct {
	svc *conn.Service
}

func NewHandler(svc *conn.Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) List(param *parser.Param) error {
	connections, err := h.svc.GetAllConnections()
	if err != nil {
		return err
	}

	for _, conn := range connections {
		logger.Log(conn.Name)
	}

	return nil
}

func (h *Handler) Create(param *parser.Param) error {
	name, ok := param.Next()
	if !ok {
		return errors.New("conn create <name>")
	}

	conn := models.Connection{
		Name: name,
	}

	newConn, err := h.svc.CreateConnnection(conn)
	if err != nil {
		return err
	}

	logger.Log(fmt.Sprintf("Created %s", newConn))

	return nil
}
