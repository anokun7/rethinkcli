package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	r "gopkg.in/gorethink/gorethink.v3"
	"io/ioutil"
	"log"
	"os"
)

func GetFields(session *r.Session) []string {
	// Fetch the 1st row from the database
	res, err := r.DB(os.Args[2]).Table(os.Args[3]).Nth(0).Keys().Run(session)
	if err != nil {
		fmt.Print(err)
	}
	defer res.Close()

	if res.IsNil() {
		fmt.Print("Row not found")
	}

	var keys []string
	err = res.All(&keys)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}
	return keys
}

func GetTableContents(url string) {
	roots := x509.NewCertPool()
	cert, err := ioutil.ReadFile("/tmp/cert.pem")
	roots.AppendCertsFromPEM(cert)
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
		TLSConfig: &tls.Config{
			RootCAs: roots,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Fetch everything from the table
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

	fields := GetFields(session)
	var row map[string]interface{}
	i := 0
	for res.Next(&row) {
		if err != nil {
			fmt.Printf("Error scanning database result: %s", err)
			return
		}
		i += 1
		fmt.Printf("%2d:\n", i)
		for _, field := range fields {
			fmt.Printf("%15s: %v\n", field, row[field])
		}
	}
	fmt.Printf("\n\t\t====  Total rows returned: %d ====\n\n", i)
}

func main() {
	GetTableContents(os.Args[1])
}
