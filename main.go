package main

import (
	"time"
	"goji.io"
	"goji.io/pat"
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2"
	"./models"
	"./controllers"
)

func main() {
	conn := models.Connection{
		"localhost:27017",
		"my_database",
		"dev1",
		"password123",
		"cedict",
		"http://localhost:4200",
		"http://www.localhost:4200"}

	fmt.Println("Starting Application!")	
		info := &mgo.DialInfo {
			Addrs:    []string{conn.Hosts},
			Timeout:  60 * time.Second,
			Database: conn.Database,
			Username: conn.Username,
			Password: conn.Password,
		}

		session, err1 := mgo.DialWithInfo(info)

		if err1 != nil {
			panic(err1)
		}

		defer session.Close()
	
		mux := goji.NewMux()

	
	mux.HandleFunc(pat.Get("/getAllHsk"), api.AllHsk(session,conn))
	mux.HandleFunc(pat.Get("/pagedHsk"), api.PagedHsk(session,conn))
	mux.HandleFunc(pat.Get("/pagedcedict"), api.Pagedcedict(session,conn))
	mux.HandleFunc(pat.Post("/pagedcedictDefinitionSearch"), api.PagedcedictDefinitionSearch(session,conn))

	fmt.Println("Starting server listen and serve!")
	http.ListenAndServe(":8000", mux)

	
}