package model_test

import (
	"fmt"
	"strconv"
	"strings"
)
import "testing"



// 长度为 1w 的数据使用系统自带排序
func BenchmarkSort10k(t *testing.B) {
	for i := 0; i < 1; i++ {
		fmt.Sprintf("%d_%d_%d_%d",100,100,1000,10000)
	}
}
// 长度为 1w 的数据使用系统自带排序
func BenchmarkSort20k(t *testing.B) {
	const sp = "_"
	for i := 0; i < 1; i++ {
		builder := strings.Builder{}
		builder.WriteString(strconv.FormatInt(100,10))
		builder.WriteString(sp)
		builder.WriteString(strconv.FormatInt(100,10))
		builder.WriteString(sp)
		builder.WriteString(strconv.FormatInt(1000,10))
		builder.WriteString(sp)
		builder.WriteString(strconv.FormatInt(10000,10))
	}
}