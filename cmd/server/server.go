package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/felipefbs/graphql/graph"
	"github.com/felipefbs/graphql/internal/databases"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	categoryRepo := databases.NewCategory(db)
	courseRepo := databases.NewCourse(db)

	_, err = db.Exec("CREATE TABLE categories (id string, name string, description string)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE courses  (id string, name string, description string, category_id) ")
	if err != nil {
		log.Fatal(err)
	}

	createdCategory, err := categoryRepo.Create("Tecnologia", "Cursos de Tecnologia")
	if err != nil {
		log.Fatal(err)
	}
	_, err = courseRepo.Create("Full Cycle", "curso full cycle", createdCategory.ID)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			CategoryRepository: *categoryRepo,
			CourseRepository:   *courseRepo,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
