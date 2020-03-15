package handlers

import (
	"log"
	"datastore"
	"encoding/json"
	"net/http"
	"mux"
)

type Handlers struct {
	dataStore *datastore.DataStore
}

func NewHandlers(data *datastore.DataStore) *Handlers {
	return &Handlers{data}
}

func (h *Handlers) AddNewSymbol(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]
	h.dataStore.AddNewSymbol(symbol)
}

func (h *Handlers) GetAll(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := h.dataStore.GetAllSymbolData()
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal("Error in GetAll ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Error while serving getting the data"}`))
    }else {
	   w.WriteHeader(http.StatusOK)
	}
}

func (h *Handlers) GetSymbolData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	symbol := vars["symbol"]
	data := h.dataStore.GetSymbolData(symbol)
	if len(data) == 0 {
		errMsg := "No data found for symbol " + symbol
		log.Printf(errMsg)
		w.WriteHeader(http.StatusNotFound)
		jsonErr , _ := json.Marshal(map[string]string{"message": errMsg})
		w.Write([]byte(jsonErr))
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal("Error in GetAll ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Error while serving getting the data"}`))
    }else {
	   w.WriteHeader(http.StatusOK)
	}
}
