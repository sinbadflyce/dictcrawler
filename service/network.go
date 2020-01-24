package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Network ...
type Network struct {
}

// Listen ...
func (n *Network) Listen() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:3000/graphql?query={Word(Name:\"Hello!\"){Name,Entries{Homnum,Topics,Senses{SignPost,Definition}}}}'")
	http.ListenAndServe(":3000", nil)
}
