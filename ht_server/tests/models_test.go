package test

import(
	"testing"
	"hometools/ht_server/models"
)

func TestGetAllTool(t *testing.T){
	m:=new(models.Tool)
	list:= m.GetAll()
	t.Log(list)
}
func TestGetTool(t *testing.T){
	m:=new(models.Tool)
	t.Log(m.Get(1))
}