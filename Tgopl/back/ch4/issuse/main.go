package main

import (
	"fmt"
	"log"
	"os"
	"tgopl/back/ch4/github"
)

func main() {
	result, err := github.SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issuse:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
