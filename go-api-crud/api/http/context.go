package http

import (
	"encoding/json"
	"go-api-crud/errors"
	"net/http"
)

type Context struct {
	writer  http.ResponseWriter
	request *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		writer:  w,
		request: r,
	}
}

func (c *Context) PathValue(key string) string {
	return c.request.PathValue(key)
}

func (c *Context) BodyJson(v any) bool {
	if err := json.NewDecoder(c.request.Body).Decode(v); err != nil {
		c.HandleError(errors.NewValidationError("JSON inválido."))
		return false
	}

	return true
}

func (c *Context) QueryParam(key string) string {
	return c.request.URL.Query().Get(key)
}

func (c *Context) ResponseJson(statusCode int, data any) {
	c.writer.Header().Set("Content-Type", "application/json")
	c.writer.WriteHeader(statusCode)
	json.NewEncoder(c.writer).Encode(data)
}

func (c *Context) ResponseOk(data any) {
	c.ResponseJson(http.StatusOK, data)
}

func (c *Context) ResponseCreated(data any) {
	c.ResponseJson(http.StatusCreated, data)
}

// func (c *Context) ResponseBadRequest(message string) {
// 	c.ResponseJson(http.StatusBadRequest, map[string]string{"message": message})
// }

// func (c *Context) ResponseNotFound(message string) {
// 	c.ResponseJson(http.StatusNotFound, map[string]string{"message": message})
// }

// func (c *Context) ResponseInternalError(message string) {
// 	c.ResponseJson(http.StatusInternalServerError, map[string]string{"message": message})
// }

// func (c *Context) ResponseError(data errors.AppError) {
// 	c.ResponseJson(data.StatusCode(), data)
// }

func (c *Context) ResponseNoContent() {
	c.writer.WriteHeader(http.StatusNoContent)
}

type errorResponse struct {
	Error errorDetails `json:"error"`
}

type errorDetails struct {
	Status  int                    `json:"status"`
	Type    string                 `json:"type"`
	Message string                 `json:"message"`
	Field   string                 `json:"field,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func (c *Context) HandleError(err any) {
	appErr, ok := err.(errors.AppError)

	if !ok {
		c.ResponseJson(500, errorResponse{
			Error: errorDetails{
				Status:  500,
				Type:    "INTERNAL_ERROR",
				Message: "Ops.",
			},
		})
		return
	}

	details := errorDetails{
		Status:  appErr.StatusCode(),
		Message: appErr.GetMessage(),
		Type:    appErr.Type(),
	}

	if errWithDetails, ok := err.(errors.AppErrorDetails); ok {
		details.Details = errWithDetails.GetDetails()
	}

	if errWithField, ok := err.(errors.AppErrorField); ok {
		details.Field = errWithField.GetField()
	}

	c.ResponseJson(details.Status, details)
}
