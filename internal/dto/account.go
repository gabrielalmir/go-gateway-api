package dto

import "github.com/gabrielalmir/go-gateway-api/internal/domain"

type CreateAccount struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountOutput struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Balance   float64 `json:"balance"`
	ApiKey    string  `json:"api_key,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func ToAccount(input CreateAccount) domain.Account {
	return *domain.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) *AccountOutput {
	return &AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		ApiKey:    account.ApiKey,
		CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: account.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
