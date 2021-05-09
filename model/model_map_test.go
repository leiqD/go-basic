package model_test

import "testing"

type mvp struct {
	m int
	v int
}

var aa map[int]map[int]*mvp
var bb map[mvp]*mvp

// 长度为 1w 的数据使用系统自带排序
func BenchmarkMap10k(t *testing.B) {
	for i,row := range aa{
		for j,_ := range row{
			i++;j++
		}
	}
}
// 长度为 1w 的数据使用系统自带排序
func BenchmarkMap20k(t *testing.B) {
	for _,row := range bb{
		row.m++
		row.v++
	}
}

func TestMain(m *testing.M) {
	bb = make(map[mvp]*mvp)
	aa = make(map[int]map[int]*mvp)
	for i := 0; i < 10000; i++ {
		if _,ok := aa[i];!ok{
			aa[i]=make(map[int]*mvp)
		}
		aa[i][i]=&mvp{
			m:i,
			v:i,
		}
		d := mvp{
			m:i,
			v:i,
		}
		bb[d]=&d
	}
	m.Run()
}