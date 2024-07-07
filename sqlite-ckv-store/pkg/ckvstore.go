package pkg

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const TableName string = "CKVTable"

var db *sql.DB

type CKVEntry struct {
	Category string
	Key string
	Value string
}

func InitDB(dropFile bool, filePath string) {
	if filePath == "" {
		filePath = "./ckvstore.sqlite"
	}
	if dropFile {
		os.Remove(filePath)
	}
	var err error
	db, err = sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}
}

func CreateTable() {
	sqlStmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
	category TEXT,
	key TEXT,
	value TEXT,
	PRIMARY KEY(category, key));
	`, TableName)

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func ReadAll(category string) []CKVEntry {
	sqlStmt := fmt.Sprintf(`SELECT * FROM %s WHERE collection = '%s' and key!=''`, TableName, category)
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var entries []CKVEntry
	for rows.Next() {
		var entry CKVEntry
		err = rows.Scan(&entry.Category, &entry.Key, &entry.Value)
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, entry)
	}
	return entries
}

func Read(category,key,value string) CKVEntry {
	sqlStmt := fmt.Sprintf(`SELECT * FROM %s WHERE category = ? AND key = ?;`,TableName)
	row := db.QueryRow(sqlStmt, collection, key)
	var entry CKVEntry
	err := row.Scan(&entry.Category, &entry.Key, &entry.Value)
	if err != nil {
		log.Fatal(err)
	}
	return entry
}

func Upsert(category, key, value string) {
  sqlStmt := fmt.Sprintf(`INSERT OR REPLACE INTO %s (category, key, value) VALUES (?, ?, ?)`, TableName)
	_, err := db.Exec(sqlStmt, category,key,value)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func Delete(category, key string) {
  sqlStmt := fmt.Sprintf(`DELETE FROM %s WHERE collection = ? AND key = ?;`, TableName)
	_, err := db.Exec(sqlStmt, collection, key)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func DeleteCategory(category string) {
	sqlStmt := fmt.Sprintf(`DELETE FROM %s WHERE category = ?;`,TableName)
	_, err := db.Exec(sqlStmt, category)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
