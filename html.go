package table

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
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

func FromHTML(r io.Reader) (*Table, error) {
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	table := &Table{}

	p := htmlParser{}

	if err := p.HTML(node, table); err != nil {
		return nil, err
	}

	return table, nil
}

type htmlParser struct {
}

func (p htmlParser) HTML(node *html.Node, table *Table) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}

	at := node.DataAtom.String()
	at = strings.ToLower(at)

	if at == "table" {
		return p.Table(node, table)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if err := p.HTML(c, table); err != nil {
			return err
		}
	}

	return nil
}

func (p htmlParser) Table(node *html.Node, table *Table) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}

	if table == nil {
		return fmt.Errorf("table is nil")
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		at := c.DataAtom.String()
		at = strings.ToLower(at)

		if at == "thead" {
			// if err := parseHTMLTableHead(c, table); err != nil {
			// 	return err
			// }
		}

		if at == "tbody" {
			// if err := parseHTMLTableBody(c, table); err != nil {
			// 	return err
			// }
		}
	}

	return nil
}
