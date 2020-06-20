package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"godoksip/docks/repo"
	"godoksip/docks/usecase"
	"godoksip/middleware"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	docksHandler "godoksip/docks/handler"
	shipsHandler "godoksip/ships/handler"
)

func main() {
	port := "8080"
	conStr := "root:root@tcp(127.0.0.1:3306)/ship_dock"

	db, err := sql.Open("mysql", conStr)
	if err != nil {
		log.Fatal("Error When Connect to DB " + conStr + ": " + err.Error())
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	shipsRepo := repo.CreateShipsRepoMysqlImpl(db)
	shipsUsecase := usecase.CreateShipsUsecase(shipsRepo)
	docksRepo := repo.CreateDocksRepoMysqlImpl(db)
	docksUsecase := usecase.CreateDocksUsecase(docksRepo)

	shipsHandler.CreateShipsHandler(router, shipsUsecase)
	docksHandler.CreateDocksHandler(router, docksUsecase)

	router.Use(middleware.Logger)

	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Fatal(err)
	}


}
