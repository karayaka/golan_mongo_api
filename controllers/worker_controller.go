package controllers

import (
	"golang_mongo_api/middlewares"
	entitymodels "golang_mongo_api/models/entity_models"
	"golang_mongo_api/models/view_models/request"
	"golang_mongo_api/repositorys"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type WorkerController struct {
	wr repositorys.IWorkerRepository
}

func NewWorkerConroller(wr repositorys.IWorkerRepository) *WorkerController {
	return &WorkerController{
		wr: wr,
	}
}

func (wc WorkerController) RegisterWorkerController(e *echo.Echo) {
	g := e.Group("api/worker", middlewares.AuthMiddleware) //groplayıp koruma sağlanıyor
	g.GET("/GetAllWorkers", wc.GetWorkers)
	g.POST("/AddWorker", wc.AddWorker)
	g.PUT("/UpdateWorker", wc.UpdateWorker)
	g.GET("/GetById/:id", wc.GetWorkerById)
}

func (wc WorkerController) GetWorkers(c echo.Context) error {
	workers, err := wc.wr.GetAllWorkers()
	if err != nil {
		return ErrorResponse(c, err)
	}
	return Ok(c, workers)
}

func (wc WorkerController) AddWorker(c echo.Context) error {
	var worker request.WorkerRequestModel
	err := c.Bind(&worker)
	if err != nil {
		return ErrorResponse(c, err)
	}
	sessionId, _ := strconv.ParseUint(c.Get("sesionId").(string), 10, 32)
	wc.wr.AddWorker(&entitymodels.Worker{
		UserId:      worker.UserId,
		Name:        worker.Name,
		Surname:     worker.Surname,
		Department:  worker.Department,
		Email:       worker.Email,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		CreatedBy:   sessionId,
		UpdatedBy:   sessionId,
	})
	return Ok(c, "")
}
func (wc WorkerController) UpdateWorker(c echo.Context) error {
	var worker request.WorkerRequestModel
	err := c.Bind(&worker)
	if err != nil {
		return ErrorResponse(c, err)
	}
	sessionId, _ := strconv.ParseUint(c.Get("sesionId").(string), 10, 32)
	wc.wr.UpdateWorkers(&entitymodels.Worker{
		Id:          worker.Id,
		UserId:      worker.UserId,
		Name:        worker.Name,
		Surname:     worker.Surname,
		Department:  worker.Department,
		Email:       worker.Email,
		UpdatedDate: time.Now(),
		UpdatedBy:   sessionId,
	})
	return Ok(c, "")
}
func (wc WorkerController) GetWorkerById(c echo.Context) error {
	id := c.Param("id")
	worker, err := wc.wr.GetById(id)
	if err != nil {
		return ErrorResponse(c, err)
	}
	return Ok(c, worker)
}
