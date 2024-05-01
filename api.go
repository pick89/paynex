package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// WriteJSON writes JSON response with appropriate headers
func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// ApiError represents an error response
type ApiError struct {
	Error string `json:"error"`
}

// apiFunc defines the function signature for API handlers
type apiFunc func(http.ResponseWriter, *http.Request) error

// makeHTTPHandleFunc creates an HTTP handler function from an API function
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println("Error:", err) // Logging the error for debugging
			WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Internal Server Error"})
		}
	}
}

// APIServer represents the API server
type APIServer struct {
	listenAddr string
}

// NewAPIServer creates a new instance of APIServer
func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
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


// handleAccount handles requests for /account
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetAccount(w, r)
	case http.MethodPost:
		return s.handleCreateAccount(w, r)
	case http.MethodDelete:
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("Method not allowed: %s", r.Method)
	}
}

// handleGetAccount handles GET requests for /account/{id}
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	log.Println("Requested account ID:", id)
	// Implement logic to fetch account details from database using the id
	// For now, let's just return an empty account
	return WriteJSON(w, http.StatusOK, &Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	// Implement logic to create a new account
	// For now, let's return a dummy account
	account := NewAccount("John", "Doe")
	return WriteJSON(w, http.StatusCreated, account)
}


func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfert(w http.ResponseWriter, r *http.Request) error {
	return nil
}