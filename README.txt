Running:- 
export GOPATH=`workspace location`
go run main.go from workspace

Testing:---
To get the price of ETHBTC
curl -X GET 'http://localhost:8088/symbols/ETHBTC'

To get the price of all supported symbols
curl -X GET 'http://localhost:8088/symbols/all'

To add the symbol to the CryptoServer to support, Note it is PUT call. Soon after this call, symbol  LTCTUSD details are fetched and cached
curl -X PUT 'http://localhost:8088/symbols/LTCTUSD'
