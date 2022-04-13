package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

//HandlerController function
func HandlerController() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cluster/{clusterId}", HandleCluster).Methods(http.MethodPut)

	return r
}
