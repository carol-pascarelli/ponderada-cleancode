package repository

import (
	"ponderada3/domain"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Create(expense *domain.Expense) error
	FindAll(category string) ([]domain.Expense, error)
	FindByID(id uint) (*domain.Expense, error)
	Update(expense *domain.Expense) error
	Delete(id uint) error
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) Create(expense *domain.Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepository) FindAll(category string) ([]domain.Expense, error) {
	var expenses []domain.Expense
	q := r.db.Order("created_at DESC")
	if category != "" {
		q = q.Where("category = ?", category)
	}
	return expenses, q.Find(&expenses).Error
}

func (r *expenseRepository) FindByID(id uint) (*domain.Expense, error) {
	var expense domain.Expense
	err := r.db.First(&expense, id).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *expenseRepository) Update(expense *domain.Expense) error {
	return r.db.Save(expense).Error
}

func (r *expenseRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Expense{}, id).Error
}
