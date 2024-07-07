package database

import (
	"backend/pkg/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const TableName string = "MemoryTable"

type Metadata struct {
	Chunk_ID    string
	Parent_ID   string
	Description string
	Chunk       string
	Other       string
}

type DatabaseEntry struct {
	Collection string
	Key        string
	Metadata   string
	Embedding  string
	Timestamp  string
	Vector     *[]float64
}

// type MemoryRecord struct{}
type Embeddings struct {
	MemoryRecord DatabaseEntry
	Relevance    float64
}

func UnparseMetadata(data Metadata) string {
	// serialize to json
	metadata, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return string(metadata)
}

func ParseMetadata(metadata string) Metadata {
	var data Metadata
	err := json.Unmarshal([]byte(metadata), &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func InitDB(dropFile bool, filePath string) {
	if filePath == "" {
		filePath = "./vector.sqlite"
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

func CloseDB() {
	db.Close()
}

func CreateMemoryTable() {
	sqlStmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
	collection TEXT,
	key TEXT,
	metadata TEXT,
	embedding TEXT,
	timestamp TEXT,
	PRIMARY KEY(collection, key));
	`, TableName)

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func TableExists() bool {
	sqlStmt := fmt.Sprintf(`SELECT name FROM sqlite_master WHERE type='table' AND name='%s'`, TableName)
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return rows.Next()
}

func DropTable() {
	sqlStmt := `DROP TABLE ` + TableName + `;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func TruncateTable() {
	sqlStmt := `DELETE FROM ` + TableName + `;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func CreateCollection(collection string) {
	if CollectionExists(collection) {
		return
	}

	sqlStmt := fmt.Sprintf(`INSERT INTO %s (collection) VALUES (?)`, TableName)
	_, err := db.Exec(sqlStmt, collection)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func Upsert(collection, key, metadata, embedding string) {
	sqlStmt := `INSERT OR REPLACE INTO ` + TableName + ` (collection, key, metadata, embedding, timestamp) VALUES (?, ?, ?, ?, ?)`
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec(sqlStmt, collection, key, metadata, embedding, timestamp)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func GetCollections() []string {
	sqlStmt := `SELECT DISTINCT collection FROM ` + TableName + `;`
	rows, err := db.Query(sqlStmt)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var collections []string
	for rows.Next() {
		var collection string
		err = rows.Scan(&collection)
		if err != nil {
			log.Fatal(err)
		}
		collections = append(collections, collection)
	}
	return collections
}

func CollectionExists(collection string) bool {
	sqlStmt := `SELECT collection FROM ` + TableName + ` WHERE collection = ?;`
	rows, err := db.Query(sqlStmt, collection)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return rows.Next()
}

func DropColllection(collection string) {
	sqlStmt := `DELETE FROM ` + TableName + ` WHERE collection = ?;`
	_, err := db.Exec(sqlStmt, collection)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}


func ReadAll(collection string) []DatabaseEntry {
	sqlStmt := fmt.Sprintf(`SELECT * FROM %s WHERE collection = '%s' and key!=''`, TableName, collection)
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var entries []DatabaseEntry
	for rows.Next() {
		var entry DatabaseEntry
		err = rows.Scan(&entry.Collection, &entry.Key, &entry.Metadata, &entry.Embedding, &entry.Timestamp)
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, entry)
	}
	return entries
}

func Read(collection, key string) DatabaseEntry {
	sqlStmt := `SELECT * FROM ` + TableName + ` WHERE collection = ? AND key = ?;`
	row := db.QueryRow(sqlStmt, collection, key)
	var entry DatabaseEntry
	err := row.Scan(&entry.Collection, &entry.Key, &entry.Metadata, &entry.Embedding, &entry.Timestamp)
	if err != nil {
		log.Fatal(err)
	}
	return entry
}

func Delete(collection, key string) {
	sqlStmt := `DELETE FROM ` + TableName + ` WHERE collection = ? AND key = ?;`
	_, err := db.Exec(sqlStmt, collection, key)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func DeleteCollection(collection string) {
	sqlStmt := `DELETE FROM ` + TableName + ` WHERE collection = ?;`
	_, err := db.Exec(sqlStmt, collection)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func UnparseEmbedding(vector []float64) string {
	// serialize to json
	embedding, err := json.Marshal(vector)
	if err != nil {
		log.Fatal(err)
	}
	return string(embedding)
}

func ParseEmbedding(embedding string) []float64 {
	// embedding is a json array of floats
	var vector []float64
	// deserialize json
	err := json.Unmarshal([]byte(embedding), &vector)
	if err != nil {
		log.Fatal(err)
	}
	return vector
}

func NearSearch(collection, embedding string, limit int, minRelevance float64, withEmbedding bool) []DatabaseEntry {

	// Read all records in the collection
	entries := ReadAll(collection)

	// Find the relevance of the embedding with each record
	var list []Embeddings
	for _, entry := range entries {
		if entry.Embedding == "" {
			continue
		}
		embedding1 := ParseEmbedding(entry.Embedding)
		embedding2 := ParseEmbedding(embedding)
		relevance := util.CosineSimilarity(embedding1, embedding2)
		if relevance >= minRelevance {
			if withEmbedding {
				entry.Vector = &embedding1
			}
			list = append(list, Embeddings{entry, relevance})
		}
	}

	// Sort the list by relevance
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i].Relevance < list[j].Relevance {
				list[i], list[j] = list[j], list[i]
			}
		}
	}

	// Limit the list
	var final []DatabaseEntry
	count := 0
	for _, e := range list {
		if withEmbedding {
			final = append(final, e.MemoryRecord)
		}
		count++
		if count >= limit {
			break
		}
	}
	return final
}
