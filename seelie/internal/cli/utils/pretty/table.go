package pretty

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

// PrintTable pretty print lines to ascii table style
func PrintTable(rows [][]interface{}) {
	t := table.NewWriter()
	t.AppendHeader(rows[0])
	for _, row := range rows[1:] {
		t.AppendRow(row)
	}
	fmt.Println(t.Render())
}
