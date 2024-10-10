package main

import (
	"fmt"
	"log"

	"github.com/barturba/ticket-tracker/internal/server"
	_ "github.com/lib/pq"
)

func main() {

	srv := server.NewServer()
	fmt.Println("server started on ", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the ticket-tracker has started\n")
}
