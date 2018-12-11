package main

import (
	"log"
	"net/http"
    "os"
	"github.com/dimfeld/httptreemux"
    "github.com/globalsign/mgo"
	"github.com/richardkevin/go-beers/api"
	"github.com/richardkevin/go-beers/beers"
)

func main() {
    PORT := os.Getenv("PORT")
	addr := "127.0.0.1:" + PORT
	router := httptreemux.NewContextMux()

    session, err := mgo.Dial("localhost:27017/go-beers")
	if err != nil {
		log.Fatal(err)
	}
	repository := beers.NewBeerRepository(session)

	router.Handler("GET", "/", &api.DefaultHandler{repository})
	router.Handler(http.MethodGet, "/beer/:id", &api.GetBeerHandler{repository})
	router.Handler(http.MethodGet, "/create/:name", &api.UpsertBeerHandler{repository})

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))

	// execute
	// curl http://localhost:8081/beer/heineken
	// curl -XPUT http://localhost:8081/beer/heineken -d'{"name": "heineken"}'
}
