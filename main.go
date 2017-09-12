package main

import (
	"database/sql"
	"flag"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbPathPtr := flag.String("db", "", "Path of database file")
	flag.Parse()
	fastaPath := flag.Arg(0)

	// Check if fasta path exists
	if _, err := os.Stat(fastaPath); os.IsNotExist(err) {
		panic(err)
	}

	// Check if folder to db path exists
	dbDirPath := path.Dir(*dbPathPtr)
	if _, err := os.Stat(dbDirPath); os.IsNotExist(err) {
		panic(err)
	}

	fastaBasePath := path.Base(fastaPath)
	sequences := ReadFasta(fastaPath, fastaBasePath)

	// Open db
	db, err := sql.Open("sqlite3", *dbPathPtr)
	if err != nil {
		panic(err)
	}
	_, err = RecordFasta(db, sequences)
}
