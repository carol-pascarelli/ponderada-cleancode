package service

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"ponderada-cleancode/domain"
	"ponderada-cleancode/repository"
)

var (
	ErrNotFound           = errors.New("figurinha não encontrado")
	ErrInvalidTipo        = errors.New("tipo inválido")
	ErrInvalidPosicao     = errors.New("posição inválida")
	ErrCamposObrigatorios = errors.New("todos os campos são obrigatórios")
)

type FigurinhaService interface {
	Create(req domain.CreateFigurinhaRequest) (*domain.Figurinha, error)
	List(posicao, tipo string) ([]domain.Figurinha, error)
	GetByID(id uint) (*domain.Figurinha, error)
	Update(id uint, req domain.UpdateFigurinhaRequest) (*domain.Figurinha, error)
	Delete(id uint) error
}

type figurinhaService struct {
	repo repository.FigurinhaRepository
}

func NewFigurinhaService(repo repository.FigurinhaRepository) FigurinhaService {
	return &figurinhaService{repo: repo}
}

func isValidTipo(tipo string) bool {
	for _, t := range domain.TiposValidos {
		if tipo == t {
			return true
		}
	}
	return false
}

func isValidPosicao(posicao string) bool {
	for _, p := range domain.PosicoesValidas {
		if posicao == p {
			return true
		}
	}
	return false
}

func (s *figurinhaService) Create(req domain.CreateFigurinhaRequest) (*domain.Figurinha, error) {
	if req.Numero == "" || req.Tipo == "" || req.Posicao == "" {
		return nil, ErrCamposObrigatorios
	}
	if !isValidTipo(req.Tipo) {
		return nil, ErrInvalidTipo
	}
	if !isValidPosicao(req.Posicao) {
		return nil, ErrInvalidPosicao
	}

	now := time.Now()
	figurinha := &domain.Figurinha{
		Numero:    req.Numero,
		Tipo:      req.Tipo,
		Posicao:   req.Posicao,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.repo.Create(figurinha); err != nil {
		return nil, err
	}
	return figurinha, nil
}

func (s *figurinhaService) List(posicao, tipo string) ([]domain.Figurinha, error) {
	if posicao != "" && !isValidPosicao(posicao) {
		return nil, ErrInvalidPosicao
	}
	if tipo != "" && !isValidTipo(tipo) {
		return nil, ErrInvalidTipo
	}
	return s.repo.FindAll(posicao, tipo)
}

func (s *figurinhaService) GetByID(id uint) (*domain.Figurinha, error) {
	figurinha, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return figurinha, err
}

func (s *figurinhaService) Update(id uint, req domain.UpdateFigurinhaRequest) (*domain.Figurinha, error) {
	figurinha, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if req.Numero != "" {
		figurinha.Numero = req.Numero
	}
	if req.Tipo != "" {
		if !isValidTipo(req.Tipo) {
			return nil, ErrInvalidTipo
		}
		figurinha.Tipo = req.Tipo
	}
	if req.Posicao != "" {
		if !isValidPosicao(req.Posicao) {
			return nil, ErrInvalidPosicao
		}
		figurinha.Posicao = req.Posicao
	}
	figurinha.UpdatedAt = time.Now()
	if err := s.repo.Update(id, figurinha); err != nil {
		return nil, err
	}
	return figurinha, nil
}

func (s *figurinhaService) Delete(id uint) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
