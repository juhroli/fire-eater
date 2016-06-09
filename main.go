package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"encoding/json"
)

type (
	FireController struct {
		session *mgo.Session
	}
)

func NewFireController(s *mgo.Session) *FireController {
	return &FireController{s}
}


func (fc FireController) handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	content_type := r.Header.Get("Content-Type")
	if content_type != "application/json" || content_type != "text/plain" {
		w.WriteHeader(400)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	var json_body map[string]interface{}

	json.Unmarshal(buf, &json_body)
	source := r.Header.Get("source")
	fc.session.DB("FireEaterData").C(source).Insert(json_body)
	w.WriteHeader(200)
}

func getMongoSession() *mgo.Session {
	mongo := os.Getenv("MONGO_PORT")[6:]
	fmt.Println("Using mongo: ", mongo)
	s, error := mgo.Dial("mongodb://" + mongo)

	if error != nil {
		panic(error)
	}

	return s
}


func main() {
	r := httprouter.New()
	fc := NewFireController(getMongoSession())

	r.POST("/consume", fc.handler)

	port := string(os.Getenv("PORT"))

	if port == "" {
		port = "9081"
	}

	fmt.Println("Server started listening on port: ", port)

	http.ListenAndServe(":" + port, r)
}
