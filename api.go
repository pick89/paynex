package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


// apiFunc defines the function signature for API handlers
type apiFunc func(http.ResponseWriter, *http.Request) error

// APIServer represents the API server
type APIServer struct {
	listenAddr 	string
	store 		Storage
}

// ErrClientError represents a client-side error
type ErrClientError struct {
    Msg string
}

// ApiError represents an error response
type ApiError struct {
	Error string `json:"error"`
}

// Error method makes ErrClientError satisfy the error interface.
func (e ErrClientError) Error() string {
    return e.Msg
}


// NewAPIServer creates a new instance of APIServer
func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}

// makeHTTPHandleFunc creates an HTTP handler function from an API function
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := f(w, r); err != nil {
            // Check if the error is a client error
            if clientErr, ok := err.(ErrClientError); ok {
                // If it's a client error, send a 400 Bad Request
                WriteJSON(w, http.StatusBadRequest, ApiError{Error: clientErr.Error()})
            } else {
                // If it's not a client error, assume it's a server error
                WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Internal Server Error"})
            }
        }
    }
}

// Run starts the API server
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Println("JSON API server running on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// WriteJSON writes JSON response with appropriate headers
func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json") // Set the content type header.
    w.WriteHeader(status) // Write the HTTP status code.
	if err := json.NewEncoder(w).Encode(v); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return err
    }
    return nil
}

// handleAccount handles requests for /account
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetAccounts(w, r)
	case http.MethodPost:
		return s.handleCreateAccount(w, r)
	case http.MethodDelete:
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("Method not allowed: %s", r.Method)
	}
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
    var createAccountReq CreateAccountRequest
    if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
        // Use WriteJSON to send an error message if decoding fails
        return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid JSON data"})
    }
    
    account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)
    if err := s.store.CreateAccount(account); err != nil {
        // Use WriteJSON to send an error message if creating the account fails
        return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to create account"})
    }
    
    // Successfully created the account, return it as JSON
    return WriteJSON(w, http.StatusCreated, account)
}

// handleGetAccount handles GET requests for /account/{id}
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	log.Println("Requested account ID:", id)
	// Implement logic to fetch account details from database using the id
	// For now, let's just return an empty account
	return WriteJSON(w, http.StatusOK, &Account{})
}

// handleGetAccounts handles GET requests for /account
func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfert(w http.ResponseWriter, r *http.Request) error {
	return nil
}
