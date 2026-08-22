package services

import (
	"github.com/chishkin-afk/interval-learning-algo/backend/internal/application/dtos/responses"
)

type kind int

const (
	KindRequest kind = 0
	KindTimeout
	KindUnauthorized
	KindPermissionDenied
	KindConflict
	KindNotFound
	KindInternal
)

type ServiceError struct {
	kind kind
	err  responses.Err
}

func (se *ServiceError) Error() string {
	return se.err.Error
}

func (se *ServiceError) Kind() kind {
	return se.kind
}

func NewServiceError(k kind, e error) *ServiceError {
	return &ServiceError{
		kind: k,
		err:  responses.Err{Error: e.Error()},
	}
}
