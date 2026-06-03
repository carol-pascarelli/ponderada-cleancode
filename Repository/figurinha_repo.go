package repository

import (
	"ponderada-cleancode/domain"
	"gorm.io/gorm"
)

type FigurinhaRepository interface {
	Create(figurinha *domain.Figurinha) error
	FindAll(posicao, tipo string) ([]domain.Figurinha, error)
	FindByID(id uint) (*domain.Figurinha, error)
	Update(id uint, data *domain.Figurinha) error
	Delete(id uint) error
}

type figurinhaRepository struct {
	db *gorm.DB
}

func NewFigurinhaRepository(db *gorm.DB) FigurinhaRepository {
	return &figurinhaRepository{db: db}
}

func (r *figurinhaRepository) Create(figurinha *domain.Figurinha) error {
	return r.db.Create(figurinha).Error
}

func (r *figurinhaRepository) FindAll(posicao, tipo string) ([]domain.Figurinha, error) {
	var figurinhas []domain.Figurinha
	query := r.db.Order("created_at DESC")

	if posicao != "" {
		query = query.Where("posicao = ?", posicao)
	}
	if tipo != "" {
		query = query.Where("tipo = ?", tipo)
	}

	return figurinhas, query.Find(&figurinhas).Error
}

func (r *figurinhaRepository) FindByID(id uint) (*domain.Figurinha, error) {
	var figurinha domain.Figurinha

	err := r.db.First(&figurinha, id).Error
	if err != nil {
		return nil, err
	}

	return &figurinha, nil
}

func (r *figurinhaRepository) Update(id uint, data *domain.Figurinha) error {
	return r.db.Model(&domain.Figurinha{}).Where("id = ?", id).Updates(data).Error
}

func (r *figurinhaRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Figurinha{}, id).Error
}
