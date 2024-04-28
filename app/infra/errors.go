package infra

import (
	"errors"

	"github.com/ganganbiz1/go-echo-gorm-template/app/domain"
	"github.com/ganganbiz1/go-echo-gorm-template/app/logger"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func HandleError(originErr error) error {
	if errors.Is(originErr, gorm.ErrRecordNotFound) {
		return domain.ErrNotFound
	}

	logger.Error(originErr.Error(), originErr)

	if mysqlErr, ok := originErr.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062:
			// Error Code: 1062. Duplicate entry ‘xxx’ for key ‘xxx.PRIMARY’
			return domain.ErrAlreadyExist
		case 1213:
			// ERROR 1213 (40001): Deadlock found when trying to get lock;
			return domain.ErrDeadlock
		}
	}
	return originErr
}
