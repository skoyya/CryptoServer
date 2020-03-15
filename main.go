package main

import ( "datastore"
		 "server"
		 "log"
	 )

func main() {
	log.Print("CryptoServer started")
	validSymbols := []string{"ETHBTC"}
    ds := datastore.NewDataStore(validSymbols)
    ds.Run()

	s := server.NewServer(ds)
	s.Run()
}

