package table

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Markdown_Print(t *testing.T) {
	// t.Skip()
	t.Parallel()
	r := require.New(t)

	bb := &bytes.Buffer{}

	table := quickTable(t)

	md := Markdown{}

	err := md.Print(bb, table)
	r.NoError(err)

	act := bb.String()
	act = strings.TrimSpace(act)

	// fmt.Println(act)
	exp := `| a   | b   | c   |
| --- | --- | --- |
| 1   | 2   | 3   |
| 2   | 4   | 6   |
| 3   | 6   | 9   |
| 4   | 8   | 12  |
| 5   | 10  | 15  |
| 6   | 12  | 18  |
| 7   | 14  | 21  |
| 8   | 16  | 24  |
| 9   | 18  | 27  |
| 10  | 20  | 30  |`

	r.Equal(exp, act)

}

func Test_Markdown_Print_DisableHeaderLine(t *testing.T) {
	// t.Skip()
	t.Parallel()
	r := require.New(t)

	bb := &bytes.Buffer{}

	table := quickTable(t)

	md := Markdown{
		DisableHeaderLine: true,
	}

	err := md.Print(bb, table)
	r.NoError(err)

	act := bb.String()
	act = strings.TrimSpace(act)

	// fmt.Println(act)
	exp := `| a   | b   | c   |
| 1   | 2   | 3   |
| 2   | 4   | 6   |
| 3   | 6   | 9   |
| 4   | 8   | 12  |
| 5   | 10  | 15  |
| 6   | 12  | 18  |
| 7   | 14  | 21  |
| 8   | 16  | 24  |
| 9   | 18  | 27  |
| 10  | 20  | 30  |`

	r.Equal(exp, act)

}

func Test_Markdown_Print_DisableHeader(t *testing.T) {
	// t.Skip()
	t.Parallel()
	r := require.New(t)

	bb := &bytes.Buffer{}

	table := quickTable(t)

	md := Markdown{
		DisableHeader: true,
	}

	err := md.Print(bb, table)
	r.NoError(err)

	act := bb.String()
	act = strings.TrimSpace(act)

	// fmt.Println(act)
	exp := `| 1   | 2   | 3   |
| 2   | 4   | 6   |
| 3   | 6   | 9   |
| 4   | 8   | 12  |
| 5   | 10  | 15  |
| 6   | 12  | 18  |
| 7   | 14  | 21  |
| 8   | 16  | 24  |
| 9   | 18  | 27  |
| 10  | 20  | 30  |`

	r.Equal(exp, act)

}
