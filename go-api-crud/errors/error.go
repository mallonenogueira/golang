package errors

type AppError interface {
	StatusCode() int
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) StatusCode() int {
	return 400
}

type FieldValidationError struct {
	Field   string
	Message string
}

func (e *FieldValidationError) StatusCode() int {
	return 400
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) StatusCode() int {
	return 404
}
