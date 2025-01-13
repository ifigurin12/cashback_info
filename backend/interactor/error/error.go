package prjerror

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "record not found"
}

type NotAuthorizedError struct{}

func (e *NotAuthorizedError) Error() string {
	return "not authorized"
}

type InvalidJWTError struct{}

func (e *InvalidJWTError) Error() string {
	return "invalid token format"
}
