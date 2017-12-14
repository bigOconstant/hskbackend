package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"os"
	"gopkg.in/mgo.v2"
    "time"
    "strings"
    
)

const (
    hosts      = "localhost:27017"
    database   = "hanyu"
    username   = "prodcaleb"
    password   = "studyhanyudev"
    collection = "cedict"
)

type HSK struct {
	Traditional string `bson:"Traditional" json:"Traditional"`
	Simplified string `bson:"Simplified" json:"Simplified"`
	PinyinNumbered string `bson:"PinyinNumbered" json:"PinyinNumbered"`
	Pinyin string `bson:"Pinyin" json:"Pinyin"`
    Definition string `bson:"Definition" json:"Definition"`
    Search[] string `bson:"Search" json:"Search"`
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
    // for _, p := range pages {
    //     fmt.Println(p.toString())
    // }

	//fmt.Println(toJson(pages))
	
    fmt.Println("Length:",len(pages))
    
    for i := 0; i < len(pages); i++{
        pages[i].Search = append(pages[i].Search,pages[i].Traditional)
        pages[i].Search = append(pages[i].Search,pages[i].Simplified)

        PinyinNumbered := strings.Fields(pages[i].PinyinNumbered)

        pages[i].Search = append(pages[i].Search,PinyinNumbered...)

        Pinyin := strings.Fields(pages[i].Pinyin)

        pages[i].Search = append(pages[i].Search,Pinyin...)
        pages[i].Definition = strings.Replace(pages[i].Definition,";"," ",-1)

        definitionCleaned := strings.Replace(pages[i].Definition,";"," ",-1)
        definitionCleaned = strings.Replace(definitionCleaned,"("," ",-1)
        definitionCleaned = strings.Replace(definitionCleaned,")"," ",-1)
        definitionCleaned = strings.ToLower(definitionCleaned);

        definitionSplit := strings.Fields(definitionCleaned)

        pages[i].Search = append(pages[i].Search,definitionSplit...)
		
    }

	for i := 0; i < len(pages); i++{
		col.Insert(pages[i])
		//fmt.Println(pages[i])
    }
    fmt.Println("Done inserting")

}

func getPages() []HSK {
    raw, err := ioutil.ReadFile("./cedict.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []HSK
    json.Unmarshal(raw, &c)
    return c
}