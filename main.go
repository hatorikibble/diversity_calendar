package main

import (
        "fmt"
	"github.com/hatorikibble/diversity_calendar/service"
        )
        
var appName = "diversity_calendar"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
