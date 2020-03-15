package router

import (
	"net/http"
	"server/handlers"
	"mux"
	"time"
	"log"
	)

func NewRouter(h *handlers.Handlers) *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/symbols/all").Handler(httpWrapper(h.GetAll,"GetAll"))
	router.Methods("GET").Path("/symbols/{symbol}").Handler(httpWrapper(h.GetSymbolData,"GetSymbolData"))
	router.Methods("PUT").Path("/symbols/{symbol}").Handler(httpWrapper(h.AddNewSymbol,"AddNewSymbol"))

    return router
}

func httpWrapper(inner http.HandlerFunc, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        inner.ServeHTTP(w, r)
        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}

