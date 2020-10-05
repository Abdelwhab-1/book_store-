package datamarshal

import (
	"encoding/json"
	"log"
	"net/http"
)
func MarshUser(wr http.ResponseWriter, data interface{}){
	end := json.NewEncoder(wr)
	err := end.Encode(data)
	if  err != nil {
		log.Fatal(err)
	}
}
func UnmarshUser(r *http.Request, data interface{}){
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {

		log.Fatal(err)
	}
}