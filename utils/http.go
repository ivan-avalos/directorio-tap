package utils

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	// Response representa un response base
	Response struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}
	// ErrorResponse representa un response de error
	ErrorResponse struct {
		Code    int         `json:"code"`
		Type    string      `json:"type"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

// BaseResponse regresa una response standard
func BaseResponse(code int, data interface{}) Response {
	return Response{
		Code: code,
		Data: data,
	}
}

func (er ErrorResponse) getErrorResponse() *echo.HTTPError {
	return echo.NewHTTPError(er.Code, er)
}

func processValidationError(errs validator.ValidationErrors) *echo.HTTPError {
	// Pueden copiar y pegar esto si quieren
	errsMap := make([]map[string]string, 0)
	for _, err := range errs {
		fieldErr := make(map[string]string)
		fieldErr["field"] = err.Field()
		fieldErr["tag"] = err.Tag()
		if err.Param() != "" {
			fieldErr["param"] = err.Param()
		}
		errsMap = append(errsMap, fieldErr)
	}
	return ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
		Type:    "validation_failed",
		Data:    errsMap,
	}.getErrorResponse()
}

// ProcessError maneja adecuadamente un error para response
func ProcessError(err error) *echo.HTTPError {
	// También pueden copiar y pegar esto
	switch e := err.(type) {
	case validator.ValidationErrors:
		return processValidationError(e)
	case *mysql.MySQLError:
		return ErrorResponse{
			Code:    http.StatusInternalServerError,
			Type:    "database_error",
			Message: "database_error",
			Data:    e.Number,
		}.getErrorResponse()
	}
	switch err {
	case gorm.ErrRecordNotFound:
		return ErrorResponse{
			Code:    http.StatusNotFound,
			Type:    "not_found_error",
			Message: "Record not found",
			Data:    err.Error(),
		}.getErrorResponse()
	}
	return ErrorResponse{
		Code:    http.StatusInternalServerError,
		Type:    "unknown_error",
		Message: "Unknown error",
		Data:    nil,
	}.getErrorResponse()
}
