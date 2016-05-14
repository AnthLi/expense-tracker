package main

import (
  "net/http"
  "encoding/json"
  "strings"
  "html/template"

  "golang.org/x/crypto/bcrypt"
)

// Decode the requests's body for POST values
func decodeReqJson(r *http.Request, i interface{}) error {
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(i)
  if err != nil {
    return err
  }

  return nil;
}

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
func handle(w http.ResponseWriter, r *http.Request, path string) {
  // Parse the file name from the relative path
  f := path
  fIndex := strings.LastIndex(path, "/")
  if fIndex > 0 {
    f = path[fIndex + 1:]
  }

  // Resolve AngularJS template conflict and render the template
  t := template.New(f).Delims("[[", "]]")
  t, _ = t.ParseFiles(path)
  t.Execute(w, nil)
}