package main

import (
        "fmt"
        "net/http"
        "log"
)

func main() {
        http.HandleFunc("/", handler)
        port := "8080"
        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World!\n")
}
