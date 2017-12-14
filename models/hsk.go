package models
type HSK struct {
    Hanzi string `bson:"Hanzi" json:"Hanzi"`
	Pinyin string `bson:"Pinyin" json:"Pinyin"` 
	Definition string `bson:"Definition" json:"Definition"`
	Level string `bson:"Level" json:"Level"`
}