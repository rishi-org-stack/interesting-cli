package main

import (
	"fmt"
	qu "github.com/rishi-org-stack/cli/events/query"
)

func main() {
	var s string
	// for {
		fmt.Scan(&s)
		switch s {
		case "events":
			qu.Solverquery()
		}
	
}
