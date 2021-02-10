package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/jakewitcher/pos-server/graph"
	"github.com/jakewitcher/pos-server/graph/generated"
	"github.com/jakewitcher/pos-server/internal/datastore"
	"github.com/jakewitcher/pos-server/internal/datastore/sqlite"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func init() {
	db := sqlite.OpenConnection()
	datastore.Customers = sqlite.NewCustomerProvider(db)
	datastore.Stores = sqlite.NewStoreProvider(db)
	datastore.StoreLocations = sqlite.NewStoreLocationProvider(db)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
