package errors

type AppError interface {
	StatusCode() int
	Type() string
	GetMessage() string
}

type AppErrorField interface {
	GetField() string
}

type AppErrorDetails interface {
	GetDetails() map[string]interface{}
}

type ValidationError struct {
	Message string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}

func (e *ValidationError) StatusCode() int    { return 400 }
func (e *ValidationError) Type() string       { return "VALIDATION_ERRROR" }
func (e *ValidationError) GetMessage() string { return e.Message }

type FieldValidationError struct {
	Field   string
	Message string
}

func NewFieldValidationError(message string, field string) *FieldValidationError {
	return &FieldValidationError{Message: message, Field: field}
}

func (e *FieldValidationError) StatusCode() int    { return 400 }
func (e *FieldValidationError) GetMessage() string { return e.Message }
func (e *FieldValidationError) Type() string       { return "FIELD_VALIDATION_ERRROR" }
func (e *FieldValidationError) GetField() string   { return e.Field }

type NotFoundError struct {
	Message  string
	Resource string
}

func NewNotFoundError(message string, resource string) *NotFoundError {
	return &NotFoundError{Message: message, Resource: resource}
}

func (e *NotFoundError) StatusCode() int    { return 404 }
func (e *NotFoundError) GetMessage() string { return e.Message }
func (e *NotFoundError) Type() string       { return "NOT_FOUND" }

func (e *NotFoundError) GetDetails() map[string]interface{} {
	return map[string]interface{}{"resource": e.Resource}
}
