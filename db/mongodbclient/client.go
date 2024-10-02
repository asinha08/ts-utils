package mongodbclient

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfigForMongo struct {
	UserName     string
	Password     string
	Host         string
	Port         string
	AppName      string
	DataBaseName string
}

var dbClient *mongo.Client
var dbName string

func InitMongoDB(config *DBConfigForMongo) (err error) {
	dbName = config.DataBaseName
	connectionString := "mongodb://" + config.UserName + ":" + config.Password + "@" + config.Host + config.Port + "/?appName=" + config.AppName
	dbClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	return
}

func InitMongoDbClinet(config *DBConfigForMongo) (err error) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(
		"mongodb+srv://" +
			config.UserName +
			":" +
			config.Password +
			"@" +
			config.Host +
			"/?retryWrites=true&w=majority&appName=" +
			config.AppName).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	dbClient, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	return
}

func GetDB() (db *mongo.Database) {
	db = dbClient.Database(dbName)
	return
}

func CloseClient() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_ = dbClient.Disconnect(ctx)
}
