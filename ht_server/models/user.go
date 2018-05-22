package models

import(
	"hometools/ht_server/entity"
	"hometools/ht_server/dal"
	"time"
	"fmt"
)

type User struct {
	Id       int64
	Username string
	Openid string
	Photo  string
	Description string
}
var db=dal.NewUserDal()

//登陆 1：成功 0：未注册 -1：拒绝
func (this *User) Login(openid string) (int,*entity.Sys_user){
	if len(openid)==0{
		return -1,nil
	}
	u:=this.GetUser(openid)
	if u==nil{
		return 0,nil
	}
	return 1,u
}
func (this *User) GetUser(openid string) *entity.Sys_user{
	return db.GetByOpenID(openid)
}
func (this *User) RegisterUser() error{
	u:=entity.Sys_user{
		Openid:this.Openid,
		Name:this.Username,
		Photo:this.Photo,
		Description:this.Description,
		Createby:"root",
		Createtime:time.Now(),
	}
	
	id,err:= db.Add(&u);
	if err!=nil{
		u.Id=id
	}
	return err
}
func (this *User) UpdateUser() error{
	if(this.Id==0){
		return fmt.Errorf("数据不合法")
	}
	u:=entity.Sys_user{
		Id:this.Id,
		Openid:this.Openid,
		Name:this.Username,
		Photo:this.Photo,
		Description:this.Description,
		// Createby:"root",
		// Createtime:time.Now(),
	}
	return db.Update(&u)
}
func (this *User) GetFamilyList(userID int) *[]entity.Sys_family{
	return db.GetFamilyList(userID)
}