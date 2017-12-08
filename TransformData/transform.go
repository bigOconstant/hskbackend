package main

import (
   // "encoding/json"
    "fmt"
    //"io/ioutil"
	//"os"
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"io/ioutil"
)
const (
    hosts      = "localhost:27017"
    database   = "my_database"
    username   = "dev1"
    password   = "password123"
    collection = "cedict"
)
// type HSK struct {
// 	Hanzi string `bson:"Hanzi" json:"Hanzi"`
// 	Pinyin string `bson:"Pinyin" json:"Pinyin"` 
// 	Definition string `bson:"Definition" json:"Definition"`
// }
type HSKLevel struct {
	Traditional string `bson:"Traditional" json:"Traditional"`
	Simplified string `bson:"Simplified" json:"Simplified"` 
	PinyinNumbered string `bson:"PinyinNumbered" json:"PinyinNumbered"`
	Pinyin string `bson:"Pinyin" json:"Pinyin"`
	Definition string `bson:"Definition" json:"Definition"`
	Search[] string `bson:"Search" json:"Search"`
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
		fmt.Println("Erro on connect")
		 return
	 }

	
		
		col := session.DB(database).C(collection)

		var cedict []HSKLevel
		

		

		err := col.Find(bson.M{}).All(&cedict)
		if err != nil {
           fmt.Println("Error")
            return
		}

		
		// for i := 0; i < len(hsk); i++{
		// 	//col.Insert(hsk1[i]);
		// 	//fmt.Println(pages[i])
			
			
		// }
		pagesJson, err := json.Marshal(cedict)
		ioutil.WriteFile("C:/Users/caleb/Documents/GolangCode/hskParser/hskFull.json",pagesJson, 0644)
		
	

		
	}


