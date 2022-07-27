// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// 	"github.com/go-chi/cors"
// 	"github.com/go-chi/httplog"

// 	"github.com/sam-app/hackernews/graph"
// 	"github.com/sam-app/hackernews/graph/generated"
// 	database "github.com/sam-app/hackernews/packages/db/postgress"
// )

// const defaultPort = "8080"

// func commonMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Add("Content-Type", "application/json")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	// init database
// 	dbError := database.InitDB()
// 	if dbError != nil {
// 		log.Fatalf("Error connecting to database: %s", dbError)
// 	}

// 	r := chi.NewRouter()
// 	r.Use(cors.Handler(opts))

// 	// A good base middleware stack
// 	r.Use(middleware.RequestID)
// 	r.Use(middleware.RealIP)
// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)
// 	// Logger
// 	logger := httplog.NewLogger("httplog-example", httplog.Options{
// 		JSON: true,
// 	})

// 	r.Use(httplog.RequestLogger(logger))

// 	// Set a timeout value on the request context (ctx), that will signal
// 	// through ctx.Done() that the request has timed out and further
// 	// processing should be stopped.
// 	srv := handler.NewDefaultServer(generated.NewExecutableSchema((generated.Config{Resolvers: &graph.Resolver{}})))
// 	r.Use(middleware.Timeout(60 * time.Second))
// 	r.Handle("/graph", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, r))
// }

// var opts = cors.Options{
// 	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
// 	AllowedOrigins: []string{"https://*", "http://*"},
// 	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 	AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
// 	//ExposedHeaders:   []string{"Link"},
// 	AllowCredentials: false,
// 	MaxAge:           300, // Maximum value not ignored by any of major browsers
// }

// // r := chi.NewRouter()
// // r.Use(cors.New(cors.Options{
// // 	AllowedOrigins: []string{"*"},
// // 	MaxAge:         86000,
// // }).Handler)
// // r.Use(commonMiddleware)

// // server := handler.NewDefaultServer(generated.NewExecutableSchema((generated.Config{Resolvers: &graph.Resolver{}})))

// // r.Handle("/", playground.Handler("GraphQL playground", "query"))
// // r.Handle("/query", server)

// // log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// // log.Fatal(http.ListenAndServe(":"+port, r))

// //}