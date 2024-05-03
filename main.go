package main

import (

	"log"
)

// Here are the entry point to run my database and my server 

func main(){
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.init(); err != nil {
		log.Fatal(err)
	}

	 server := NewAPIServer(":3000", store)
	 server .Run()

}
