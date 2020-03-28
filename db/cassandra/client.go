package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

const maxRetry = 60

var session *gocql.Session = nil

func InitDBSession(host string, keySpace string, serviceName string) *gocql.Session {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keySpace
	var err error
	var counter int
	for {
		session, err = cluster.CreateSession()
		if err != nil && counter < maxRetry {
			counter++
			time.Sleep(time.Millisecond * 1000)
			fmt.Println(serviceName, ", Unable to connect database. Retrying :: ", counter)
			fmt.Println(err)
		} else if err == nil {
			fmt.Println(serviceName, " is now connected to database")
			break
		}
		if counter == maxRetry {
			panic(err)
		}
	}
	return session
}

func GetDBSession() *gocql.Session {
	return session
}

func CloseDBSession() {
	session.Close()
}
