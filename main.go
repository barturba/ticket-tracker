package main

import (
	"fmt"
	"log"

	"github.com/barturba/ticket-tracker/internal/server"
	_ "github.com/lib/pq"
)

func main() {

	srv := server.NewServer()
	err := srv.ListenAndServe()
	fmt.Println("server started on ", srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the ticket-tracker has started\n")
}
