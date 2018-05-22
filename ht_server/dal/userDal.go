package dal

import(
	"hometools/ht_server/entity"
)

type UserDal struct{
	baseDal
}
func NewUserDal() *UserDal{
	u:=new(UserDal)
	u.tabname=new(entity.Sys_user).TableName()
	return u
}
func(this *UserDal) Update(user *entity.Sys_user) error{
	 _,err:= DbContext.Update(user,"name","photo","Description")
	 return err
}
func(this *UserDal) GetByOpenID(openid string) *entity.Sys_user{
	row:= DbContext.QueryTable(this.tabname).Filter("openid",openid)
	user:=new(entity.Sys_user)
	if err:=row.One(user);err!=nil{
		return nil
	}else{
		return user
	}
}
func(this *UserDal) GetFamilyList(userid int) *[]entity.Sys_family{
	var list []entity.Sys_family
	row:=DbContext.Raw("select t1.* from Sys_Family t1,Sys_FamilyMember t2 where t1.ID=t2.FamilyID and t2.UserID=?",userid)
	if _,err:=row.QueryRows(&list);err!=nil{
		this.errorHandler(err)
	}
	return &list
}