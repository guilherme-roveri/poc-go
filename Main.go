package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gbroveri/users/gateways/cassandra"
	"github.com/gbroveri/users/usecases"

	"github.com/gbroveri/users/api"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "dev"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}

func main() {
	defer Session.Close()
	r := mux.NewRouter()
	handler := api.NewUserAPIHandler(usecases.NewUserCrud(cassandra.NewUserGatewayCassandra(Session)))
	r.HandleFunc("/users", handler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
