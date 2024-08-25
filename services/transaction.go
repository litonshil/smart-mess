package services

import (
	"fmt"
	"smart-mess/domain"
	"smart-mess/infra/logger"
	"smart-mess/utils/consts"
)

func TransactionRollback(
	txc *domain.TXClient,
	loggeruc logger.LogClient,
	entity consts.Entity,
	action consts.Action,
) error {
	if err := txc.Rollback(); err != nil {
		loggeruc.Error(
			fmt.Sprintf(
				"error occurred while transaction rollbacked for %v %v",
				entity,
				action,
			),
			err,
		)
		return err
	}

	loggeruc.Info("transaction rollbacked successfully ...")
	return nil
}

func TransactionCommit(
	txc *domain.TXClient,
	loggeruc domain.LoggerUseCase,
	entity consts.Entity,
	action consts.Action,
) error {
	if err := txc.Commit(); err != nil {
		loggeruc.Error(
			fmt.Sprintf(
				"error occurred while %v %v transaction commit",
				entity,
				action,
			),
			err,
		)
		return err
	}

	loggeruc.Info("transaction successfully committed...")
	return nil
}
