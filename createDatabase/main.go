//Script to insert data for project
//Need to think of a cleaner way to do it

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

const (
	hosts      = "localhost:27017"
	database   = "my_database"
	username   = "dev1"
	password   = "password123"
	collection = "cedict"
)

type CEDICTSTRUCT struct {
	Traditional    string   `bson:"Traditional" json:"Traditional"`
	Simplified     string   `bson:"Simplified" json:"Simplified"`
	PinyinNumbered string   `bson:"PinyinNumbered" json:"PinyinNumbered"`
	Pinyin         string   `bson:"Pinyin" json:"Pinyin"`
	Definition     string   `bson:"Definition" json:"Definition"`
	Search         []string `bson:"Search" json:"Search"`
}

type HSK struct {
	Hanzi      string `bson:"Hanzi" json:"Hanzi"`
	Pinyin     string `bson:"Pinyin" json:"Pinyin"`
	Definition string `bson:"Definition" json:"Definition"`
	Level      string `bson:"Level" json:"Level"`
}

func (p CEDICTSTRUCT) toString() string {
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

	fmt.Println("Hi there! Lets create those collections...")

	info := &mgo.DialInfo{
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

	fmt.Println("Beginning read of cedict file...")

	CedictPages := getCedictPages("./Data/cedict.json")

	fmt.Println("Cedict Length:", len(CedictPages))
	fmt.Println("Begenning inserting cedict documents into mongodb. Hold on a sec this could take a minute...")

	for i := 0; i < len(CedictPages); i++ {
		col.Insert(CedictPages[i])
		//fmt.Println(pages[i])
	}
	fmt.Println("Done inserting Cedict")
	fmt.Println("")

	col = session.DB(database).C("hsk")

	fmt.Println("Beginning read of HSK file...")

	HskPages := getHskPages("./Data/hskAll.json")

	fmt.Println("hsk Length:", len(HskPages))

	fmt.Println("Begenning inserting hsk documents into mongodb...")

	for i := 0; i < len(HskPages); i++ {
		col.Insert(HskPages[i])
		//fmt.Println(pages[i])
	}

	fmt.Println("Done inserting Hsk")
	fmt.Println("")

	fmt.Println("Done Creating database.")

}

func getCedictPages(directory string) []CEDICTSTRUCT {
	raw, err := ioutil.ReadFile(directory)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var pages []CEDICTSTRUCT
	json.Unmarshal(raw, &pages)

	for i := 0; i < len(pages); i++ {
		pages[i].Search = append(pages[i].Search, pages[i].Traditional)
		pages[i].Search = append(pages[i].Search, pages[i].Simplified)

		PinyinNumbered := strings.Fields(pages[i].PinyinNumbered)

		pages[i].Search = append(pages[i].Search, PinyinNumbered...)

		Pinyin := strings.Fields(pages[i].Pinyin)

		pages[i].Search = append(pages[i].Search, Pinyin...)
		pages[i].Definition = strings.Replace(pages[i].Definition, ";", " ", -1)

		definitionCleaned := strings.Replace(pages[i].Definition, ";", " ", -1)
		definitionCleaned = strings.Replace(definitionCleaned, "(", " ", -1)
		definitionCleaned = strings.Replace(definitionCleaned, ")", " ", -1)
		definitionCleaned = strings.ToLower(definitionCleaned)

		definitionSplit := strings.Fields(definitionCleaned)

		pages[i].Search = append(pages[i].Search, definitionSplit...)

	}

	return pages
}

func getHskPages(directory string) []HSK {
	raw, err := ioutil.ReadFile(directory)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var pages []HSK
	json.Unmarshal(raw, &pages)

	return pages

}
