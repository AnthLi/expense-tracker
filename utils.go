package main

import (
  "net/http"
  "log"
  "html/template"
)

// Handle exceution of parsed templates
func render(w http.ResponseWriter, path string, data interface{}) {
  // Resolve AngularJS template conflict and render the template
  t, _ := template.ParseFiles(path)
  t.Delims("[[", "]]").Execute(w, data)
}

// 404 handler
func notFound(w http.ResponseWriter, r *http.Request, status int) {
  w.WriteHeader(status)
  if status == http.StatusNotFound {
    render(w, "public/templates/404.html", nil)
  }
}

// Handle HTTP error replies to the request
func httpError(w http.ResponseWriter, err string, status int) {
  if status == http.StatusInternalServerError {
    log.Println(err)
  }

  http.Error(w, err, status)
}