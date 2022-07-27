package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/sam-app/hackernews/graph"
	"github.com/sam-app/hackernews/graph/generated"
	database "github.com/sam-app/hackernews/packages/db/postgress"
)

const defaultPort = "8080"

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

var opts = cors.Options{
	//AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	AllowedOrigins:   []string{"https://*", "http://*"},
	AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: false,
	MaxAge:           300, // Maximum value not ignored by any of major browsers
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// init database
	dbError := database.InitDB()
	if dbError != nil {
		log.Fatalf("Error connecting to database: %s", dbError)
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(opts))

	r.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		MaxAge:         86000,
	}).Handler)
	r.Use(commonMiddleware)

	server := handler.NewDefaultServer(generated.NewExecutableSchema((generated.Config{Resolvers: &graph.Resolver{}})))

	//r.Handle("/graph", playground.Handler("GraphQL playground", "query"))
	r.Handle("/graphql", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
