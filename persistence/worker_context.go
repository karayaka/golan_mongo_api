package persistence

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WorkerContext struct {
	ConnectionString string
	log              *echo.Logger
	DbName           string
}

func RegisterWorkerContext(cs string, dbName string, log echo.Logger) *WorkerContext {
	return &WorkerContext{
		ConnectionString: cs,
		log:              &log,
		DbName:           dbName,
	}
}
func (uc WorkerContext) Init() *mongo.Database {
	clientOptions := options.Client().ApplyURI(uc.ConnectionString)

	// MongoDB Client başlatma
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Bağlantıyı doğrulama
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("MongoDB bağlantısı başarısız:", err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Asenkron olarak uygulama kapanışını dinleyin
	go func() {
		<-quit
		client.Disconnect(context.TODO())
		os.Exit(0)
	}()

	log.Println("MongoDB bağlantısı başarılı")
	return client.Database(uc.DbName)
}
