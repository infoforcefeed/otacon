package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleterry/otacon/config"
)

type OtaconApp struct {
	router *mux.Router
}

func New(c *config.Config) *OtaconApp {
	app := &OtaconApp{}
}

func (a OtaconApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
