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
		logger.Log(fmt.Sprintf("%s %s", conn.Name, conn.Url))
	}

	return nil
}

func (h *Handler) Create(param *parser.Param) error {
	name, ok := param.Next()
	if !ok {
		return errors.New("conn create <name> <url>")
	}

	url, ok := param.Next()
	if !ok {
		return errors.New("conn create <name> <url>")
	}

	conn := models.Connection{
		Name: name,
		Url:  url,
	}

	newConn, err := h.svc.CreateConnnection(conn)
	if err != nil {
		return err
	}

	logger.Log(fmt.Sprintf("Created %s", newConn))

	return nil
}

func (h *Handler) Remove(param *parser.Param) error {
	name, ok := param.Next()
	if !ok {
		return errors.New("conn remove <name>")
	}

	removedConn, err := h.svc.RemoveConnnection(name)
	if err != nil {
		return err
	}

	logger.Log(fmt.Sprintf("Removed %s", removedConn))

	return nil
}
