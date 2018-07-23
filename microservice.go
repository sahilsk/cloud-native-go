package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sahilsk/cloud-native-go/api"
)

/*
HelloFuncHandler

*/
func HelloFuncHandler(resp http.ResponseWriter, req *http.Request) {
	url := req.URL
	q := url.Query()

	output := fmt.Sprintf("Hello World %v %v ", q.Get("name"), q.Get("last"))
	resp.WriteHeader(http.StatusOK)
	_, err := resp.Write([]byte(output))
	if err == nil {
		fmt.Printf("Error: %v \n", err)
	}
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return port
}

func main() {

	fmt.Printf("API Version: %v \n", api.API_VERSION)

	//routes
	http.HandleFunc("/hello", HelloFuncHandler)

	http.HandleFunc("/books", api.ListBooks)
	http.HandleFunc("/books/", api.BookActions)

	fmt.Printf("listening on port: %v \n", port())
	addr := fmt.Sprintf("localhost:%s", port())

	http.ListenAndServe(addr, nil)

}
