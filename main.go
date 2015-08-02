package main

import "fmt"
import "flag"

func main() {
	repoSiteCode := flag.String("repoSiteCode","gitlab", "possible values: gitlab (github in the future)")
	destinationPath := flag.String("destination", ".", "clone to this location")

	fmt.Println("Hello")
	fmt.Printf("repoSiteCode:%s destinationPath:%s", repoSiteCode, destinationPath)

	//TODO: get the html
	//TODO: parse the repo links
	//TODO: pull the repos to destination

}