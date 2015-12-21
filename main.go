package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

const rndBackendURLEnv string = "RND_BACKEND_URL"
const calcPiBackendURLEnv string = "CALC_PI_BACKEND_URL"
const serverPortEnv string = "FRONTEND_PORT"

const defaultServerPort string = "8080"

func main() {
	var envErrors []string

	rndBackendURL := os.Getenv(rndBackendURLEnv)
	calcPiBackendURL := os.Getenv(calcPiBackendURLEnv)
	servePort := os.Getenv(serverPortEnv)

	if rndBackendURL == "" {
		envErrors = append(envErrors, fmt.Sprintf("%s", rndBackendURLEnv))
	}

	if calcPiBackendURL == "" {
		envErrors = append(envErrors, fmt.Sprintf("%s", calcPiBackendURLEnv))
	}

	if servePort == "" {
		servePort = defaultServerPort
	}

	if len(envErrors) > 0 {
		log.Fatalf("cannot continue, we are missing the following enviornemnt variables: %v", strings.Join(envErrors, ", "))
	}

	bindPort := fmt.Sprintf(":%s", servePort)
	log.Printf("using RND_BACKEND_URL:%s\n", rndBackendURL)
	log.Printf("using CALC_PI_BACKEND_URL:%s\n", calcPiBackendURL)
	log.Printf("serving on forntend on :%s\n", bindPort)

	rndHdl := NewRndHandler(rndBackendURL)
	calcPiHdl := NewCaclPiHandler(calcPiBackendURL)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/rnd", rndHdl.Calc)
	r.HandleFunc("/pi", calcPiHdl.Calc)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(bindPort, r))

}
