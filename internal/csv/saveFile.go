package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	ghRepos "github.com/mstefa/github-cli-tool/internal"
)

type csvRepository struct {
}

func NewRepository() ghRepos.ReposCsv {
	return &csvRepository{}
}

func (c *csvRepository) SaveRepos(repos []ghRepos.GhRepo) {
	fmt.Println(repos)
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("values.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	csvwriter := csv.NewWriter(f)

	for _, v := range repos {
		var row []string
		row = append(row, strconv.FormatInt(v.RepoID, 10))
		row = append(row, v.Name)
		row = append(row, v.Description)
		row = append(row, strconv.FormatBool(v.Fork))
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()

	f.Close()
}
