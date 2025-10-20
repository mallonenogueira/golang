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
		c.ResponseBadRequest("JSON inválido")
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

func (c *Context) ResponseBadRequest(message string) {
	c.ResponseJson(http.StatusBadRequest, map[string]string{"message": message})
}

func (c *Context) ResponseNotFound(message string) {
	c.ResponseJson(http.StatusNotFound, map[string]string{"message": message})
}

func (c *Context) ResponseInternalError(message string) {
	c.ResponseJson(http.StatusInternalServerError, map[string]string{"message": message})
}

func (c *Context) ResponseError(data errors.AppError) {
	c.ResponseJson(data.StatusCode(), data)
}

func (c *Context) ResponseNoContent() {
	c.writer.WriteHeader(http.StatusNoContent)
}
