package server

import "net/http"


type IDroneAPI interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type droneAPI struct {

}

func NewDroneAPI() IDroneAPI{
	return &droneAPI{}
}

func (api *droneAPI) Get(w http.ResponseWriter, r * http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
