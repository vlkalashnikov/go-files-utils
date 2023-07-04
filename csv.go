package files

import (
	"encoding/csv"
	"os"
)

func CreateCsvFile(path string, records [][]string) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	return w.WriteAll(records)
}
