package dal

import(
	"hometools/ht_server/entity"
)

type FamilyDal struct{
	baseDal
}

func NewFamilyDal() *FamilyDal{
	m:=new (FamilyDal)
	m.tabname=new(entity.Sys_family).TableName()
	return m
}

//获取家庭成员列表
func (this *FamilyDal) GetMembers(familyid int64) *[]entity.V_Sys_FamilyUserMembers{
	list:=make([]entity.V_Sys_FamilyUserMembers,0)
	if _,err:=DbContext.QueryTable(new(entity.V_Sys_FamilyUserMembers)).Filter("FamilyID",familyid).All(&list);err!=nil{
		this.errorHandler(err)
	} 
	return &list
}

//移除从家庭一名成员
// func (this *FamilyDal) RemoveMember(userid int,familyid int){
// 	DbContext.Delete
// }