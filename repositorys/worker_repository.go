package repositorys

import (
	"context"
	"fmt"
	entitymodels "golang_mongo_api/models/entity_models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IWorkerRepository interface {
	AddWorker(model *entitymodels.Worker) (string, error)
	DeleteWorker(id string) error
	GetById(id string) (*entitymodels.Worker, error)
	GetAllWorkers() ([]entitymodels.Worker, error)
	UpdateWorkers(model *entitymodels.Worker) error
}

type WorkerRepository struct {
	collection *mongo.Collection
}

func NewWorkerRepository(db *mongo.Database) IWorkerRepository {
	return &WorkerRepository{
		collection: db.Collection("workers"),
	}
}

// UpdateWorkers implements IWorkerRepository.
func (wr *WorkerRepository) UpdateWorkers(model *entitymodels.Worker) error {
	oid, _ := primitive.ObjectIDFromHex(model.Id)

	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": model.ToMap(), //bu konu düşünülmeli!!!
	}
	rs, err := wr.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println(rs)

	return nil
}

func (wr WorkerRepository) AddWorker(model *entitymodels.Worker) (string, error) {
	result, err := wr.collection.InsertOne(context.TODO(), model)
	if err != nil {
		return "", err
	}
	return fmt.Sprint("", result.InsertedID), nil
}

func (wr WorkerRepository) DeleteWorker(id string) error {
	filter := bson.M{"_id": id}
	_, err := wr.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil
	}
	return nil
}

// GetAllWorkers implements IWorkerRepository.
func (wr *WorkerRepository) GetAllWorkers() ([]entitymodels.Worker, error) {
	filter := bson.D{}
	cursor, err := wr.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var workers []entitymodels.Worker

	for cursor.Next(context.TODO()) {

		var worker entitymodels.Worker

		if err := cursor.Decode(&worker); err != nil {
			return nil, err
		}
		workers = append(workers, worker)
		if err := cursor.Err(); err != nil {
			return nil, err
		}
	}
	return workers, nil
}

// GetById implements IWorkerRepository. yazılacak
func (wr *WorkerRepository) GetById(id string) (*entitymodels.Worker, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": oid}
	var worker entitymodels.Worker
	err := wr.collection.FindOne(context.TODO(), filter).Decode(&worker)
	if err != nil {
		return nil, err
	}
	return &worker, err
}
