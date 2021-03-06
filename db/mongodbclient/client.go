package mongodbclient

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DBConfigForMongo struct {
	UserName string
	Password string
	Host     string
	Port     string
	AppName  string
}

var dbClient *mongo.Client

func InitMongoDB(config *DBConfigForMongo) (err error) {
	connectionString := "mongodb://" + config.UserName + ":" + config.Password + "@" + config.Host + config.Port + "/?appName=" + config.AppName
	dbClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	return
}

func GetDB(dbName string) (db *mongo.Database) {
	db = dbClient.Database(dbName)
	return
}

func CloseClient() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_ = dbClient.Disconnect(ctx)
}
