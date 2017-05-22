package tracker

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/infoforcefeed/otacon/config"
)

type Tracker struct {
	router *mux.Router
	config *config.Config
}

func New(cfg *config.Config) *Tracker {
	t := &Tracker{config: cfg}

	t.router = mux.NewRouter()

	t.router.Handle("/{token}/announce", errorHandler(t.announceHandler)).Methods("GET").Name("announce")
	t.router.Handle("/{token}/scrape", errorHandler(t.scrapeHandler)).Methods("GET").Name("scrape")

	return t
}

func (t *Tracker) Run() error {
	loggedRouter := handlers.LoggingHandler(os.Stdout, t)
	fmt.Printf("tracker running on %s\n", t.config.TrackerBind)
	if err := http.ListenAndServe(t.config.TrackerBind, loggedRouter); err != nil {
		return err
	}

	return nil
}

func (t *Tracker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.router.ServeHTTP(w, r)
}
