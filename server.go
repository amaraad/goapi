package main

import (
	"goapi/config"
	"goapi/directives"
	"goapi/graph"
	"goapi/graph/generated"
	"goapi/middlewares"
	migration "goapi/migrations"
	"goapi/rest"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	// Add Migration
	migration.MigrateTable()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// Defer Closing Database
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Authentication router
	router := mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/graphql/", playground.Handler("GraphQL playground", "/graphql/query"))
	router.Handle("/graphql/query", srv)
	router.HandleFunc("/rest/request/get", rest.GetRequest).Methods("GET")
	router.HandleFunc("/rest/request/post", rest.PostRequest).Methods("POST")

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
