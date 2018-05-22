package controllers

import (
	"hometools/ht_server/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}


//获取信息
// @router /:openid [get]
func (u *UserController) GetUser(){
	openid:=u.GetString(":openid")
	Return_Json(&u.Controller,new(models.User).GetUser(openid))
}

//注册
func (u *UserController) Post(){
	result:= models.Result{}
	user:= &models.User{}
	if err:=json.Unmarshal(u.Ctx.Input.RequestBody,user);err!=nil{
		result.Code=-1
		result.Msg=fmt.Sprintf("序列化数据失败:%s",err)
	}else{
		if err:=user.RegisterUser();err!=nil{
			result.Code=0
			result.Msg=fmt.Sprintf("注册失败:%s",err)
		}else{
			result.Code=1
			result.Msg="success"
		}
	}
	Return_Json(&u.Controller,&result)
}
//修改信息
func (u *UserController) Put(){
	result:= models.Result{}
	user:= &models.User{}
	if err:=json.Unmarshal(u.Ctx.Input.RequestBody,user);err!=nil{
		result.Code=-1
		result.Msg=fmt.Sprintf("序列化数据失败:%s",err)
	}else{
		if err:=user.UpdateUser();err!=nil{
			result.Code=0
			result.Msg=fmt.Sprintf("修改信息失败:%s",err)
		}else{
			result.Code=1
			result.Msg="success"
		}
	}
	Return_Json(&u.Controller,&result)
}

//登陆接口
// @router /login/:openid [get]
func (u *UserController) Login() {
	result:= models.Result{}
	openid:=u.GetString(":openid")
	result.Code,result.Data=new(models.User).Login(openid)
	switch result.Code{
	case 1:
			result.Msg="success"
	case 0:
			result.Msg="not register user"
	default:
			result.Msg="denial login"
	}
	Return_Json(&u.Controller,&result)
}
//获取family列表
// @router /familylist/:userid [get]
func (u *UserController) FamilyList(){
	userid,err:=u.GetInt(":userid")
	if err==nil{
		Return_Json(&u.Controller,new(models.User).GetFamilyList(userid))
	}
}
