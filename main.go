package main

import (
	"fmt"
	"log"
	"os"

	"github.com/drewstinnett/flipip/internal/dns"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("Usage: %v ip1 [ip2...]", os.Args[0])
	}
	for _, item := range os.Args[1:] {
		rev, err := dns.Reverse(item)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rev)
	}
}
