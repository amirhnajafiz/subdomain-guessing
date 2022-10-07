package main

import (
	"fmt"

	"github.com/amirhnajafiz/subdomain-guessing/pkg"
)

// Lookup
// manages 3 methods for lookup types.
type Lookup interface {
	// TypeA ns lookup.
	TypeA(fqdn, serverAddr string) ([]string, error)
	// TypeCNAME ns lookup.
	TypeCNAME(fqdn, serverAddr string) ([]string, error)
	// Navigate ns lookup.
	Navigate(fqdn, serverAddr string) []pkg.Result
}

func main() {
	fmt.Println("vim-go")
}
