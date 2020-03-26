package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"server"
	"net/http"
	"strconv"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var port *string = flag.String("p", "8080", "Port number (int>0)")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}
	if portnumb, err := strconv.Atoi(*port); err != nil || portnumb < 1 {
		fmt.Println("Invalid port number")
		return
	}
	commaport := fmt.Sprintf(":%s", *port)
	r := mux.NewRouter()
	r.Path("/locs").HandlerFunc(server.GetLocals).Methods("GET")
	r.Path("/locs").HandlerFunc(server.AddLocal).Methods("POST")
	server := &http.Server{Addr: commaport, Handler: r}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Unnexpected Error: %s", err)
	}

}

