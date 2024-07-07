package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var filePath string
var folderPath string

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&filePath, "file", "f", "", "path to the markdown file")
	generateCmd.Flags().StringVarP(&folderPath, "folder", "d", "", "path to the folder containing markdown files")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a new AI Tutor json file from a markdown file",
	Long:  "Generates a new AI Tutor json file from a markdown file. Generation can be done for a single file or for all files in a folder.\n For a single file, use the --file flag.\n For all files in a folder, use the --folder flag.",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
		// Check if file or folder path is provided
		if filePath == "" && folderPath == "" {
			fmt.Println("Please provide a file or folder path")
			fmt.Println(" File: aigen generate --file /folder/file.md\n Folder: aigen generate --folder /folder.")
			return
		}
		// Convert markdown file to json
		if filePath != "" {
			// Convert markdown file to json
			//pkg.ConvertMdToJSON(filePath)
      fmt.Println(filePath)
		}
		// Convert all markdown files in a folder to json
		if folderPath != "" {
			// Walk through the folder and get all markdown files
			var files []string
			err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
					files = append(files, path)
				}
				return nil
			})
			if err != nil {
				log.Fatal(err)
			}
			// Convert all markdown files to json in parallel
			wg := sync.WaitGroup{}
			for _, filePath := range files {
				wg.Add(1)
				go func(path string) {
					//pkg.ConvertMdToJSON(path)
          fmt.Println(path)
					wg.Done()

				}(filePath)
			}
			wg.Wait()
		}
	},
}
