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
	if r.Method != "GET" {
		http.Error(w, "Could not create " + r.Method + " request ", http.StatusBadRequest)
		return
	}
	dateString := vars["date"]
	var splitedDate = strings.Split(dateString, "-")
	if len(splitedDate) < 3 || len(splitedDate) > 3 {
		http.Error(w, "not a valid date", http.StatusBadRequest)
		return
	} else {
		day, _ := strconv.Atoi(splitedDate[2])
		month, _ := strconv.Atoi(splitedDate[1])
		year, _ := strconv.Atoi(splitedDate[0])
		date, err := bsdate.New(day, month, year)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		gregorianDate, _ := date.GetGregorianDate()
		fmt.Fprintf(w, gregorianDate.Format("2006-01-02"))
	}
}

func getBsFromAd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if r.Method != "GET" {
		http.Error(w, "Could not create " + r.Method + " request ", http.StatusBadRequest)
		return
	}
	dateString := vars["date"]
	var splitedDate = strings.Split(dateString, "-")
	if len(splitedDate) < 3 || len(splitedDate) > 3 {
		http.Error(w, "cannot convert date, invalid or missing data", http.StatusBadRequest)
		return
	}
	day, _ := strconv.Atoi(splitedDate[2])
	month, _ := strconv.Atoi(splitedDate[1])
	year, _ := strconv.Atoi(splitedDate[0])
	bsDate, err := bsdate.NewFromGregorian(day, month, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%d-%02d-%02d", bsDate.GetYear(), bsDate.GetMonth(), bsDate.GetDay())
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/ad-from-bs/{date}", getAdFromBs)
	myRouter.HandleFunc("/bs-from-ad/{date}", getBsFromAd)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
