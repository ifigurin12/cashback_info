package utility

import (
	"errors"
	"fmt"
	"net/http"

	prjerror "cashback_info/interactor/error"
)

func TransformErrorToHttpError(err error) (int, string) {
	switch {
	case errors.Is(err, &prjerror.NotFoundError{}):
		return http.StatusNotFound, "Not found"
	case errors.Is(err, &prjerror.NotAuthorizedError{}):
		return http.StatusUnauthorized, "Not authorized"
	case errors.Is(err, &prjerror.InvalidJWTError{}):
		return http.StatusBadRequest, "Invalid JWT token format"
	default:
		fmt.Println(err)
		return http.StatusInternalServerError, "Internal server error"
	}
}
