package main

import (
	"database/sql"
	"fmt"
)

func InitDB(dbPath string) (err error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	createStmt := fmt.Sprintf(`
		create table seq (
			id integer primary key autoincrement,
			alignment_name text,
			sample_name text,
			description text,
			seq text,
			unique(alignment_name, sample_name) on conflict fail);`)
	_, err = db.Exec(createStmt)
	if err != nil {
		return err
	}
	return
}

func RecordFasta(dbPtr *sql.DB, sequences []sequence) (cnt int, err error) {
	sqlStmt := "insert into seq(alignment_name, sample_name, description, seq) values(?, ?, ?, ?)"

	// Create transaction per replicate
	tx, err := dbPtr.Begin()
	if err != nil {
		return cnt, err
	}

	stmt, err := tx.Prepare(sqlStmt)
	if err != nil {
		return cnt, err
	}

	for _, s := range sequences {
		_, err := stmt.Exec(s.alignmentName, s.id, s.description, s.seq)
		if err != nil {
			return cnt, err
		}
		cnt++
	}
	return cnt, err
}
