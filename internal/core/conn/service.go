package conn

import "github.com/LiddleChild/slingshot/internal/core/models"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAllConnections() ([]models.Connection, error) {
	return s.repo.GetAllConnections()
}
