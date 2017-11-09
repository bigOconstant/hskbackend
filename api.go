
package main

import (
	
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"log"
	"strconv"
//	"common/controller.go"
   
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


func pagedHsk(s *mgo.Session) func(w http.ResponseWriter,r *http.Request){
    return func(w http.ResponseWriter,r *http.Request){
		session := s.Copy()
		hskLevel := r.URL.Query().Get("hskLevel")
		pageSize,err := strconv.Atoi(r.URL.Query().Get("pageSize"))
		

		pageNumber,err := strconv.Atoi(r.URL.Query().Get("page"))

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
		col := session.DB(database).C(collection)

		var hsk []HSK

		q := col.Find(bson.M{"Level":colectionvalue}).Limit(pageSize);
		q = q.Skip((pageNumber-1)*pageSize)
		err = q.All(&hsk)
		respBody, err := json.MarshalIndent(hsk, "", "  ")
        if err != nil {
            log.Fatal(err)
		}
		
		ResponseWithJSON(w, respBody, http.StatusOK)





	}
}



func allHsk(s *mgo.Session) func(w http.ResponseWriter,r *http.Request){
    return func(w http.ResponseWriter,r *http.Request){
		session := s.Copy()
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
		col := session.DB(database).C(collection)

		var hsk []HSK

		if colectionvalue != "" {

		err := col.Find(bson.M{"Level":colectionvalue}).All(&hsk)
		if err != nil {
            ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
            log.Println("Failed get all objects: ", err)
            return
        }
		}else{
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




// func createUser(s *mgo.Session) func(w http.ResponseWriter,r *http.Request){
//     return func(w http.ResponseWriter,r *http.Request){
// 		session := s.Copy()

// 		var user Users
// 		var userCheck Users

// 		decoder := json.NewDecoder(r.Body)
// 		err := decoder.Decode(&user)

// 		if err != nil {
//             ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
//             return
// 		}
		
// 		c := session.DB(database).C(collection)

// 		c.Find(bson.M{"user": user.User}).One(&userCheck)
// 		fmt.Println(userCheck.User)
		
// 		if (Users{}) == userCheck {
// 		    err = c.Insert(user)
// 		} else {log.Println("failed, user exists")}

// 		if err != nil  {
//             if mgo.IsDup(err) {
//                 ErrorWithJSON(w, "User already exists", http.StatusBadRequest)
//                 return
//             }
	   

// 	   ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
// 	       log.Println("Failed insert user: ", err)
// 	      return
//       }
// 	  w.Header().Set("Content-Type", "application/json")
// 	  w.WriteHeader(http.StatusCreated)
// 	}
// }