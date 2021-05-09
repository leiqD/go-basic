package model

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)
type conf struct {
	a int64 `default:"10" db_conf_name:"a"`
	b string `default:"1000" db_conf_name:"b"`
	c bool `default:"true" db_conf_name:"c"`
	d []string `default:"1,2,3,4" db_conf_name:"d"`
}

func get_default_name(cc conf,i int)string {
	t := reflect.TypeOf(cc)
	x := t.Field(i)
	xx := x.Tag.Get("db_conf_name")
	return xx
}

func set_default_int64(cc conf,name string)(int64) {
	t := reflect.TypeOf(cc)
	x,ok := t.FieldByName(name)
	if !ok{
		return -1
	}
	c := x.Tag.Get("default")
	r ,err:= strconv.ParseInt(c,10,64)
	if err != nil{
		return -1
	}
	return r
}

func set_default_string(cc conf,name string)string {
	t := reflect.TypeOf(cc)
	x,ok := t.FieldByName(name)
	if !ok{
		return "-1"
	}
	c := x.Tag.Get("default")
	return c
}

func set_default_strings(cc conf,name string)[]string {
	 s := set_default_string(cc,name)
	 c := strings.Split(s,",")
	return c
}

func set_default_bool(cc conf,name string)bool {
	t := reflect.TypeOf(cc)
	x,ok := t.FieldByName(name)
	if !ok{
		return false
	}
	c := x.Tag.Get("default")
	r ,err:= strconv.ParseBool(c)
	if err != nil{
		return false
	}
	return r
}

func TestStruct(T *testing.T){
	cc := conf{}
	d := make(map[interface{}]string)
	d[&cc.a]=get_default_name(cc,0)
	d[&cc.b]=get_default_name(cc,1)
	d[&cc.c]=get_default_name(cc,2)
	d[&cc.d]=get_default_name(cc,3)
	cc.a = set_default_int64(cc,d[&cc.a])
	cc.b = set_default_string(cc,d[&cc.b])
	cc.c = set_default_bool(cc,d[&cc.c])
	cc.d = set_default_strings(cc,d[&cc.d])
	fmt.Printf("%v",d)
	fmt.Printf("%v",cc)
}
