package repository

import (
	"github.com/gocql/gocql"
	"log"
	"websocket-example/model"
)

func Store(t *model.Timeline) (err error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		t.TimeLine, gocql.TimeUUID(), t.Text).Exec(); err != nil {
		log.Fatal(err)
	}
	return
}
