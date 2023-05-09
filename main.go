package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tushar2708/altcsv"
	"os"
	"path"
	"path/filepath"
)

const file string = "./metadata.db"

func main() {
	db, err := sql.Open("sqlite3", file)
	if nil != err {
		fmt.Println("Failed to to read " + file + " - " + err.Error())
		return
	}
	defer db.Close()

	var filePath, _ = filepath.Abs(file)
	var fileDir = path.Dir(filePath)
	fmt.Println(fileDir)

	var sql = "SELECT title, path FROM books ORDER BY title;"

	rows, err := db.Query(sql)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	csvFilePath := path.Join(fileDir, "catalog.csv")
	csvFile, _ := os.Create(csvFilePath)
	defer csvFile.Close()

	csv := altcsv.NewWriter(csvFile)
	csv.AllQuotes = true
	defer csvFile.Close()

	var titleCol = ""
	var pathCol = ""
	for rows.Next() {
		err = rows.Scan(&titleCol, &pathCol)
		if nil != err {
			continue
		}

		err := csv.Write([]string{titleCol, pathCol})
		if nil != err {
			continue
		}
	}
	csv.Flush()
}
