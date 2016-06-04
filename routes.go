package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "net/url"
  "strings"

  "golang.org/x/crypto/bcrypt"
  "github.com/icza/session"
)

// Session is global so it doesn't get reinitialized after page refresh
var sess session.Session
var sessionCount = 0

// Index handler
func index(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    notFound(w, r, 404)
    return
  }

  render(w, "public/index.html", nil)
}

// Login handler
func login(w http.ResponseWriter, r *http.Request) {
  req := r.Method
  if req == "GET" {
    if r.URL.Path != "/login" {
      notFound(w, r, 404)
      return
    }

    render(w, "public/index.html", nil)
  } else if req == "POST" {
    if r.FormValue("email") == "" {
      httpError(w, "Please enter an email to log in", 401)
      return
    }

    acct, err := getAccount(db, strings.ToLower(r.FormValue("email")))
    if err != nil {
      httpError(w, err.Error(), 500)
      return
    }

    // Account does not exist
    if acct.Email == "" {
      httpError(w, "That account does not exist", 401)
      return
    }

    // Hashed password and login form password comparison
    hBytes := []byte(acct.Password)
    pBytes := []byte(r.FormValue("password"))
    err = bcrypt.CompareHashAndPassword(hBytes, pBytes)
    if err != nil {
      httpError(w, "The password you've entered is incorrect", 401)
      return
    }

    // Successful login
    sessionCount += 1
    sess = session.NewSessionOptions(&session.SessOptions {
      CAttrs: map[string]interface{}{"id": acct.Aid, "email": acct.Email},
      Attrs:  map[string]interface{}{"count": sessionCount},
    });

    session.Add(sess, w)
  }
}

// Sign-up handler
func signup(w http.ResponseWriter, r *http.Request) {
  req := r.Method
  if req == "GET" {
    if r.URL.Path != "/signup" {
      notFound(w, r, 404)
      return
    }

    render(w, "public/index.html", nil)
  } else if req == "POST" {
    // Parse the sign up form values to make sure the user filled out all
    // required fields
    r.ParseForm()
    for _, v := range r.Form {
      if v[0] == "" {
        httpError(w, "Please fill out the required fields", 401)
        return
      }
    }

    // Encrypt the password before saving it to the database
    b := []byte(r.FormValue("password"))
    hashedPassword, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
    if err != nil {
      httpError(w, err.Error(), 500)
      return
    }

    acct, err := getAccount(db, strings.ToLower(r.FormValue("email")))
    if err != nil {
      httpError(w, err.Error(), 500)
      return
    }

    if acct.Email != "" {
      httpError(w, "The email you've entered is already in use", 401)
      return
    }

    form := &Account{
      Fname: r.FormValue("fname"),
      Lname: r.FormValue("lname"),
      Email: strings.ToLower(r.FormValue("email")),
      Password: string(hashedPassword),
    }

    // Query the new Account into the database
    if err := addAccount(db, form); err != nil {
      httpError(w, err.Error(), 500)
      return
    }
  }
}

// Logout handler
func logout(w http.ResponseWriter, r *http.Request) {
  if sess != nil {
    session.Remove(sess, w)
    sess = nil
  }
}

// Search expense handler
func search(w http.ResponseWriter, r *http.Request) {
  if sess == nil {
    http.Redirect(w, r, "/login", 301)
    return
  }

  req := r.Method
  if req == "GET" {
    if r.URL.Path != "/search" {
      notFound(w, r, 404)
      return
    }

    render(w, "public/index.html", nil)
  } else if req == "POST" {
    form := &Search {
      Email: r.FormValue("email"),
      Name: r.FormValue("name"),
      Date: r.FormValue("date"),
    }

    expenses, err := searchExpenses(db, r.FormValue("email"), form)
    if err != nil {
      httpError(w, err.Error(), 403)
      return
    }

    fmt.Println("Expenses:", expenses)
  }
}

// Add expense handler
func add(w http.ResponseWriter, r *http.Request) {
  if sess == nil {
    http.Redirect(w, r, "/login", 301)
    return
  }

  req := r.Method
  if req == "GET" {
    if r.URL.Path != "/add" {
      notFound(w, r, 404)
      return
    }

    render(w, "public/index.html", nil)
  } else if req == "POST" {
    // Parse the date from the form, excluding time
    cutoff := strings.Index(r.FormValue("date"), "00:00:00")
    form := &Expense {
      Aid: sess.CAttr("id").(int),
      Name: r.FormValue("name"),
      Amount: r.FormValue("amount"),
      Date: r.FormValue("date")[:cutoff - 1],
    }

    if err := addExpense(db, form); err != nil {
      httpError(w, err.Error(), 500)
      return
    }
  }
}

// REST API for AngularJS to fetch data

// Gets the current account's name and email and passes it along as JSON
func accounts(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/accounts" {
    notFound(w, r, 404)
    return
  }

  query, _ := url.ParseQuery(r.URL.RawQuery)
  acct, err := getAccountName(db, strings.ToLower(query["email"][0]))
  if err != nil {
    httpError(w, err.Error(), 500)
    return
  }

  json.NewEncoder(w).Encode(acct)
}

// Gets the current account's recent expenses
func recentExpenses(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/expenses/recent" {
    notFound(w, r, 404)
    return
  }

  query, _ := url.ParseQuery(r.URL.RawQuery)
  expenses, err := getRecentExpenses(db, strings.ToLower(query["email"][0]))
  if err != nil {
    httpError(w, err.Error(), 500)
    return
  }

  json.NewEncoder(w).Encode(expenses)
}