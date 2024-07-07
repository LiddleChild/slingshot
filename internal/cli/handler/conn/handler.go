package conn

import (
	"fmt"

	"github.com/LiddleChild/slingshot/internal/cli/parser"
	"github.com/LiddleChild/slingshot/internal/core/conn"
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
		fmt.Println(conn.Name)
	}

	return nil
}
