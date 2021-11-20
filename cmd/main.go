package main

import (
	"flag"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iDevoid/cptx"
	"github.com/iDevoid/recipes/internal/glue/routing"
	"github.com/iDevoid/recipes/internal/handler/rest"
	"github.com/iDevoid/recipes/internal/module/recipes"
	"github.com/iDevoid/recipes/internal/storage/persistence"
	"github.com/iDevoid/recipes/platform/routers"
	"github.com/sirupsen/logrus"
)

var testInit bool

func init() {
	flag.BoolVar(&testInit, "test", false, "initialize test mode without serving")
	flag.Parse()
}

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	psql := cptx.Initialize(dbURL, dbURL, "recipes")
	postgresDB, postgresTX := psql.Open("mysql")

	dbRecipes := persistence.RecipesInit(postgresDB)
	usecase := recipes.Initialize(postgresTX, dbRecipes)

	handler := rest.RecipesInit(usecase)
	router := routing.RecipesRouting(handler)

	server := routers.Initialize(":8080", router, "recipes")
	if testInit {
		logrus.Info("Initialize test mode Finished!")
		os.Exit(0)
	}

	server.Serve()
}
