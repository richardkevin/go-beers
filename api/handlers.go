package api

import (
	"fmt"
    "log"
    "net/http"
    "github.com/dimfeld/httptreemux"
    "github.com/richardkevin/go-beers/beers"
)

type DefaultHandler struct{
    Repository *beers.BeerRepository
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Breetings, lets have a beer! \n")

	fmt.Fprintln(w, "O que você deseja?")
	fmt.Fprintln(w, "Tomar uma: /beer/:id")
	fmt.Fprintln(w, "Feeling luck: /random")

}

type UpsertBeerHandler struct{
    Repository *beers.BeerRepository
}
func (h *UpsertBeerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	new_beer := &beers.Beer{Name: params["name"]}
    err := h.Repository.Create(new_beer)
    if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Cerveja %s adicionada com sucesso!", new_beer.Name)
}

type GetBeerHandler struct{
    Repository *beers.BeerRepository
}
func (h *GetBeerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
    beer, _ := h.Repository.FindById(params["id"])
	fmt.Fprintf(w, "Eu queria tomar a cerveja: %s!", beer.Name)
}
