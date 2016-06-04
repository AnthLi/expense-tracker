package main

import (
  "fmt"
  "log"
  "net/http"
  "os"

  "gopkg.in/pg.v4"
  "github.com/icza/session"
)

var db *pg.DB

func main() {
  db = connect()

  // if err := deleteSchema(db); err != nil {
  //   log.Println(err)
  // }

  if err := createSchema(db); err != nil {
    log.Println(err)
  }

  // Initialize the session manager
  session.Global = session.NewCookieManagerOptions(
    session.NewInMemStore(),
    &session.CookieMngrOptions{AllowHTTP: true},
  )

  // Serve static files
  http.Handle("/public/",
    http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

  http.HandleFunc("/", index)
  http.HandleFunc("/login", login)
  http.HandleFunc("/signup", signup)
  http.HandleFunc("/logout", logout)
  http.HandleFunc("/search", search)
  http.HandleFunc("/add", add)
  http.HandleFunc("/accounts", accounts)
  http.HandleFunc("/expenses/recent", recentExpenses)

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  portStr := fmt.Sprint(":", port)

  log.Println("Listening on:", portStr)
  http.ListenAndServe(portStr, nil)
}