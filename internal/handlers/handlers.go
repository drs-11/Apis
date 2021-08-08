package handlers

import (
	"Apis/internal/utils"
	"Apis/pkg/storage"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	DB *storage.DataBase
}

func (h *Handler) HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user", h.createUser).Methods("POST")
	router.HandleFunc("/user/database", h.createTable).Methods("POST")
	router.HandleFunc("/user/database/items", h.putData).Methods("POST")
	router.HandleFunc("/user/database/items", h.getData).Methods("GET")
	router.HandleFunc("/user/database/items", h.deleteData).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var payload storage.PayLoad
	var data []byte
	r.Body.Read(data)
	json.Unmarshal(data, &payload)
	h.DB.CreateUser(payload.Username, payload.Password)
}

func (h *Handler) createTable(w http.ResponseWriter, r *http.Request) {
	var payload storage.PayLoad
	var data []byte
	r.Body.Read(data)
	json.Unmarshal(data, &payload)
	h.DB.CreateTable(&payload)
}

func (h *Handler) putData(w http.ResponseWriter, r *http.Request) {
	var payload storage.PayLoad
	var data []byte
	r.Body.Read(data)
	json.Unmarshal(data, &payload)
	h.DB.AddItems(&payload)
}

func (h *Handler) getData(w http.ResponseWriter, r *http.Request) {
	var payload storage.PayLoad
	var data []byte
	r.Body.Read(data)
	json.Unmarshal(data, &payload)
	items := h.DB.GetItems(&payload)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(utils.ConvertDataToJson(items))
}

func (h *Handler) deleteData(w http.ResponseWriter, r *http.Request) {
	var payload storage.PayLoad
	var data []byte
	r.Body.Read(data)
	json.Unmarshal(data, &payload)
	h.DB.DeleteItems(&payload)
}
