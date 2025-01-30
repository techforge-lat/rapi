package rapi

import (
	"net/http"
)

const responseType = "https://datatracker.ietf.org/doc/html/rfc9457"

type Response struct {
	Data       any    `json:"data,omitempty"`
	Type       string `json:"type"`
	Title      string `json:"title,omitempty"`
	Detail     string `json:"detail,omitempty"`
	DebugError string `json:"debugError,omitempty"`
	Status     uint   `json:"status"`
	StatusCode string `json:"statusCode,omitempty"`
	Instance   string `json:"instance,omitempty"`
}

func New() *Response {
	return &Response{Type: responseType}
}

func (r *Response) SetType(t string) *Response {
	r.Type = t

	return r
}

func (r *Response) SetTitle(title string) *Response {
	r.Title = title

	return r
}

func (r *Response) SetDetail(detail string) *Response {
	r.Detail = detail

	return r
}

func (r *Response) SetDebugError(debugErr string) *Response {
	r.DebugError = debugErr

	return r
}

func (r *Response) SetStatus(status uint) *Response {
	r.Status = status

	return r
}

func (r *Response) SetStatusCode(status string) *Response {
	r.StatusCode = status

	return r
}

func (r *Response) SetInstance(instance string) *Response {
	r.Instance = instance

	return r
}

func Ok(data any) *Response {
	return &Response{
		Type:       responseType,
		Data:       data,
		StatusCode: string(Success),
		Status:     http.StatusOK,
	}
}

func UnexpectedError() *Response {
	return &Response{
		Type:       responseType,
		Detail:     DetailByStatusCode[string(InternalError)],
		StatusCode: string(InternalError),
		Status:     http.StatusInternalServerError,
	}
}

func Created(data any) *Response {
	return &Response{
		Type:       responseType,
		Title:      "Petición exitosa",
		Detail:     "Recurso creado",
		StatusCode: "created",
		Status:     http.StatusCreated,
	}
}

func NotFoundResponse() *Response {
	return &Response{
		Type:       responseType,
		Title:      "Recurso no encontrado",
		Detail:     "El recurso solicitado no existe",
		StatusCode: "not_found",
		Status:     http.StatusNotFound,
	}
}

func Updated() *Response {
	return &Response{
		Type:       responseType,
		Title:      "Petición exitosa",
		Detail:     "Recurso actualizado",
		StatusCode: "updated",
		Status:     http.StatusOK,
	}
}
