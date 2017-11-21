package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// PageVariables represent the variables those are sent to the page
type PageVariables struct {
	Date string
	Time string
}

func root(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	HomePageVars := PageVariables{
		Date: now.Format("02-01-2015"),
		Time: now.Format("15:04:04"),
	}

	t, err := template.ParseFiles("home.html")

	if err != nil {
		log.Print("Template parsing error: ", err)
	}
	err = t.Execute(w, HomePageVars)

	if err != nil {
		log.Print("Template excuting error: ", err)
	}
}

func main() {
	http.HandleFunc("/", root)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)

	fmt.Println("Using the awesome port:", port)
	openbrowser(strings.Join([]string{"http://localhost:", port}, ""))
	panic(http.Serve(listener, nil))

}
