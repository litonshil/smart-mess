package transaction

import (
	"context"
	"gorm.io/gorm"
	"smart-mess/domain"
	"smart-mess/infra/logger"
)

type DBTransactionService struct {
	loggeruc logger.LogClient
	repo     domain.TXRepo
}

func NewTXClient(ctx context.Context, client *gorm.DB) *domain.TXClient {
	return &domain.TXClient{
		Ctx:    ctx,
		Client: client,
	}
}

func NewDBTransaction(loggeruc logger.LogClient, repo domain.TXRepo) *DBTransactionService {
	return &DBTransactionService{
		repo:     repo,
		loggeruc: loggeruc,
	}
}

func (svc *DBTransactionService) CreateTransaction(ctx context.Context) (*domain.TXClient, error) {
	txc, err := svc.repo.CreateTransaction(ctx)
	if err != nil {
		svc.loggeruc.Error("transaction creation failed", err)
		return nil, err
	}
	return txc, nil
}
