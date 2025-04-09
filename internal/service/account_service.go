package service

import (
	"github.com/gabrielalmir/go-gateway-api/internal/domain"
	"github.com/gabrielalmir/go-gateway-api/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository}
}

func (s *AccountService) CreateAccount(input dto.CreateAccount) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)
	existingAccount, err := s.repository.FindByApiKey(account.ApiKey)

	if existingAccount != nil {
		return nil, domain.ErrDuplicatedApiKey
	}

	if err != nil {
		return nil, err
	}

	if err := s.repository.Save(&account); err != nil {
		return nil, err
	}

	return dto.FromAccount(&account), nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, domain.ErrAccountNotFound
	}

	account.AddBalance(amount)

	if err := s.repository.UpdateBalance(account); err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}

func (s *AccountService) FindByApiKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, domain.ErrAccountNotFound
	}

	return dto.FromAccount(account), nil
}
