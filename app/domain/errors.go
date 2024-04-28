package domain

import (
	"errors"

	"github.com/ganganbiz1/go-echo-gorm-template/app/logger"
)

var ErrMap = map[error]struct{}{}

var ErrInternal = errors.New("ErrInternal")
var ErrNotFound = errors.New("ErrNotFound")
var ErrAlreadyExist = errors.New("ErrAlreadyExist")
var ErrDB = errors.New("ErrDB")
var ErrDeadlock = errors.New("ErrDeadlock")
var ErrTransaction = errors.New("ErrTransaction")
var ErrParmission = errors.New("ErrParmission")
var ErrBadRequest = errors.New("ErrBadRequest")
var ErrTypeConvert = errors.New("ErrTypeConvert")

func init() {
	ErrMap[ErrInternal] = struct{}{}
	ErrMap[ErrNotFound] = struct{}{}
	ErrMap[ErrAlreadyExist] = struct{}{}
	ErrMap[ErrDB] = struct{}{}
	ErrMap[ErrDeadlock] = struct{}{}
	ErrMap[ErrTransaction] = struct{}{}
	ErrMap[ErrParmission] = struct{}{}
	ErrMap[ErrBadRequest] = struct{}{}
	ErrMap[ErrTypeConvert] = struct{}{}
}

func HandleError(domainErr error, originErr error) error {
	logger.Error("", originErr)
	return domainErr
}
