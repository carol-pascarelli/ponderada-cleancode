package service

import (
	"errors"
	"ponderada3/domain"
	"ponderada3/repository"
	"time"

	"gorm.io/gorm"
)

var (
	ErrNotFound        = errors.New("gasto não encontrado")
	ErrInvalidAmount   = errors.New("o valor deve ser maior que zero")
	ErrInvalidCategory = errors.New("categoria inválida")
)

type ExpenseService interface {
	Create(req domain.CreateExpenseRequest) (*domain.Expense, error)
	List(category string) ([]domain.Expense, error)
	GetByID(id uint) (*domain.Expense, error)
	Update(id uint, req domain.UpdateExpenseRequest) (*domain.Expense, error)
	Delete(id uint) error
}

type expenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{repo: repo}
}

func (s *expenseService) Create(req domain.CreateExpenseRequest) (*domain.Expense, error) {
	if req.Amount <= 0 {
		return nil, ErrInvalidAmount
	}
	if !req.Category.IsValid() {
		return nil, ErrInvalidCategory
	}
	expense := &domain.Expense{
		Description: req.Description,
		Amount:      req.Amount,
		Category:    req.Category,
		Date:        time.Now(),
	}
	if err := s.repo.Create(expense); err != nil {
		return nil, err
	}
	return expense, nil
}

func (s *expenseService) List(category string) ([]domain.Expense, error) {
	if category != "" && !domain.ExpenseCategory(category).IsValid() {
		return nil, ErrInvalidCategory
	}
	return s.repo.FindAll(category)
}

func (s *expenseService) GetByID(id uint) (*domain.Expense, error) {
	expense, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return expense, err
}

func (s *expenseService) Update(id uint, req domain.UpdateExpenseRequest) (*domain.Expense, error) {
	expense, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if req.Amount != nil {
		if *req.Amount <= 0 {
			return nil, ErrInvalidAmount
		}
		expense.Amount = *req.Amount
	}
	if req.Category != nil {
		if !req.Category.IsValid() {
			return nil, ErrInvalidCategory
		}
		expense.Category = *req.Category
	}
	if req.Description != nil {
		expense.Description = *req.Description
	}
	if err := s.repo.Update(expense); err != nil {
		return nil, err
	}
	return expense, nil
}

func (s *expenseService) Delete(id uint) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
