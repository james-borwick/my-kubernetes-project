package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
)

func main() {
    http.HandleFunc("/", handler)
    port := "8080"
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    fmt.Fprintf(w, "Hostname: %s\nDeployment: violet\n", hostname)
}
