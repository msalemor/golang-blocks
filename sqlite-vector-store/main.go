package main

import (
	"backend/pkg/database"
	"fmt"
)

func main() {
	database.InitDB(true, "")
	database.CreateMemoryTable()
	if database.TableExists() {
		//database.DropTable()
		//database.CreateMemoryTable()
		database.CreateCollection("collection")
		database.Upsert("collection", "key1",
			database.UnparseMetadata(database.Metadata{
				Chunk_ID:    "chunk_key1",
				Parent_ID:   "parent_chunk_key1",
				Description: "Description",
				Chunk:       "This is very long text.",
				Other:       "Other data"}),
			database.UnparseEmbedding([]float64{1, 2, 3, 4, 5}))
		database.Upsert("collection", "key2", "metadata", database.UnparseEmbedding([]float64{1, 2, 3, 4, 5}))
		database.Upsert("collection", "key3", "metadata", database.UnparseEmbedding([]float64{1, 2, 3, 4, 5}))
		rows := database.ReadAll("collection")
		for _, row := range rows {
			fmt.Println(row)
		}
		database.Delete("collection", "key1")
		rows = database.ReadAll("collection")
		for _, row := range rows {
			fmt.Println(row)
		}
		database.DeleteCollection("collection")
		rows = database.ReadAll("collection")
		if len(rows) == 0 {
			fmt.Println("Collection deleted")
		}
	}
}
