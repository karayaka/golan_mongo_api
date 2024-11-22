package main

import (
	"fmt"
	"golang_mongo_api/controllers"
	"golang_mongo_api/persistence"
	"golang_mongo_api/repositorys"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	fmt.Print("deene")
	envErr := godotenv.Load()

	e := echo.New()
	if envErr != nil {
		e.Logger.Fatal("env yüklenemdi")
	}

	dbContext := persistence.RegisterWorkerContext(
		os.Getenv("USER_DB_CONNECTION"),
		os.Getenv("DB_NAME"),
		e.Logger)
	db := dbContext.Init()
	workerRepository := repositorys.NewWorkerRepository(db)
	controllers.RegisterControllers(e, workerRepository)

	e.GET("/", func(c echo.Context) error {

		worker, _ := workerRepository.GetAllWorkers()
		return c.JSON(http.StatusOK, worker)
	})

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
	fmt.Println(port)

}
