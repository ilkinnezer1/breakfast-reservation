package handlers

import (
	"fmt"
	"net/http"
)
const port = ":9233"

func Start(){
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	fmt.Println(fmt.Sprintf("Server is running on port %s", port))
	http.ListenAndServe(port, nil)
}