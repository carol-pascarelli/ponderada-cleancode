package service 

import (
	"errors"
	"gorm.io/gorm"

	"ponderada-cleancode/domain"
	"ponderada-cleancode/repository"

)

var ErrFigurinhaNotFound = errors.New("figurinha não encontrada")
var ErrInvalidTipo = errors.New("tipo inválido")
var ErrInvalidPosicao = errors.New("posição inválida")
var	ErrCamposObrigatorios = errors.New("todos os campos são obrigatórios")


type FigurinhaService interface {
	CreateFigurinha(data *domain.CreateFigurinhaRequest) (*domain.Figurinha, error)
	GetFigurinhas(posicao, tipo string) ([]domain.Figurinha, error)
	GetFigurinhaByID(id uint) (*domain.Figurinha, error)
	UpdateFigurinha(id uint, data *domain.UpdateFigurinhaRequest) (*domain.Figurinha, error)
	DeleteFigurinha(id uint) error
}

type figurinhaService struct {
	repo repository.FigurinhaRepository
}

func NewFigurinhaService(repo repository.FigurinhaRepository) FigurinhaService {
	return &figurinhaService{
		repo: repo,
	}
}

func validarPosicao(posicao string) bool {
	for _, pnge domain.PosicoesValidas {
		if posicao == p {
			return true
		}
	}
}

func (s *figurinhaService) CreateFigurinha(data *domain.CreateFigurinhaRequest) (*domain.Figurinha, error) {
	if req.Numero == ||
	
}