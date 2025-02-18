package interfaces

import (
	"context"
)

type DbInterface interface {
	CloseDB() error

	rpsRegistrarInterface
}

type rpsRegistrarInterface interface {
	ServiceAvailable(ctx context.Context, inst string) bool
}

type ValidationDbInterface interface {
	CloseDB() error
}
