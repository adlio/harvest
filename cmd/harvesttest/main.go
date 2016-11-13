package main

import (
	"flag"
	"fmt"

	"github.com/adlio/harvest"
)

func main() {
	var domain string
	var username string
	var password string

	flag.StringVar(&domain, "domain", "culturefoundry", "Harvest subdomain")
	flag.StringVar(&username, "username", "user@example.com", "Your username")
	flag.StringVar(&password, "password", "password", "Your password")
	flag.Parse()

	h := harvest.NewBasicAuthAPI(domain, username, password)
	project, err := h.GetProject(9292184, harvest.Defaults())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(project.Name)
}
