package main

import (
  "html/template"
  "log"
  "net/http"
  "strings"
)

// Handle exceution of parsed templates
func render(w http.ResponseWriter, path string, data interface{}) {
  // Resolve AngularJS template conflict and render the template
  t := template.New(path).Delims("[[", "]]")
  t, _ = t.ParseFiles(path)
  t.ExecuteTemplate(w, path[strings.LastIndex(path, "/") + 1:], data)
}

// 404 handler
func notFound(w http.ResponseWriter, r *http.Request, status int) {
  w.WriteHeader(status)
  if status == 404 {
    render(w, "public/templates/404.html", nil)
  }
}

// Handle HTTP error replies to the request
func httpError(w http.ResponseWriter, err string, status int) {
  if status == 500 {
    log.Println(err)
  }

  http.Error(w, err, status)
}