package utils

import (
	"bytes"
	"encoding/csv"
)

func GenerateCsv(rows [][]string) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	
	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}

	// flush dulu SEBELUM ambil buf.Bytes()
	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}