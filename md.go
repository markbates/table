package table

import (
	"fmt"
	"io"
	"os"
)

type Markdown struct {
	DisableHeader     bool
	DisableHeaderLine bool
}

func (md Markdown) Print(w io.Writer, table *Table) error {
	if w == nil {
		w = os.Stdout
	}

	if table == nil {
		return fmt.Errorf("table is nil")
	}

	cols, err := table.Columns()
	if err != nil {
		return err
	}

	if len(cols) > 0 && !md.DisableHeader {
		if err := md.printColumns(w, table); err != nil {
			return err
		}
	}

	rows, err := table.Rows()
	if err != nil {
		return err
	}

	ml, err := table.MaxColumnWidth()
	if err != nil {
		return err
	}

	if ml < 3 {
		ml = 3
	}

	for _, row := range rows {
		fmt.Fprint(w, "\n|")

		for _, col := range row {
			s := fmt.Sprintf("%v", col)

			for len(s) < ml {
				s += " "
			}

			fmt.Fprintf(w, " %s |", s)
		}

	}

	return nil
}

func (md Markdown) printColumns(w io.Writer, table *Table) error {
	if w == nil {
		w = os.Stdout
	}

	if table == nil {
		return fmt.Errorf("table is nil")
	}

	ml, err := table.MaxColumnWidth()
	if err != nil {
		return err
	}

	if ml < 3 {
		ml = 3
	}

	cols, err := table.Columns()
	if err != nil {
		return err
	}

	fmt.Fprint(w, "|")

	for _, col := range cols {
		for len(col) < ml {
			col += " "
		}

		fmt.Fprintf(w, " %s |", col)
	}

	if md.DisableHeaderLine {
		return nil
	}

	fmt.Fprint(w, "\n|")

	for range cols {
		var s string

		for len(s) < ml {
			s += "-"
		}

		fmt.Fprintf(w, " %s |", s)
	}

	return nil
}
