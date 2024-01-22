package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pramodkoujalagi/REST-API-Mongo/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	err := http.ListenAndServe("localhost:9000", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(fmt.Errorf("unable to connect to MongoDB: %v", err))
	}
	return s
}
