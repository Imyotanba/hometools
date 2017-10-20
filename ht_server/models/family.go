package models
import(
	"hometools/ht_server/entity"
	"hometools/ht_server/dal"
	"time"
)

type Family struct{
	//entity.Sys_family
}

//获取信息
func (this *Family) Get(id int64) *entity.Sys_family{
	fam:=new(entity.Sys_family)
	fam.Id=id
	dal.NewFamilyDal().Get(fam)
	return fam
}
//创建
func (this *Family) Create(mode *entity.Sys_family) error{
	mode.Createtime=time.Now()
	if id,err:=dal.NewFamilyDal().Add(mode);err!=nil{
		return err
	}else{
		mode.Id=id
		return nil
	}
}
//修改
func (this *Family) Update(mode *entity.Sys_family) error{
	return dal.NewFamilyDal().Update(mode)
}
//删除
func (this *Family) Delete(id int64) error{
	mode:=entity.Sys_family{
		Id:id,
	}
	return dal.NewFamilyDal().Delete(&mode)
}
//获取家庭成员
func (this *Family) Members(familyid int) *[]entity.V_Sys_FamilyUserMembers{
	return dal.NewFamilyDal().GetMembers(familyid)
}
//添加一名成员
func (this *Family) AddMember(familyid,userid int64) error{
	fm:=entity.Sys_familymember{
		Familyid:familyid,
		Userid:userid,
		Userstatus:0,
	}
	_,err:=dal.NewFamilyDal().Add(&fm)
	return err
}
//移除一名成员
func (this *Family) RemoveMember(familyid,userid int64) error{
	fm:=entity.Sys_familymember{
		Familyid:familyid,
		Userid:userid,
	}
	return dal.NewFamilyDal().Delete(&fm)
}
//获取工具列表
/*
func (this *Family) Tools(id int) *[]entity.Sys_tool{

}
*/
