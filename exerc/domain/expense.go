package domain

import "time"

type ExpenseCategory string

const (
	CategoryFood      ExpenseCategory = "alimentacao"
	CategoryTransport ExpenseCategory = "transporte"
	CategoryHealth    ExpenseCategory = "saude"
	CategoryEducation ExpenseCategory = "educacao"
	CategoryOther     ExpenseCategory = "outro"
)

var validCategories = map[ExpenseCategory]bool{
	CategoryFood:      true,
	CategoryTransport: true,
	CategoryHealth:    true,
	CategoryEducation: true,
	CategoryOther:     true,
}

func (c ExpenseCategory) IsValid() bool {
	return validCategories[c]
}

type Expense struct {
	ID          uint            `json:"id" gorm:"primaryKey"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	Category    ExpenseCategory `json:"category"`
	Date        time.Time       `json:"date"`
	CreatedAt   time.Time       `json:"created_at"`
}

type CreateExpenseRequest struct {
	Description string          `json:"description" binding:"required,min=3"`
	Amount      float64         `json:"amount"      binding:"required,gt=0"`
	Category    ExpenseCategory `json:"category"    binding:"required"`
}

type UpdateExpenseRequest struct {
	Description *string          `json:"description"`
	Amount      *float64         `json:"amount"`
	Category    *ExpenseCategory `json:"category"`
}
