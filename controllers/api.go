package api

import (
	"../components"
	"../models"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(json)
}

func PagedHsk(s *mgo.Session, conn models.Connection) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()

		origin := r.Header.Get("Origin")

		if conn.Prod && origin != conn.Origin1 && origin != conn.Origin2 {
			ErrorWithJSON(w, "Database error", http.StatusBadRequest)
			log.Println("Not allowed ", nil)
			return
		} else {

			hskLevel := r.URL.Query().Get("hskLevel")
			ccomponent := component.NewCedictComponent(session, conn)
			pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
			pageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))

			var response = ccomponent.GetPagedHsk(hskLevel, pageSize, pageNumber)

			respBody, err := json.MarshalIndent(response, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

			ResponseWithJSON(w, respBody, http.StatusOK)

		}
	}
}

func Pagedcedict(s *mgo.Session, conn models.Connection) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		origin := r.Header.Get("Origin")

		if conn.Prod && origin != conn.Origin1 && origin != conn.Origin2 {
			ErrorWithJSON(w, "Database error", http.StatusBadRequest)
			log.Println("Not allowed ", nil)
			return
		} else {
			pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))

			pageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))

			col := session.DB(conn.Database).C("cedict")

			var cedict []models.CEDICT
			count, err := col.Count()

			if err != nil {
				log.Fatal(err)
			}
			q := col.Find(bson.M{}).Sort("Pinyin").Limit(pageSize)
			q = q.Skip((pageNumber - 1) * pageSize)
			err = q.All(&cedict)

			var response = models.CEDICTWITHSIZE{cedict, count}

			respBody, err := json.MarshalIndent(response, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

			ResponseWithJSON(w, respBody, http.StatusOK)

		}
	}
}

func AllHsk(s *mgo.Session, conn models.Connection) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		origin := r.Header.Get("Origin")

		if conn.Prod && origin != conn.Origin1 && origin != conn.Origin2 {
			ErrorWithJSON(w, "Database error", http.StatusBadRequest)
			log.Println("Not allowed ", nil)
			return
		} else {
			hskLevel := r.URL.Query().Get("hskLevel")

			var colectionvalue = "hsk"
			switch hskLevel {
			case "1":
				colectionvalue = "hsk1"
			case "2":
				colectionvalue = "hsk2"
			case "3":
				colectionvalue = "hsk3"
			case "4":
				colectionvalue = "hsk4"
			case "5":
				colectionvalue = "hsk5"
			case "6":
				colectionvalue = "hsk6"
			default:
				colectionvalue = ""
			}

			col := session.DB(conn.Database).C(conn.Collection)

			var hsk []models.HSK

			if colectionvalue != "" {

				err := col.Find(bson.M{"Level": colectionvalue}).All(&hsk)
				if err != nil {
					ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
					log.Println("Failed get all objects: ", err)
					return
				}
			} else {
				err := col.Find(bson.M{}).All(&hsk)
				if err != nil {
					ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
					log.Println("Failed get all objects: ", err)
					return
				}
			}

			respBody, err := json.MarshalIndent(hsk, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

			ResponseWithJSON(w, respBody, http.StatusOK)

		}
	}

}

func PagedcedictDefinitionSearch(s *mgo.Session, conn models.Connection) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		body, err := ioutil.ReadAll(r.Body)
		origin := r.Header.Get("Origin")

		if conn.Prod && origin != conn.Origin1 && origin != conn.Origin2 {
			ErrorWithJSON(w, "Database error", http.StatusBadRequest)
			log.Println("Not allowed ", nil)
			return
		} else {
			var t models.Search
			err = json.Unmarshal(body, &t)
			if err != nil {
				panic(err)
			}
			pageSize := t.PageSize
			pageNumber := t.Page + 1
			var colectionvalue = "cedict"
			col := session.DB(conn.Database).C(colectionvalue)
			var cedict []models.CEDICT
			var stringfields = strings.Split(t.Search, " ")
			q := col.Find(bson.M{"Search": bson.M{"$all": stringfields}})
			count, err := q.Count()
			if err != nil {
				log.Fatal(err)
			}
			q = q.Limit(pageSize)
			q = q.Skip((pageNumber - 1) * pageSize)
			err = q.All(&cedict)
			var response = models.CEDICTWITHSIZE{cedict, count}
			respBody, err := json.MarshalIndent(response, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			ResponseWithJSON(w, respBody, http.StatusOK)
		}

	}
}

func GetLesson(s *mgo.Session, conn models.Connection) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()

		origin := r.Header.Get("Origin")

		if conn.Prod && origin != conn.Origin1 && origin != conn.Origin2 {
			ErrorWithJSON(w, "Database error", http.StatusBadRequest)
			log.Println("Not allowed", nil)
			return
		} else {

			lessonNumber, err := strconv.Atoi(r.URL.Query().Get("lesson"))

			col := session.DB(conn.Database).C("lessons")

			var lessons []models.Lesson

			q := col.Find(bson.M{"Lesson": lessonNumber})

			err = q.All(&lessons)

			var lesson models.Lesson

			if len(lessons) > 0 {
				lesson = lessons[0]
			}

			respBody, err := json.MarshalIndent(lesson, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

			ResponseWithJSON(w, respBody, http.StatusOK)

		}
	}
}
