package main

import (
  "fmt"
  "log"
  "net/http"
  "os"

  "gopkg.in/pg.v4"
)

func main() {
  port := os.Getenv("PORT")

  if port == "" {
    port = "8080"
  }

  portStr := fmt.Sprint(":", port)

  log.Println("Listening on:", portStr)

  // Serve public files
  http.Handle("/public/",
    http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
  http.Handle("/", http.FileServer(http.Dir("./templates/")))
  http.ListenAndServe(portStr, nil)
}