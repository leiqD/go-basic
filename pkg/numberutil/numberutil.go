package numberutil

import (
	"math"
	"math/big"
	"net"
	"strconv"
	"strings"
)

// Abs 取x的绝对值
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// InetAtoN
func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

// Round 四舍五入
func Round(x float64) int {
	return int(math.Floor(x + 0.5))
}

func SliceToString(a []int64, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(b, sep)
}

func StringSliceToInt64Slice(a []string) []int64 {
	if len(a) == 0 {
		return []int64{}
	}

	b := make([]int64, len(a))
	for _, v := range a {
		data, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		b = append(b, data)
	}
	return b
}
