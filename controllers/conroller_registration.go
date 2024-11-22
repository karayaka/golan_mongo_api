package controllers

import (
	"golang_mongo_api/models/view_models/response"
	"golang_mongo_api/repositorys"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func RegisterControllers(e *echo.Echo, wr repositorys.IWorkerRepository) {

	wc := NewWorkerConroller(wr)
	wc.RegisterWorkerController(e)
}

func Ok[T any](c echo.Context, result T) error {
	res := response.BaseResponse[T]{
		Data:    result,
		Message: "İşlem Başarılı",
		Date:    time.Now(),
	}
	c.JSON(http.StatusOK, res)
	return nil
}
func ErrorResponse(c echo.Context, err error) error {
	res := response.ErrorResponse{
		Message: err.Error(),
		Date:    time.Now(),
	}
	c.JSON(http.StatusBadRequest, res)
	return nil
}
