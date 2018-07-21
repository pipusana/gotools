package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string
	Phone   string
	ADDRESS string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world !!")
}

func createProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data Person
	body, _ := ioutil.ReadAll(r.Body)
	error := json.Unmarshal(body, &data)

	if error != nil {
		log.Println(error)
	}

	session, err := mgo.Dial("mongodb://localhost:27017")
	connect := session.DB("address").C("person")

	Info := Person{Name: data.Name, Phone: data.Phone, ADDRESS: data.ADDRESS}
	err = connect.Insert(Info)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := json.Marshal(Info)
	if err != nil {
		log.Println(err)
	}

	w.Write(payload)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	connect := session.DB("address").C("person")

	if err != nil {
		log.Println(err)
	}

	var results []Person
	error := connect.Find(nil).All(&results)
	if error != nil {
		log.Fatal(error)
	}

	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

	payload, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
	}

	w.Write(payload)
}

func getProfileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mongoID := strings.TrimPrefix(vars["id"], "id")
	session, err := mgo.Dial("mongodb://localhost:27017")
	connect := session.DB("address").C("person")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(mongoID)

	var results []Person
	error := connect.FindId(bson.M{"_id": bson.ObjectIdHex(mongoID)}).One(&results)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(results)

	payload, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
	}

	w.Write(payload)
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mongoID := strings.TrimPrefix(vars["id"], "id")
	session, err := mgo.Dial("mongodb://localhost:27017")
	connect := session.DB("address").C("person")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(mongoID)

	connect.Remove(bson.M{"_id": bson.ObjectIdHex(mongoID)})
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mongoID := strings.TrimPrefix(vars["id"], "id")
	w.Header().Set("Content-Type", "application/json")

	var data Person
	body, _ := ioutil.ReadAll(r.Body)
	error := json.Unmarshal(body, &data)

	if error != nil {
		log.Println(error)
	}

	session, err := mgo.Dial("mongodb://localhost:27017")
	connect := session.DB("address").C("person")

	fmt.Println(mongoID)

	err = connect.Update(bson.M{"_id": bson.ObjectIdHex(mongoID)}, bson.M{"$set": bson.M{"name": data.Name, "address": data.ADDRESS, "phone": data.Phone}})

	if err != nil {
		log.Println(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).
		Methods("GET")

	r.HandleFunc("/find", getProfile).
		Methods("GET")

	r.HandleFunc("/findById/{id}", getProfileByID).
		Methods("GET")

	r.HandleFunc("/create", createProfile).
		Methods("POST")

	r.HandleFunc("/delete/{id}", deleteProfile).
		Methods("DELETE")

	r.HandleFunc("/update/{id}", updateProfile).
		Methods("PUT")

	http.ListenAndServe(":3000", r)
}
