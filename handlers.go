package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RndHandler struct {
	BackendURL string
}

func NewRndHandler(backendURL string) *RndHandler {
	return &RndHandler{BackendURL: backendURL}
}

type Num struct {
	RndNum string
}

func (rnd *RndHandler) Calc(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/rnd", rnd.BackendURL)
	log.Println(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, "%s currently not availble", rnd.BackendURL)
		return
	}

	defer res.Body.Close()

	var myNum Num
	b, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(b, &myNum)

	fmt.Fprintf(w, "From Rnd calc page, secure random number using crypto/rand backend service :%s", myNum)
}

type CalcPiHandler struct {
	BackendURL string
}

func NewCaclPiHandler(backendURL string) *CalcPiHandler {
	return &CalcPiHandler{BackendURL: backendURL}
}

func (p *CalcPiHandler) Calc(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/pi", p.BackendURL)
	res, err := http.Get(url)
	log.Printf("%s", err)
	if err != nil {
		fmt.Fprintf(w, "%s currently not availble", p.BackendURL)
		return
	}

	fmt.Fprintf(w, "From Pi calc page, backendUrl:%s %s", p.BackendURL, res)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	g := "<p><a href=\"\\rnd\">Generate Secure Random Number</a></p>"
	// g := "<p><a href=\"\\rnd\">Generate Random Number</a></p> <p><a href=\"\\pi\">Calc Pi from Rnd Number</a></p>"
	fmt.Fprintf(w, g)
}
