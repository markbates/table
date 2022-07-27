package table

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func quickTable(t *testing.T) *Table {
	t.Helper()

	table := &Table{}

	cols := []string{"a", "b", "c"}

	r := require.New(t)

	err := table.SetColumns(cols...)
	r.NoError(err)

	for i := 0; i < 10; i++ {
		i := i + 1
		nums := []any{i * 1, i * 2, i * 3}
		err = table.QuickRow(nums...)
		r.NoError(err)
	}

	return table
}

func Test_Table_QuickRow(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	table := &Table{}

	err := table.QuickRow(1, 2, 3)
	r.NoError(err)

	r.Len(table.rows, 1)

	rows, err := table.Rows()
	r.NoError(err)

	r.Len(rows, 1)
}

func Test_Table_SetColumns(t *testing.T) {
	t.Parallel()

	cols := []string{"a", "b", "c"}
	nums := []any{1, 2, 3}

	short := &Table{}

	long := &Table{}

	tcs := []struct {
		name  string
		table *Table
		vals  []any
		err   bool
	}{
		{name: "empty table", table: &Table{}, vals: nums},
		{name: "short size", table: short, vals: []any{1}, err: true},
		{name: "long size", table: long, vals: []any{1, 2, 3, 4}, err: true},
	}

	for _, tc := range tcs {

		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			table := tc.table

			err := table.SetColumns(cols...)

			r.NoError(err)

			vals := tc.vals

			err = table.QuickRow(vals...)

			if tc.err {
				r.Error(err)
				return
			}

			r.NoError(err)
		})

	}

}

func Test_Table_MaxColumnSize(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	a := &Table{}
	err := a.SetColumns("a", "b", "c")
	r.NoError(err)

	err = a.QuickRow(1, 1234, 3)
	r.NoError(err)

	b := &Table{}
	err = b.SetColumns("a", "bbbbb", "c")
	r.NoError(err)

	tcs := []struct {
		name  string
		table *Table
		exp   int
	}{
		{name: "a", table: a, exp: 4},
		{name: "b", table: b, exp: 5},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			table := tc.table

			act, err := table.MaxColumnWidth()
			r.NoError(err)

			r.Equal(tc.exp, act)

		})
	}
}
