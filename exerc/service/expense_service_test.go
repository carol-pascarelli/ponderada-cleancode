package service_test

import (
	"errors"
	"ponderada3/domain"
	"ponderada3/service"
	"testing"

	"gorm.io/gorm"
)

// fakeRepo implementa repository.ExpenseRepository em memória.
type fakeRepo struct {
	expenses map[uint]*domain.Expense
	next     uint
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{expenses: make(map[uint]*domain.Expense), next: 1}
}

func (f *fakeRepo) Create(e *domain.Expense) error {
	e.ID = f.next
	f.next++
	cp := *e
	f.expenses[cp.ID] = &cp
	return nil
}

func (f *fakeRepo) FindAll(category string) ([]domain.Expense, error) {
	var out []domain.Expense
	for _, e := range f.expenses {
		if category == "" || string(e.Category) == category {
			out = append(out, *e)
		}
	}
	return out, nil
}

func (f *fakeRepo) FindByID(id uint) (*domain.Expense, error) {
	e, ok := f.expenses[id]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	cp := *e
	return &cp, nil
}

func (f *fakeRepo) Update(e *domain.Expense) error {
	f.expenses[e.ID] = e
	return nil
}

func (f *fakeRepo) Delete(id uint) error {
	delete(f.expenses, id)
	return nil
}

func newSvc() service.ExpenseService {
	return service.NewExpenseService(newFakeRepo())
}

func TestCreate_ValidRequest(t *testing.T) {
	svc := newSvc()
	req := domain.CreateExpenseRequest{
		Description: "Almoço",
		Amount:      25.50,
		Category:    domain.CategoryFood,
	}
	expense, err := svc.Create(req)
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}
	if expense.ID == 0 {
		t.Error("ID não foi preenchido")
	}
	if expense.Date.IsZero() {
		t.Error("Date não foi preenchida")
	}
}

func TestCreate_InvalidAmount(t *testing.T) {
	svc := newSvc()
	req := domain.CreateExpenseRequest{
		Description: "Almoço",
		Amount:      -1,
		Category:    domain.CategoryFood,
	}
	_, err := svc.Create(req)
	if !errors.Is(err, service.ErrInvalidAmount) {
		t.Fatalf("esperava ErrInvalidAmount, obteve %v", err)
	}
}

func TestCreate_InvalidCategory(t *testing.T) {
	svc := newSvc()
	req := domain.CreateExpenseRequest{
		Description: "Almoço",
		Amount:      10,
		Category:    "lazer",
	}
	_, err := svc.Create(req)
	if !errors.Is(err, service.ErrInvalidCategory) {
		t.Fatalf("esperava ErrInvalidCategory, obteve %v", err)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	svc := newSvc()
	_, err := svc.GetByID(999)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("esperava ErrNotFound, obteve %v", err)
	}
}

func TestList_FilterByCategory(t *testing.T) {
	svc := newSvc()
	svc.Create(domain.CreateExpenseRequest{Description: "Onibus", Amount: 5, Category: domain.CategoryTransport})
	svc.Create(domain.CreateExpenseRequest{Description: "Almoço", Amount: 20, Category: domain.CategoryFood})

	results, err := svc.List(string(domain.CategoryTransport))
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 {
		t.Fatalf("esperava 1, obteve %d", len(results))
	}
}

func TestList_InvalidCategory(t *testing.T) {
	svc := newSvc()
	_, err := svc.List("lazer")
	if !errors.Is(err, service.ErrInvalidCategory) {
		t.Fatalf("esperava ErrInvalidCategory, obteve %v", err)
	}
}

func TestUpdate_NotFound(t *testing.T) {
	svc := newSvc()
	desc := "novo"
	_, err := svc.Update(999, domain.UpdateExpenseRequest{Description: &desc})
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("esperava ErrNotFound, obteve %v", err)
	}
}

func TestDelete_NotFound(t *testing.T) {
	svc := newSvc()
	err := svc.Delete(999)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("esperava ErrNotFound, obteve %v", err)
	}
}

func TestUpdate_InvalidAmount(t *testing.T) {
	svc := newSvc()
	expense, _ := svc.Create(domain.CreateExpenseRequest{
		Description: "Almoço", Amount: 10, Category: domain.CategoryFood,
	})
	neg := -5.0
	_, err := svc.Update(expense.ID, domain.UpdateExpenseRequest{Amount: &neg})
	if !errors.Is(err, service.ErrInvalidAmount) {
		t.Fatalf("esperava ErrInvalidAmount, obteve %v", err)
	}
}
