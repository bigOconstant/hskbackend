package models

type Connection struct {
	Hosts      string `bson:"Hosts" json:"Hosts"`
	Database   string `bson:"Database" json:"Database"`
	Username   string `bson:"Username" json:"Username"`
	Password   string `bson:"Password" json:"Password"`
	Collection string `bson:"Collection" json:"Collection"`
	Origin1    string `bson:"Origin1" json:"Origin1"`
	Origin2    string `bson:"Origin2" json:"Origin2"`
	Prod       bool   `bson:"Prod" json:"Prod"`
	Port       string `bson:"Port" json:"Port"`
}
