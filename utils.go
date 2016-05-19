package main

import (
  "net/http"
  "log"
  "strings"
  "html/template"

  "golang.org/x/crypto/bcrypt"
)

// Encrypt passwords using the bcrypt library
func encryptPassword(password string) (string, error) {
  b := []byte(password)
  hash, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
  if err != nil {
    return "", err
  }

  return string(hash), nil
}

// Compare the password with the hash
func equivPassword(hashedPassword, password string) bool {
  hBytes := []byte(hashedPassword)
  pBytes := []byte(password)
  err := bcrypt.CompareHashAndPassword(hBytes, pBytes)

  return err == nil
}

// Format the date as yyyy-mm-dd
func formatDate(date string) string {
  return date[:strings.Index(date, "T")]
}

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