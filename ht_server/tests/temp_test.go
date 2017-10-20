package test

import(
	"testing"
	"hometools/ht_server/entity"
	"reflect"
)

func TestTemp1(t *testing.T){
	m:=make([]entity.Sys_family,0)
	s:=reflect.TypeOf(m).Name()
	t.Logf("type is %s",s)
}