package  datastore

import ("log"
		"io/ioutil"
		"time"
		"sync"
		"net/http"
)

const BASE_URL = "https://api.hitbtc.com/api/2/public/symbol/"
const POLL_FREQUENCY_IN_SEC = 5

type DataStore struct {
	validSymbols []string
	data map[string]string
	mutex sync.RWMutex
}

func NewDataStore(vSymbols []string) *DataStore {
	return &DataStore{vSymbols, make(map[string]string), sync.RWMutex{}}
}

func (ds *DataStore) Run() {
	log.Printf("DataStore started -- will fetch the symbol data for every [%d] seconds",POLL_FREQUENCY_IN_SEC)
	ticker := time.NewTicker(POLL_FREQUENCY_IN_SEC * time.Second)
	go func() {
		for {
           select {
                case <- ticker.C:
					ds.loadCryptoData()
			}//select
		}//for ifnite
	}()
}

func (ds *DataStore) GetAllSymbolData() []string {
	data := []string{}
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()
	for _ , symbolData :=  range ds.data  {
		data = append(data, symbolData)
	}
	return data
}

func (ds *DataStore) AddNewSymbol(symbol string) {
	//append function , i think no need to synchronize it
	ds.validSymbols = append(ds.validSymbols, symbol)
}

func (ds *DataStore) GetSymbolData(symbol string) string {
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()
	//It is quicker to search in map that in list validSymbols
	if data, ok := ds.data[symbol]; ok {
		return data;
	}
	return ""
}

func (ds *DataStore) loadCryptoData() {
	newData := make(map[string]string)
	for _, symbol := range ds.validSymbols {
		resp, err := http.Get(BASE_URL + symbol)
		defer resp.Body.Close()
		var byteBody []byte
		if err == nil {
			byteBody, err = ioutil.ReadAll(resp.Body)
		}
		if err != nil || len(byteBody) == 0 {
			//Dont exit, next try might go through.  Need to count errors and exit if breaches the configured limit
			log.Fatal("Failed to fetch the data for symbol" , err)
		}
		newData[symbol] = string(byteBody)
		log.Print("Fetched new price for symbol: "+symbol)
	}//forloop
	//Not locking while loading as it adds the latency to APIs
	ds.mutex.Lock()
	ds.data = newData
	defer ds.mutex.Unlock()
}
