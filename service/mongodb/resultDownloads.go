package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

type ResultDownloadsServise interface {
	InitConnections() error
	Disconnections()
	InsertOneIntoCollect() error
	NewData(data map[string]string)
}

type servise struct {
	ctx        context.Context
	client     *mongo.Client
	collection *mongo.Collection
	inputData  map[string]string
}

func NewService() *servise {
	return &servise{}
}

func (s *servise) NewData(data map[string]string) {
	s.inputData = data
}

func (s *servise) InitConnections() error {
	mongoURL := os.Getenv("CONFIG_MONGODB_URL")

	var err error

	//s.client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kateryna:katya135@mybackendcluster.dvi6ngn.mongodb.net/?retryWrites=true&w=majority"))
	s.client, err = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		return err
	}

	s.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	err = s.client.Connect(s.ctx)
	if err != nil {
		return err
	}

	err = s.client.Ping(s.ctx, readpref.Primary())
	if err != nil {
		return err
	}

	db := s.client.Database("my-backend-mongodb")
	s.collection = db.Collection("resultDownloads")

	return nil
}

func (s *servise) Disconnections() {
	s.client.Disconnect(s.ctx)
}

func (s *servise) InsertOneIntoCollect() error {

	incData := bson.M{}
	for k, v := range s.inputData {
		incData[k] = v
	}

	_, err := s.collection.InsertOne(s.ctx, bson.M{"$inc": incData})

	if err != nil {
		return err
	}

	return nil
}
