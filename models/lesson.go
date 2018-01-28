package models
type Convo struct {
	Pinyin string `bson:"Pinyin" json:"Pinyin"`
	Hanzi  string `bson:"Hanzi" json:"Hanzi"`
	Flag   bool   `bson:"Flag" json:"Flag"`
}

type Word struct {
	Pinyin     string `bson:"Pinyin" json:"Pinyin"`
	Hanzi      string `bson:"Hanzi" json:"Hanzi"`
	Definition string `bson:"Definition" json:"Definition"`
	Isnew      bool   `bson:"Isnew" json:"Isnew"`
}

type Lesson struct {
	Conversation []Convo
	Words        []Word
	Lesson       int `bson:"Lesson" json:"Lesson"`
}