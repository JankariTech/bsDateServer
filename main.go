package main

import (
	"fmt"
	"github.com/JankariTech/GoBikramSambat"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bikram Sambat Server")
}

func getAdFromBs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dateString := vars["date"]
	var splitedDate = strings.Split(dateString, "-")
	day, _ := strconv.Atoi(splitedDate[2])
	month, _ := strconv.Atoi(splitedDate[1])
	year, _ := strconv.Atoi(splitedDate[0])
	date, _ := bsdate.New(day, month, year)
	gregorianDate, _ := date.GetGregorianDate()
	fmt.Fprintf(w, gregorianDate.Format("2006-01-02"))
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/ad-from-bs/{date}", getAdFromBs)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
