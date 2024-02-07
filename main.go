package main

import (
	"flag"
	"fmt"
	"log"
)


func seedAccount(s Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Account : ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "anthony", "GG", "hunter88")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil{
		log.Fatal(err)
	}

	if err2 := store.Init(); err2 != nil {
		log.Fatal(err2)
	}

	if *seed {
		fmt.Println("seeding the DB")
		seedAccounts(store)
	}

	server := NewAPIServer("127.0.0.1:3000", store)
	server.Run()
}