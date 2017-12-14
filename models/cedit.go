package models

type CEDICT struct {
	Traditional string `bson:"Traditional" json:"Traditional"`
	Simplified string `bson:"Simplified" json:"Simplified"`
	PinyinNumbered string `bson:"PinyinNumbered" json:"PinyinNumbered"`
	Pinyin string `bson:"Pinyin" json:"Pinyin"`
	Definition string `bson:"Definition" json:"Definition"`
	Search[] string `bson:"Search" json:"Search"`
}
type CEDICTWITHSIZE struct{
	Data[] CEDICT `bson:"Data" json:"Data"`
	Size int `bson:"Size" json:"Size"`
}
type Search struct {
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
	Search string `json:"search"`
}