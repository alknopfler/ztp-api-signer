package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleCluster(w http.ResponseWriter, r *http.Request) {
	if mux.Vars(r)["clusterId"] != "" {
		result := joinToCluster(mux.Vars(r)["clusterId"])
		if !result {
			log.Print("ERROR 404: Error Joining the cluster id: " + mux.Vars(r)["clusterId"] + " to hub cluster")
			responseWithError(w, http.StatusExpectationFailed, "Error Joining the cluster id: "+mux.Vars(r)["clusterId"]+" to hub cluster")
			return
		}
		log.Print("INFO 200 Joining successfully cluster id: " + mux.Vars(r)["clusterId"] + " to hub cluster")
		responseWithJSON(w, http.StatusOK, "Joining successfully cluster id: "+mux.Vars(r)["clusterId"]+" to hub cluster")
		return
	}
	log.Print("ERROR 400: Error Joining the cluster id: " + mux.Vars(r)["clusterId"] + " to hub cluster. URI param is not valid")
	responseWithError(w, http.StatusExpectationFailed, "The URI param is empty")
	return
}

func joinToCluster(clusterId string) bool {
	return true
}
