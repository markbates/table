package table

import (
	"fmt"
	"sync"
)

type Table struct {
	columns []string
	rows    Rows
	mu      sync.RWMutex
}

// MaxColumnWidth returns the size of the widest column.
func (table *Table) MaxColumnWidth() (int, error) {
	if table == nil {
		return 0, fmt.Errorf("table is nil")
	}

	table.mu.RLock()
	defer table.mu.RUnlock()

	var ml int

	for _, col := range table.columns {
		if len(col) > ml {
			ml = len(col)
		}
	}

	for _, row := range table.rows {
		for _, col := range row {
			x := fmt.Sprintf("%v", col)
			if len(x) > ml {
				ml = len(x)
			}
		}
	}

	return ml, nil
}

func (table *Table) Columns() ([]string, error) {
	if table == nil {
		return nil, fmt.Errorf("table is nil")
	}

	table.mu.RLock()
	defer table.mu.RUnlock()

	cols := make([]string, len(table.columns))
	copy(cols, table.columns)

	return cols, nil
}

func (table *Table) Rows() (Rows, error) {
	if table == nil {
		return nil, fmt.Errorf("table is nil")
	}

	table.mu.Lock()
	defer table.mu.Unlock()

	rows := make(Rows, len(table.rows))
	copy(rows, table.rows)

	return rows, nil
}

// SetColumns sets the names of the columns.
// Once set, the columns cannot be changed.
// Once set, rows will be constrained to the
// number of columns.
func (table *Table) SetColumns(s ...string) error {
	if table == nil {
		return fmt.Errorf("table is nil")
	}

	table.mu.Lock()
	defer table.mu.Unlock()

	if len(table.columns) > 0 {
		return fmt.Errorf("columns already set: %v", table.columns)
	}

	table.columns = s

	return nil
}

func (table *Table) QuickRow(x ...any) error {
	return table.AddRow(Row(x))
}

func (table *Table) AddRow(row Row) error {
	if table == nil {
		return fmt.Errorf("table is nil")
	}

	table.mu.Lock()
	defer table.mu.Unlock()

	cl := len(table.columns)

	if cl > 0 && len(row) != cl {
		return fmt.Errorf("expected %d columns, got %d", cl, len(row))
	}

	table.rows = append(table.rows, row)

	return nil
}
