package mongo_client

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
	"time"

	"github.com/vadym-vasilyev/collector-go/src/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	envMongoDbHosts    = "MONGO_HOSTS"
	envMongoDbUser     = "MONGO_USERNAME"
	envMongoDbPassword = "MONGO_PASSWORD"
	recordsDB          = "records"
)

var (
	Client mongoClientInterface = &mongoClient{}
)

type mongoClientInterface interface {
	setClient(c *mongo.Client)
	SaveMany(collectionName string, records []interface{}) error
	Save(collectionName string, record interface{}) error
	Fetch(collectionName string, record interface{}) (*mongo.Cursor, error)
}

type mongoClient struct {
	client   *mongo.Client
	database *mongo.Database
}

func Init() {
	currentHost := fmt.Sprintf("mongodb://%s:%s@%s",
		strings.TrimSpace(os.Getenv(envMongoDbUser)),
		strings.TrimSpace(os.Getenv(envMongoDbPassword)),
		strings.TrimSpace(os.Getenv(envMongoDbHosts)))
	logger.Info("Trying to connect to mongodb: " + currentHost)
	client, err := mongo.NewClient(options.Client().ApplyURI(currentHost))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	client.Database("records").Collection("records")
	Client.setClient(client)
}

func (c *mongoClient) Save(collectionName string, record interface{}) error {
	if _, err := c.database.Collection(collectionName).InsertOne(context.Background(), record); err != nil {
		logger.Error(fmt.Sprintf("Get an error during saving document %#v", record), err)
		return err
	}
	return nil
}

func (c *mongoClient) SaveMany(collectionName string, records []interface{}) error {
	if _, err := c.database.Collection(collectionName).InsertMany(context.Background(), records); err != nil {
		logger.Error(fmt.Sprintf("Get an error during saving documents %#v", records), err)
		return err
	}
	return nil
}

func (c *mongoClient) Fetch(collectionName string, record interface{}) (*mongo.Cursor, error) {
	cursor, err := c.database.Collection(collectionName).Find(context.Background(), bson.D{})
	if err != nil {
		logger.Error(fmt.Sprintf("Get an error during find documents for filter %#v", record), err)
		return nil, err
	} else {
		return cursor, nil
	}
}

func (c *mongoClient) setClient(client *mongo.Client) {
	c.client = client
	c.database = client.Database(recordsDB)
}
