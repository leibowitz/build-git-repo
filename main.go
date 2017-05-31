package main

import (
	"fmt"
	git "github.com/libgit2/git2go"
	"os"
)

func main() {
	repo, err := git.OpenRepository(".")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	remote, err := repo.LookupRemote("origin")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Hello %+v\n", remote)
}
