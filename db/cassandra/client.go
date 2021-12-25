package cassandra

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"go.elastic.co/apm/module/apmgocql"
)

const maxRetry = 60

var session *gocql.Session = nil

// InitDBSession TODO: write comment
func InitDBSession(host string, keySpace string, serviceName string) *gocql.Session {
	observer := apmgocql.NewObserver()
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keySpace
	cluster.QueryObserver = observer
	cluster.BatchObserver = observer
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

// GetDBSession TODO: write comment
func GetDBSession() *gocql.Session {
	return session
}

// CloseDBSession TODO: write comment
func CloseDBSession() {
	session.Close()
}
