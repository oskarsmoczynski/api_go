package main

import (
	"api/db"
	"api/router"
	"api/server"
	"api/service"

	"net/http"
)

func main() {
	db := db.InitDb()
	rtr := router.CreateDefaultRouter()
	router.AppendRoute(rtr, db, http.MethodGet, "/read-table", service.ReadTable)
	router.AppendRoute(rtr, db, http.MethodPost, "/create-entry", service.CreateEntry)
	s := server.CreateServer("localhost:9090", rtr)
	server.Run(s)
}
