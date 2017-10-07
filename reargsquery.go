package main

import (
	"fmt"
	r "gopkg.in/gorethink/gorethink.v3"
	"log"
	"os"
)

func GetTableContents(url string) {
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Fetch the row from the database
	res, err := r.DB(os.Args[2]).Table(os.Args[3]).Run(session)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer res.Close()

	if res.IsNil() {
		fmt.Print("Row not found")
		return
	}

	var row map[string]interface{}
	i := 0
	for res.Next(&row) {
		if err != nil {
			fmt.Printf("Error scanning database result: %s", err)
			return
		}
		i += 1
		fmt.Printf("%2d: %s\n", i, row["title"])
	}
}

func main() {
	GetTableContents(os.Args[1])
}
