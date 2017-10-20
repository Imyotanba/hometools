package test
import(
	"testing"
	"hometools/ht_server/models"
	_"hometools/ht_server/entity"
)
/*
func TestCreate(t *testing.T){
	m:=entity.Sys_family{
		Name:"测试组",
		Description:"一帮呆呆的",
		Label:"傻;蠢",
	}
	if err:=new(models.Family).Create(&m);err!=nil{
		t.Fatal(err)
	}else{
		t.Log(m)
	}
}
func TestUpdate(t *testing.T){
	d:=new(models.Family)
	m:=d.Get(1)
	t.Log(m)
	m.Description="拯救世界"
	if err:=d.Update(m);err!=nil{
		t.Fatal(err)
	}else{
		t.Log(m)
	}
}
*/
func TestDelete(t *testing.T){
	if err:=new(models.Family).Delete(2);err!=nil{
		t.Fatal(err)
	}
}