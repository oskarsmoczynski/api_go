package main

import (
	"api/db"
	"api/router"
	"api/server"
	"api/service"
	u "api/utils"

	"net/http"
)

func main() {
	u.InitLogger()
	u.StdLogger.Println("abc")
	db := db.InitDb()
	rtr := router.CreateDefaultRouter()
	router.AppendRoute(rtr, db, http.MethodGet, "/read-table", service.ReadTable)
	router.AppendRoute(rtr, db, http.MethodPost, "/create-entry", service.CreateEntry)
	router.AppendRoute(rtr, db, http.MethodPatch, "/update-entry", service.UpdateEntry)
	router.AppendRoute(rtr, db, http.MethodDelete, "/delete-entry", service.DeleteEntry)
	s := server.CreateServer("localhost:9090", rtr)
	server.Run(s)
}
