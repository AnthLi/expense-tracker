package main

import (
  "fmt"
  "net/http"
  // "log"
)

// Index handler
func index(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    notFound(w, r, http.StatusNotFound)
    return
  }

  handle(w, r, "public/index.html")
}

// Login handler
func login(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET":
      if r.URL.Path != "/#/login" {
        notFound(w, r, http.StatusNotFound)
        return
      }

      http.Redirect(w, r, "/#/login", http.StatusSeeOther)
    case "POST":
      r.ParseForm()

      acct, err := getAccount(db, r.Form["email"][0])
      if err != nil {
        httpError(w, fmt.Sprint("%q\n", err), http.StatusInternalServerError)
        return
      }

      // Account does not exist
      if acct.Email == "" {
        httpError(w, "That account does not exist!", http.StatusUnauthorized)
        return
      }

      // Account password and login password comparison
      if !equivPassword(acct.Password, r.Form["password"][0]) {
        httpError(w, "Incorrect Password!", http.StatusUnauthorized)
        return
      }
  }
}

// Sign-up handler
func signup(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET":
      if r.URL.Path != "/#/signup" {
        notFound(w, r, http.StatusNotFound)
        return
      }

      http.Redirect(w, r, "/#/signup", http.StatusSeeOther)
    case "POST":
      r.ParseForm()

      // Encrypt the password before saving it to the database
      hashedPassword, err := encryptPassword(r.Form["password"][0])
      if err != nil {
        httpError(w, fmt.Sprint("%q\n", err), http.StatusInternalServerError)
        return
      }

      acct, err := getAccount(db, r.Form["email"][0])
      if err != nil {
        httpError(w, fmt.Sprint("%q\n", err), http.StatusInternalServerError)
        return
      }

      if acct.Email != "" {
        httpError(w, "That email is already in use!", http.StatusUnauthorized)
        return
      }

      form := &Account{
        Fname: r.Form["fname"][0],
        Lname: r.Form["lname"][0],
        Email: r.Form["email"][0],
        Password: hashedPassword,
      }

      // Query the new Account into the database
      if err := addAccount(db, form); err != nil {
        httpError(w, fmt.Sprint("%q\n", err), http.StatusInternalServerError)
        return
      }
  }
}