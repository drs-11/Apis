package handlers

import (
	"MightyLiteDB/internal/utils"
	"MightyLiteDB/pkg/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createData(w http.ResponseWriter, r *http.Request) {
	var newData storage.Data
	utils.ConvertJsonToData(r.Body, &newData)
	storage.DataMap[newData.Key] = &newData

	w.WriteHeader(http.StatusOK)
	log.Println("Data created:", newData.String())
}

func getData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(utils.ConvertDataToJson(storage.DataMap[key]))
	log.Println("Found requested Data:", storage.DataMap[key].String())
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	val, ok := storage.DataMap[key]
	if !ok {
		log.Println("Data not found..cant't delete.")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Key not found!"))
		return
	}
	delete(storage.DataMap, key)
	log.Println("Data found. Deleting...")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(val.String()))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Is this working?")
	log.Println("Endpoint Hit: Homepage")
}

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/store", createData).Methods("POST")
	router.HandleFunc("/store/{key}", getData).Methods("GET")
	router.HandleFunc("/store/{key}", deleteData).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
