package main

import (
	"bookstore-management-app/pkg/controllers"
	"bookstore-management-app/pkg/routes"
	"bookstore-management-app/pkg/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "bookstore-management-app/pkg/utils"
// "log"
// "net/http"
// "github.com/gorilla/mux"

func main() {
	db, err := utils.InitDB()
	if err != nil {
		log.Fatal("DB connection error")
	}
	err = utils.AutoMigrate(db)
	if err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	controllers.InitDBInstance(db)
	fmt.Println("hi")

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)

	}

}
