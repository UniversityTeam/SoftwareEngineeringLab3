package balancers

import (
	"encoding/json"
	"github.com/roman-mazur/chat-channels-example/server/tools"
	"log"
	"net/http"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleBalancersList(store, rw)
		} else if r.Method == "POST" {
			handleMachineUpdate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleMachineUpdate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var c Machine
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.UpdateMachine(c.Id, c.Worked)
	if err == nil {
		tools.WriteJsonOk(rw, &c)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleBalancersList(store *Store, rw http.ResponseWriter) {
	res, err := store.ListBalancers()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
