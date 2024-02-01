package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	
	"github.com/lutfiahdewi/graphql-go/graph"
	database "github.com/lutfiahdewi/graphql-go/internal/pkg/db/mssql"


	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	envErr := godotenv.Load();
	if envErr != nil {
       fmt.Printf("Error loading credentials: %v \n", envErr)
    }
	var (
		password = os.Getenv("MSSQL_DB_PASSWORD")
		user = os.Getenv("MSSQL_DB_USER")
		// dbPort = os.Getenv("MSSQL_DB_PORT")
		dbName = os.Getenv("MSSQL_DB_DATABASE")
	 )

	/*srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)*/

	router := chi.NewRouter()

	// router.Use(auth.Middleware())

	connectionString := fmt.Sprintf("sqlserver://%s:%s@localhost?database=%s", user, password, dbName)
	database.OpenDB(connectionString, false)
	defer database.CloseDB()
	// database.Migrate()
	// server := handler.NewDefaultServer(hackernews.NewExecutableSchema(hackernews.Config{Resolvers: &hackernews.Resolver{}}))
	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
