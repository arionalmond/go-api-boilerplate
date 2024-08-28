package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/arionalmond/go-api-boilerplate/pkg/datastore"
)

// GetEmployees ...
func GetEmployees(ds *datastore.MySQLDS) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "id incorrect", http.StatusBadRequest)
			return
		}

		emps, err := ds.GetEmployeeByID(id)
		err = json.NewEncoder(w).Encode(emps)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "error enconding employee to json", http.StatusInternalServerError)
			return
		}

	}
}
