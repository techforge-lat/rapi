package rapi

import (
	"net/http"
)

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
	return &Response{Type: "https://datatracker.ietf.org/doc/html/rfc9457"}
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
		Data:       data,
		StatusCode: string(Success),
		Status:     http.StatusOK,
	}
}

func UnexpectedError() *Response {
	return &Response{
		Detail:     DetailByStatusCode[string(InternalError)],
		StatusCode: string(InternalError),
		Status:     http.StatusInternalServerError,
	}
}

func Created(data any) *Response {
	return &Response{
		Title:      "Petición exitosa",
		Detail:     "Recurso creado",
		StatusCode: "created",
	}
}

func NotFoundResponse() *Response {
	return &Response{
		Title:      "Recurso no encontrado",
		Detail:     "El recurso solicitado no existe",
		StatusCode: "not_found",
	}
}

func Updated() *Response {
	return &Response{
		Title:      "Petición exitosa",
		Detail:     "Recurso actualizado",
		StatusCode: "updated",
	}
}
