package main

import (
	"handler"
	"log"
	"net/http"
	"userHandler"
)

func main() {
    http.HandleFunc("/status", handler.StatusHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))


		users := make(map[string]userHandler.User)
    u1 := userHandler.User{
        ID:        "u1",
        FirstName: "Misha",
        LastName:  "Popov",
    }
    u2 := userHandler.User{
        ID:        "u2",
        FirstName: "Sasha",
        LastName:  "Popov",
    }
    users["u1"] = u1
    users["u2"] = u2

    http.HandleFunc("/users", userHandler.UserViewHandler(users))
    log.Fatal(http.ListenAndServe(":8080", nil))
}