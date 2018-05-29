package component

import (
	"../models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Cedictcomponent struct {
	s    *mgo.Session
	conn models.Connection
}

func NewCedictComponent(sessioninput *mgo.Session, connection models.Connection) *Cedictcomponent {
	CC := new(Cedictcomponent)
	CC.s = sessioninput
	CC.conn = connection
	return CC
}

func (m *Cedictcomponent) GetPagedHsk(hskLevel string, pageSize int, pageNumber int) models.CEDICTWITHSIZE {
	var level = "hsk1"
	switch hskLevel {
	case "1":
		level = "hsk1"
	case "2":
		level = "hsk2"
	case "3":
		level = "hsk3"
	case "4":
		level = "hsk4"
	case "5":
		level = "hsk5"
	case "6":
		level = "hsk6"
	default:
		level = ""
	}

	session := m.s.Copy()

	col := session.DB(m.conn.Database).C("cedict")

	var hsk []models.CEDICT
	var count = 0
	count, err := col.Find(bson.M{"Level": level}).Count()
	if err != nil {
		fmt.Println("Effor getting count")
		count = 0
	}
	q := col.Find(bson.M{"Level": level}).Sort("Simplified").Limit(pageSize)
	q = q.Skip((pageNumber - 1) * pageSize)
	err = q.All(&hsk)

	var response = models.CEDICTWITHSIZE{hsk, count}
	return response

}

func (m *Cedictcomponent) GetHskLevelCount(level string) int {
	fmt.Println("You are requesting level:", level)

	session := m.s.Copy()

	col := session.DB(m.conn.Database).C("cedict")
	count, err := col.Find(bson.M{"Level": level}).Count()

	if err != nil {
		log.Fatal(err)
		return -1
	}
	fmt.Println("You got a count of:", count)

	return count
}
