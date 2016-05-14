package main

import (
  "fmt"
  "net/http"
  "log"

  // "golang.org/x/crypto/bcrypt"
)

// Index handler
func index(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    errorHandler(w, r, http.StatusNotFound)
    return
  }

  handle(w, r, "public/index.html")
}

// Login handler
func login(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET":
      if r.URL.Path != "/#/login" {
        errorHandler(w, r, http.StatusNotFound)
        return
      }

      http.Redirect(w, r, "/#/login", http.StatusSeeOther)
    case "POST":
      var form struct {
        Email    string
        Password string
      }
      if err := decodeReqJson(r, &form); err != nil {
        log.Println(err)
      }

      acct, err := getAccount(db, form.Email)
      if err != nil {
        log.Println(err)
      }

      // Account password and login password comparison
      if !equivPassword(acct.Password, form.Password) {
        fmt.Println("Password did not match!")
        return
      }

      fmt.Println("Login successful!")
  }
}

// Sign-up handler
func signup(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET":
      if r.URL.Path != "/#/signup" {
        errorHandler(w, r, http.StatusNotFound)
        return
      }

      http.Redirect(w, r, "/#/signup", http.StatusSeeOther)
    case "POST":
      form := &Account{}
      if err := decodeReqJson(r, form); err != nil {
        log.Println(err)
      }

      // Encrypt the password before saving it to the database
      hashedPassword, err := encryptPassword(form.Password)
      if err != nil {
        log.Println(err)
      }

      // Replace the password and date with hashed and formated versions
      form.Password = hashedPassword

      // Query the new Account into the database
      if err := addAccount(db, form); err != nil {
        log.Println(err)
      }
  }
}

// 404 handler
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
  w.WriteHeader(status)
  if status == http.StatusNotFound {
    handle(w, r, "public/templates/404.html")
  }
}