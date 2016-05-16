package main

import (
  "fmt"
  "net/http"
  // "log"

  "github.com/icza/session"
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
      if r.URL.Path == "/login" {
        http.Redirect(w, r, "/#/login", http.StatusSeeOther)
        return
      }

      if r.URL.Path != "/#/login" {
        notFound(w, r, http.StatusNotFound)
        return
      }
    case "POST":
      r.ParseForm()

      sess := session.Get(r)
      if sess != nil {
        fmt.Println(sess)
        return
      }

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

      // Account password does not match login password
      if !equivPassword(acct.Password, r.Form["password"][0]) {
        httpError(w, "Incorrect Password!", http.StatusUnauthorized)
        return
      }

      // Successful login, create and add a new session
      sess = session.NewSessionOptions(&session.SessOptions {
        CAttrs: map[string]interface{}{"UserName": acct.Email},
        Attrs:  map[string]interface{}{"Count": 1},
      })

      session.Add(sess, w)

      fmt.Println("New session:", session.Get(r))
  }
}

// Logout handler
func logout(w http.ResponseWriter, r *http.Request) {

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