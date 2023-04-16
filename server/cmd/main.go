package main

import (
	"e/server/db"
	"log"
	"e/server/internal/user"
	"e/server/internal/ws"
	"e/server/router"
)

const PORT = ":8080"
func main() {
	
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
	

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("127.0.0.1:8080")
}	
