package main

import (
  "fmt"
  "net/http"
  "strings"

  "github.com/icza/session"
)

// Global so it doesn't get reinitialized after refresh
var sess session.Session
var sessionCount = 0

// Index handler
func index(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    notFound(w, r, http.StatusNotFound)
    return
  }

  render(w, "public/index.html", nil)
}

// Login handler
func login(w http.ResponseWriter, r *http.Request) {
  req := r.Method
  if req == "GET" {
    if r.URL.Path != "/login" {
      notFound(w, r, http.StatusNotFound)
      return
    }

    http.Redirect(w, r, "/#/login", http.StatusSeeOther)
  } else if req == "POST" {
    acct, err := getAccount(db, strings.ToLower(r.FormValue("email")))
    if err != nil {
      httpError(w, fmt.Sprint("%q\n", err), http.StatusInternalServerError)
      return
    }

    // Account does not exist
    if acct.Email == "" {
      httpError(w, "That account does not exist!", http.StatusUnauthorized)
      return
    }

    // Account password does not match login password
    if !equivPassword(acct.Password, r.FormValue("password")) {
      httpError(w, "Incorrect Password!", http.StatusUnauthorized)
      return
    }

    // Successful login
    sessionCount += 1
    sess = session.NewSessionOptions(&session.SessOptions {
      CAttrs: map[string]interface{}{"UserName": acct.Email},
      Attrs:  map[string]interface{}{"Count": sessionCount},
    })

    session.Add(sess, w)
  }
}

// Logout handler
func logout(w http.ResponseWriter, r *http.Request) {
  if sess != nil {
    session.Remove(sess, w)
    sess = nil
  }
}

// Sign-up handler
func signup(w http.ResponseWriter, r *http.Request) {
  req := r.Method
  if req == "GET" {
    if r.URL.Path != "/signup" {
      notFound(w, r, http.StatusNotFound)
      return
    }

    http.Redirect(w, r, "/#/signup", http.StatusSeeOther)
  } else if req == "POST" {
    // Encrypt the password before saving it to the database
    hashedPassword, err := encryptPassword(r.FormValue("password"))
    if err != nil {
      httpError(w, fmt.Sprint("%q\n", err), http.StatusInternalServerError)
      return
    }

    acct, err := getAccount(db, strings.ToLower(r.FormValue("email")))
    if err != nil {
      httpError(w, fmt.Sprint("%q\n", err), http.StatusInternalServerError)
      return
    }

    if acct.Email != "" {
      httpError(w, "That email is already in use!", http.StatusUnauthorized)
      return
    }

    form := &Account{
      Fname: r.FormValue("fname"),
      Lname: r.FormValue("lname"),
      Email: strings.ToLower(r.FormValue("email")),
      Password: hashedPassword,
    }

    // Query the new Account into the database
    if err := addAccount(db, form); err != nil {
      httpError(w, fmt.Sprint("%q\n", err), http.StatusInternalServerError)
      return
    }
  }
}