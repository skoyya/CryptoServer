package server
import (	"log"
			"net/http"
			"datastore"
			"server/router"
			"server/handlers"
		)


type Server struct {
	dataStore *datastore.DataStore
}

func NewServer(ds *datastore.DataStore) *Server {
	return &Server{ds}
}

func (s *Server) Run() {
	h := handlers.NewHandlers(s.dataStore)
    myRouter := router.NewRouter(h)
    log.Fatal(http.ListenAndServe(":8088", myRouter))
}
