package rapi

import (
	"net/http"
)

type Code string

const (
	Success              Code = "success"
	BadRequest           Code = "bad_request"
	UnprocessableEntity  Code = "unprocessable_entity"
	InternalError        Code = "internal_error"
	BindFailed           Code = "bind_failed"
	Unauthorized         Code = "unauthorized"
	Forbidden            Code = "forbidden"
	NotFound             Code = "not_found"
	DataValidationFailed Code = "data_validation_failed"
)

var HTTPStatusByStatusCode = map[Code]int{
	Success:              http.StatusOK,
	BadRequest:           http.StatusBadRequest,
	UnprocessableEntity:  http.StatusUnprocessableEntity,
	InternalError:        http.StatusInternalServerError,
	BindFailed:           http.StatusBadRequest,
	Unauthorized:         http.StatusUnauthorized,
	Forbidden:            http.StatusForbidden,
	NotFound:             http.StatusNotFound,
	DataValidationFailed: http.StatusUnprocessableEntity,
}

var DetailByStatusCode = map[string]string{
	string(InternalError):        "Se ha producido un error interno en el servidor. Por favor, inténtelo de nuevo más tarde o contactese con soporte.",
	string(BadRequest):           "La solicitud no es válida o no puede ser procesada.",
	string(BindFailed):           "Verifique la información de la solicitud y vuelva a intentarlo.",
	string(Unauthorized):         "Asegúrese de proporcionar credenciales válidas y vuelva a intentarlo.",
	string(Forbidden):            "No tiene los permisos necesarios para realizar esta acción.",
	string(NotFound):             "El recurso solicitado no existe o no está disponible en el servidor.",
	string(DataValidationFailed): "La información proporcionada no cumple con los requisitos de validación. Por favor, verifique los datos e inténtelo de nuevo.",
	string(UnprocessableEntity):  "La información proporcionada no cumple con los requisitos de validación. Por favor, verifique los datos e inténtelo de nuevo.",
}

var TitleByStatusCode = map[string]string{
	string(InternalError):        "Error del servidor",
	string(BadRequest):           "Solicitud no válida",
	string(BindFailed):           "Error al procesar los datos",
	string(Unauthorized):         "No identificado",
	string(Forbidden):            "Acceso denegado",
	string(NotFound):             "Recurso no encontrado",
	string(DataValidationFailed): "Validación de datos fallida",
	string(UnprocessableEntity):  "Validación de datos fallida",
}
