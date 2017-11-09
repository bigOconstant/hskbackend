package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"os"
	"gopkg.in/mgo.v2"
	"time"
)

const (
    hosts      = "localhost:27017"
    database   = "my_database"
    username   = "dev1"
    password   = "password123"
    collection = "hsk6"
)

type HSK struct {
	Hanzi string `bson:"Hanzi" json:"Hanzi"`
	Pinyin string `bson:"Pinyin" json:"Pinyin"` 
	Definition string `bson:"Definition" json:"Definition"`
}

func (p HSK) toString() string {
    return toJson(p)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}

func main() {
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

			col := session.DB(database).C(collection)

    pages := getPages()
    for _, p := range pages {
        fmt.Println(p.toString())
    }

	fmt.Println(toJson(pages))
	
	fmt.Println("Length:",len(pages))

	for i := 0; i < len(pages); i++{
		col.Insert(pages[i]);
		//fmt.Println(pages[i])
	}

}

func getPages() []HSK {
    raw, err := ioutil.ReadFile("./hsk6.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []HSK
    json.Unmarshal(raw, &c)
    return c
}