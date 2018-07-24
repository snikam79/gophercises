package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	errorHelper(db.Init(dbPath))
	cmd.Execute()
}

func errorHelper(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
