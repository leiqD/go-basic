package conf

import (
	"reflect"
	"strconv"
)

type conf struct {
	a int64 `default:10`
	b string `default:1000`
}
func setDefault(v interface{}){

}
func confDefault()*conf{
	con := &conf{}
	typ := reflect.TypeOf(con.a)
	t,_ := typ.FieldByName("a")
	v := t.Tag.Get("default")
	con.a,_ = strconv.ParseInt(v,10,64)
	return con
}