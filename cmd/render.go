package cmd

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func render(memos []Memo) {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Content", "Tags"})

	layout := time.RFC3339

	for _, memo := range memos {
		tt, _ := time.Parse(layout, memo.CreateTime)
		createTimeFormatted := tt.Format("02.01.2006 15:04")
		t.AppendRows([]table.Row{{createTimeFormatted, memo.Content, memo.Property.Tags}})

	}
	t.Render()
}
