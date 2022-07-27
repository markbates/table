package table

import (
	"fmt"
	"io"
	"os"
)

type HTML struct {
	DisableHeader bool
}

func (html HTML) Print(w io.Writer, table *Table) error {
	if w == nil {
		w = os.Stdout
	}

	if table == nil {
		return fmt.Errorf("table is nil")
	}

	fmt.Fprintln(w, "<table>")

	if !html.DisableHeader {
		if err := html.printColumns(w, table); err != nil {
			return err
		}
	}

	rows, err := table.Rows()
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "<tbody>")
	for _, row := range rows {

		fmt.Fprintln(w, "<tr>")

		for _, col := range row {
			fmt.Fprintf(w, "\t<td>%v</td>\n", col)
		}

		fmt.Fprintln(w, "</tr>")
	}

	fmt.Fprintln(w, "</tbody>")
	fmt.Fprintln(w, "</table>")

	return nil
}

func (html HTML) printColumns(w io.Writer, table *Table) error {
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

	if len(cols) == 0 {
		return nil
	}

	fmt.Fprintln(w, "<thead>")

	for _, col := range cols {

		fmt.Fprintf(w, "<th>%s</th>", col)
	}

	fmt.Fprintln(w, "<thead>")

	return nil
}
