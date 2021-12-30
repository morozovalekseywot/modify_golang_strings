package best_strings

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestBestString(t *testing.T) {
	str := "simka and stimka"
	bs := newBS(str)

	bs.ToUpperFirst()
	require.Equal(t, bs.str, "Simka and stimka")

	bs.ReplaceAll("stimka", "second simka")
	require.Equal(t, bs.str, "Simka and second simka")

	bs.Set("123414*&()#@$%  \n\t  testify")
	bs.ToUpperFirst()
	require.Equal(t, bs.str, "123414*&()#@$%  \n\t  Testify")
	ans, _ := bs.IsUpper(strings.Index(bs.str, "T"))
	require.True(t, ans)

	bs.ToLowerFirst()
	require.Equal(t, bs.str, "123414*&()#@$%  \n\t  testify")
	ans, _ = bs.IsLower(strings.Index(bs.str, "t"))
	require.True(t, ans)

	bs.Set("\t\n  \t word \n\t")
	bs.TrimSpace()
	require.Equal(t, bs.str, "word")
	require.Equal(t, bs.size, 4)

	_, err := bs.IsDigit(4)
	require.EqualError(t, err, "Index out of range, index: 4, string size: 4\n")
}
