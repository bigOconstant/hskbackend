package main

import (
	//libraries
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
	//localPackages
	"os"

	"./controllers"
	"./models"
)

func main() {

	conn := getConnection("./connection.json")

	fmt.Println(conn)

	fmt.Println("Starting Application!")
	info := &mgo.DialInfo{
		Addrs:    []string{conn.Hosts},
		Timeout:  60 * time.Second,
		Database: conn.Database,
		Username: conn.Username,
		Password: conn.Password,
	}
	fmt.Println("Attempting to connect to mongodb")
	session, err1 := mgo.DialWithInfo(info)

	if err1 != nil {
		panic(err1)
	}

	defer session.Close()

	mux := goji.NewMux()

	mux.HandleFunc(pat.Get("/getAllHsk"), api.AllHsk(session, conn))
	mux.HandleFunc(pat.Get("/pagedHsk"), api.PagedHsk(session, conn))
	mux.HandleFunc(pat.Get("/pagedcedict"), api.Pagedcedict(session, conn))
	mux.HandleFunc(pat.Post("/pagedcedictDefinitionSearch"), api.PagedcedictDefinitionSearch(session, conn))
	mux.HandleFunc(pat.Get("/Lesson"), api.GetLesson(session, conn))

	fmt.Println("Database Connection Successfull :) Listening on port ", conn.Port)
	http.ListenAndServe(conn.Port, mux)

}

//./connection.json
func getConnection(filename string) models.Connection {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c models.Connection
	json.Unmarshal(raw, &c)
	return c
}
