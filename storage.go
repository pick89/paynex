package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	
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
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"
	db, err := sql.Open ("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil{
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil

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

func (s *PostgresStore) CreateAccount(*Account) error{
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
