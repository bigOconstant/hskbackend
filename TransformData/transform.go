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
    collection = "hsk"
)
// type HSK struct {
// 	Hanzi string `bson:"Hanzi" json:"Hanzi"`
// 	Pinyin string `bson:"Pinyin" json:"Pinyin"` 
// 	Definition string `bson:"Definition" json:"Definition"`
// }
type HSKLevel struct {
	Hanzi string `bson:"Hanzi" json:"Hanzi"`
	Pinyin string `bson:"Pinyin" json:"Pinyin"` 
	Definition string `bson:"Definition" json:"Definition"`
	Level string `bson:"Level" json:"Level"`
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

	
		
		col := session.DB(database).C("hsk")

		var hsk []HSKLevel
		

		

		err := col.Find(bson.M{}).All(&hsk)
		if err != nil {
           fmt.Println("Error")
            return
		}

		
		// for i := 0; i < len(hsk); i++{
		// 	//col.Insert(hsk1[i]);
		// 	//fmt.Println(pages[i])
			
			
		// }
		pagesJson, err := json.Marshal(hsk)
		ioutil.WriteFile("C:/Users/caleb/Documents/GolangCode/hskParser/hskFull.json",pagesJson, 0644)
		
	

		
	}


