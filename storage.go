package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


//Interface for defining the method CRUD
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int)error
	UpdateAccount(*Account)error
	GetAccountByID(int)(*Account,error)
}

// This structure holds a pointer to excecte SQL Command
type PostgresStore struct{
	db *sql.DB
}

// This function for creating a connection to the database

func NewPostgresStore() (*PostgresStore, error) {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    // Get environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbSSLMode := os.Getenv("DB_SSLMODE")

    // Create connection string
    connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
        dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)

    // Connect to the database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) init() error{
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
    query := `
    CREATE TABLE IF NOT EXISTS accounts (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) NOT NULL UNIQUE,
        email VARCHAR(50) NOT NULL UNIQUE,
		number serial,
		balance serial,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    )`
    _, err := s.db.Exec(query)
    if err != nil {
        return err
    }
    return nil
}

func (s *PostgresStore) CreateAccount(acc *Account) error{
	query := `
	insert into account 
	(first_name, last_name, number, balance , created_at)
	values($1,$2,$3,$4,$5)`

	resp, err := s.db.Query(
		query, 
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance, 
		acc.CreateAt)
		if err != nil{	
			return err
		}
		fmt.Printf("%+v\n", resp)
		return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error{
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error{
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account,error){
	return nil, nil
}

//Function CRUD on my database
