package table

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const HTML_TABLE = "<table>\n<thead>\n<th>a</th><th>b</th><th>c</th><thead>\n<tbody>\n<tr>\n\t<td>1</td>\n\t<td>2</td>\n\t<td>3</td>\n</tr>\n<tr>\n\t<td>2</td>\n\t<td>4</td>\n\t<td>6</td>\n</tr>\n<tr>\n\t<td>3</td>\n\t<td>6</td>\n\t<td>9</td>\n</tr>\n<tr>\n\t<td>4</td>\n\t<td>8</td>\n\t<td>12</td>\n</tr>\n<tr>\n\t<td>5</td>\n\t<td>10</td>\n\t<td>15</td>\n</tr>\n<tr>\n\t<td>6</td>\n\t<td>12</td>\n\t<td>18</td>\n</tr>\n<tr>\n\t<td>7</td>\n\t<td>14</td>\n\t<td>21</td>\n</tr>\n<tr>\n\t<td>8</td>\n\t<td>16</td>\n\t<td>24</td>\n</tr>\n<tr>\n\t<td>9</td>\n\t<td>18</td>\n\t<td>27</td>\n</tr>\n<tr>\n\t<td>10</td>\n\t<td>20</td>\n\t<td>30</td>\n</tr>\n</tbody>\n</table>"

func Test_HTML_Print(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	table := quickTable(t)

	html := HTML{}

	bb := &bytes.Buffer{}

	err := html.Print(bb, table)
	r.NoError(err)

	act := bb.String()
	act = strings.TrimSpace(act)

	// fmt.Println(act)

	r.Equal(HTML_TABLE, act)
}
