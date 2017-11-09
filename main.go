package main

import (
	"time"
	"goji.io"
	"goji.io/pat"
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2"
	
   
)
const (
    hosts      = "localhost:27017"
    database   = "my_database"
    username   = "dev1"
    password   = "password123"
	collection = "hsk"
	
)

type HSK struct {
    Hanzi string `bson:"Hanzi" json:"Hanzi"`
	Pinyin string `bson:"Pinyin" json:"Pinyin"` 
	Definition string `bson:"Definition" json:"Definition"`
	Level string `bson:"Level" json:"Level"`
}

func main() {

	fmt.Println("Starting Application!")
	
		info := &mgo.DialInfo {
			Addrs:    []string{hosts},
			Timeout:  60 * time.Second,
			Database: database,
			Username: username,
			Password: password,
		}

		session, err1 := mgo.DialWithInfo(info)

		if err1 != nil {
			panic(err1)
		}

		defer session.Close()
	
		mux := goji.NewMux()

	//mux.HandleFunc(pat.Get("/"), allUsers(session))
	
	mux.HandleFunc(pat.Get("/getAllHsk"), allHsk(session))
	mux.HandleFunc(pat.Get("/pagedHsk"), pagedHsk(session))
	//mux.HandleFunc(pat.Post("/adduser"), createUser(session))
	fmt.Println("Starting server listen and serve!")
	http.ListenAndServe(":8000", mux)

	
}