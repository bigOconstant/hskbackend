package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type HSK struct {
	Hanzi      string `bson:"Hanzi" json:"Hanzi"`
	Pinyin     string `bson:"Pinyin" json:"Pinyin"`
	Definition string `bson:"Definition" json:"Definition"`
	Level      string `bson:"Level" json:"Level"`
	Found      string `bson:"Found" json:"Found"`
}

type CEDICTSTRUCT struct {
	Traditional    string `bson:"Traditional" json:"Traditional"`
	Simplified     string `bson:"Simplified" json:"Simplified"`
	PinyinNumbered string `bson:"PinyinNumbered" json:"PinyinNumbered"`
	Pinyin         string `bson:"Pinyin" json:"Pinyin"`
	Definition     string `bson:"Definition" json:"Definition"`
	Level          string `bson:"Level" json:"Level"`
}

func main() {
	fmt.Println("Starting File!")
	hskPages := getHskPages("C:/Users/wolf/Documents/hsk/backend/hskbackend/detectLevel/hskAll.json")
	cedictPages := getCedictPages("C:/Users/wolf/Documents/hsk/backend/hskbackend/detectLevel/cedict.json")
	for i := 0; i < len(hskPages); i++ {
		//fmt.Println(hskPages[i].Hanzi)
		for j := 0; j < len(cedictPages); j++ {

			if cedictPages[j].Simplified == hskPages[i].Hanzi {
				//fmt.Println(cedictPages[j].Simplified, " Matched, its level ", hskPages[i].Level)
				cedictPages[j].Level = hskPages[i].Level
				hskPages[i].Found = "True"
			}
		}

	}

	counter := 1

	for i := 0; i < len(hskPages); i++ {
		if hskPages[i].Found != "True" {
			//	fmt.Println(counter)
			counter++
			//	fmt.Println("Hanzi:", hskPages[i].Hanzi)
			//	fmt.Println("Level:", hskPages[i].Level)
			//	fmt.Println("Pinyin:", hskPages[i].Pinyin)
			//	fmt.Println("Definition", hskPages[i].Definition, "\n")

		}

	}
	fmt.Println("not found:", counter)
	pagesJson, _ := json.MarshalIndent(cedictPages, "", " ")
	ioutil.WriteFile("cedictwithhsklabels.json", pagesJson, 0644)

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
func getCedictPages(directory string) []CEDICTSTRUCT {
	raw, err := ioutil.ReadFile(directory)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var pages []CEDICTSTRUCT
	json.Unmarshal(raw, &pages)

	return pages

}
