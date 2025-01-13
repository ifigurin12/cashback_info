package utility

import (
	prjerror "cashback_info/interactor/error"

	"github.com/jackc/pgx/v5"
)

func TransformError(err error) error {
	if err == pgx.ErrNoRows {
		return &prjerror.NotFoundError{}
	} else {
		return err
	}
}
