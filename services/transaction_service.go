package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func (s *TransactionService) Checkout(items []models.CheckoutItem, useLock bool) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}
