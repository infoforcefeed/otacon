package tracker

import "github.com/gorilla/mux"

type Tracker struct {
	router mux.Router
}

func New() *Tracker {
	t := &Tracker{}

	t.router = mux.NewRouter()

	t.router.Handle("/announce", t.announceHandler).Methods("GET").Name("announce")
	t.router.Handle("/scrape", t.scrapeHandler).Methods("GET").Name("scrape")
}
