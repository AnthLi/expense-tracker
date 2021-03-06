package main

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "log"
  "strings"

  "gopkg.in/pg.v4"
)

type Credentials struct {
  Host string     `json:"host"`
  User string     `json:"user"`
  Database string `json:"database"`
  Password string `json:"password"`
}

type Account struct {
  Aid      int
  Fname    string `json:"fname"`
  Lname    string `json:"lname"`
  Email    string `json:"email"`
  Password string
}

type Expense struct {
  Eid    int
  Aid    int
  Name   string `json:"name"`
  Amount string `json:"amount"`
  Date   string `json:"date"`
}

type Search struct {
  Email string `json:"email"`
  Name  string `json:"amount"`
  Date  string `json:"date"`
}

// Establish a connection to the database
func connect() *pg.DB {
  // Get the database credentials
  file, err := ioutil.ReadFile("db-credentials.json")
  if err != nil {
    log.Println(err)
  }

  var cred Credentials
  json.Unmarshal(file, &cred)
  db := pg.Connect(&pg.Options{
    Addr: cred.Host,
    User: cred.User,
    Password: cred.Password,
  })

  return db
}

// Testing purposes only
func deleteSchema(db *pg.DB) error {
  queries := []string {
    `DROP TABLE IF EXISTS Account CASCADE`,
    `DROP TABLE IF EXISTS Expense CASCADE`,
    `DROP TABLE IF EXISTS AccountExpense`,
  }

  for _, q := range queries {
    _, err := db.Exec(q)
    if err != nil {
      return err
    }
  }

  return nil
}

// Initialize the Schema
func createSchema(db *pg.DB) error {
  queries := []string {
    `CREATE TABLE IF NOT EXISTS Account(
      aid      SERIAL PRIMARY KEY,
      fname    TEXT NOT NULL,
      lname    TEXT NOT NULL,
      email    VARCHAR(1000) UNIQUE NOT NULL,
      password TEXT NOT NULL
    )`,
    `CREATE TABLE IF NOT EXISTS Expense(
      eid    SERIAL PRIMARY KEY,
      aid    INTEGER REFERENCES Account(aid),
      name   TEXT NOT NULL,
      amount MONEY NOT NULL,
      date   DATE NOT NULL
    )`,
  }

  for _, q := range queries {
    _, err := db.Exec(q)
    if err != nil {
      return err
    }
  }

  return nil
}

// Account sign up
func addAccount(db *pg.DB, acct *Account) error {
  q := `INSERT INTO Account (fname, lname, email, password)
    VALUES (?fname, ?lname, ?email, ?password)`

  _, err := db.Exec(q, acct)
  if err != nil {
    return err
  }

  return nil
}

// Search a account based on the email
func getAccount(db *pg.DB, email string) (Account, error) {
  var account Account
  q := `SELECT * FROM Account WHERE email = ?`
  _, err := db.Query(&account, q, email)

  return account, err
}

// Get the user's name based on their email
func getAccountName(db *pg.DB, email string) (Account, error) {
  var account Account
  q := `SELECT fname, lname FROM Account WHERE email = ?`
  _, err := db.Query(&account, q, email)

  return account, err
}

// Retrieve all accounts
func allAccounts(db *pg.DB) ([]Account, error) {
  var accounts []Account
  _, err := db.Query(&accounts, `SELECT * FROM Account`)

  return accounts, err
}

// Add an expense to a specific account
func addExpense(db *pg.DB, expense *Expense) error {
  q := `INSERT INTO Expense (aid, name, amount, date)
    VALUES (?aid, ?name, ?amount, ?date)`
  _, err := db.Exec(q, expense)
  if err != nil {
    return err
  }

  return nil
}

// Get the most recent (last 10) expenses
func getRecentExpenses(db *pg.DB, email string) ([]Expense, error) {
  var expenses []Expense
  q := `SELECT e.name, e.amount, e.date
    FROM Expense e, Account a
    WHERE e.aid = a.aid AND a.email = ?
    ORDER BY e.date DESC
    LIMIT 10`
  _, err := db.Query(&expenses, q, email)

  return expenses, err
}

func searchExpenses(db *pg.DB,
                    email string,
                    search *Search) ([]Expense, error) {
  var expenses []Expense
  q := `SELECT e.name, e.amount, e.date
    FROM Expense e, Account a
    WHERE e.aid = a.aid AND a.email = ?email`

  switch {
    case search.Name != "" && search.Date == "null":
      q += ` AND a.name = ?name`
      _, err := db.Query(&expenses, q, search)
      if err != nil {
        return nil, err
      }
    case search.Name == "" && search.Date != "null":
      cutoff := strings.Index(search.Date, "00:00:00")
      search.Date = search.Date[:cutoff - 1]

      q += ` AND e.date = ?date`

    default:
      return nil, errors.New("Please fill out at least one field")
  }

  return expenses, nil
}

// Retrieve all expenses
func allExpenses(db *pg.DB) ([]Expense, error) {
  var expenses []Expense
  _, err := db.Query(&expenses, `SELECT * FROM Expense`)

  return expenses, err
}