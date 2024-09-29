package database

import (
	"context"
	"fmt"
	"go-graphql-mongodb-project/graph/model"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb://localhost:27017"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetBoards() []*model.Board {
	boardCollec := db.client.Database("reeCraft").Collection("boards")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var boardListings []*model.Board
	cursor, err := boardCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &boardListings); err != nil {
		panic(err)
	}

	return boardListings
}

func (db *DB) CreateBoard(ctx context.Context, input model.CreateBoardInput) (*model.Board, error) {
	newBoard := &model.Board{
		ID:       primitive.NewObjectID().Hex(),
		Name:     input.Name,
		IsActive: input.IsActive,
		Columns:  make([]*model.Column, len(input.Columns)),
	}

	for i, col := range input.Columns {
		newColumn := &model.Column{
			Name:  col.Name,
			Tasks: make([]*model.Task, len(col.Tasks)),
		}

		for j, task := range col.Tasks {
			newTask := &model.Task{
				Title:       task.Title,
				Description: task.Description,
				Status:      task.Status,
				Subtasks:    make([]*model.Subtask, len(task.Subtasks)),
			}

			for k, subtask := range task.Subtasks {
				newTask.Subtasks[k] = &model.Subtask{
					Title:       subtask.Title,
					IsCompleted: subtask.IsCompleted,
				}
			}
			newColumn.Tasks[j] = newTask
		}
		newBoard.Columns[i] = newColumn
	}

	boardCollection := db.client.Database("your_database_name").Collection("boards") 
	_, err := boardCollection.InsertOne(ctx, newBoard)
	if err != nil {
		return nil, fmt.Errorf("failed to insert board: %v", err)
	}

	return newBoard, nil
}

func (db *DB) GetBoardByID(ctx context.Context, id string) (*model.Board, error) {
	boardCollection := db.client.Database("your-database-name").Collection("boards")
	objectID, err := primitive.ObjectIDFromHex(id)
	fmt.Println("Filter being used:", bson.M{"_id": objectID}, "Type of objectID:", reflect.TypeOf(id)) 

	if err != nil {
		return nil, fmt.Errorf("invalid board ID: %v", err)
	}

	var board model.Board
	err = boardCollection.FindOne(ctx, bson.M{"id": id}).Decode(&board)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("board not found")
		}
		return nil, fmt.Errorf("failed to fetch board: %v", err)
	}

	return &board, nil
}
