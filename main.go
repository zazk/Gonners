package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ruling the World")
	fmt.Println("Reload Page. Request /")
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
