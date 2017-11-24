package controllers

import(
	"github.com/astaxie/beego"
	"hometools/ht_server/models"
	"encoding/json"
	"fmt"
)
type FamilyController struct {
	beego.Controller
}

//获取信息
// @router /:id [get]
func(this *FamilyController) GetFamily(){
	id,err:=this.GetInt64(":id")
	if(err==nil){
		info:=new(models.Family).Get(id)
		this.Data["json"]=info
		this.ServeJSON()
	}
}

//创建
func(this *FamilyController) Post(){
	result:=models.Result{}
	var fam models.Family
	if err:=json.Unmarshal(this.Ctx.Input.RequestBody,fam);err!=nil{
		result.Code=-1
		result.Msg=fmt.Sprintf("序列化数据失败:%s",err)
	} else{
		createBy:=this.GetSession("CurrentUser").(models.User)
		if err:=fam.Create(createBy.Username);err!=nil{
			result.Code=0
			result.Msg=fmt.Sprintf("创建失败:%s",err)
		}else{
			result.Success()
		}
	}
	this.Data["json"]=result
	this.ServeJSON()
}
//修改资料
func(this *FamilyController) Put(){
	result:=models.Result{}
	var fam models.Family
	if err:=json.Unmarshal(this.Ctx.Input.RequestBody,fam);err!=nil{
		result.Code=-1
		result.Msg=fmt.Sprintf("序列化数据失败:%s",err)
	} else{
		if err:=fam.Update();err!=nil{
			result.Code=0
			result.Msg=fmt.Sprintf("创建失败:%s",err)
		}else{
			result.Success()
		}
	}
	this.Data["json"]=result
	this.ServeJSON()
}
//获取所有成员
// @router /:id/members [get]
func(this *FamilyController) GetMembers(){
	id,err:=this.GetInt64(":id")
	if(err==nil){
		members:=new(models.Family).Members(id)
		this.Data["json"]=members
		this.ServeJSON()
	}
}
//添加成员
// @router /:id/members [post]
func(this *FamilyController) AddMember(){
	result:=models.Result{}
	fid,err:=this.GetInt64(":id")
	if(err!=nil){
		return
	}
	userid,err:=this.GetInt64("userid")
	if(err!=nil){
		return
	}
	if err:= new(models.Family).AddMember(fid,userid);err==nil{
		result.Success()
	}else{
		result.Code=0
		result.Msg=fmt.Sprint(err)
	}
	this.Data["json"]=result
	this.ServeJSON()
}
//移除成员
//@router /:id/members [delete]
func(this *FamilyController) DelMember(){
	result:=models.Result{}
	fid,err:=this.GetInt64(":id")
	if(err!=nil){
		return
	}
	userid,err:=this.GetInt64("userid")
	if(err!=nil){
		return
	}
	if err:= new(models.Family).RemoveMember(fid,userid);err==nil{
		result.Success()
	}else{
		result.Code=0
		result.Msg=fmt.Sprint(err)
	}
	this.Data["json"]=result
	this.ServeJSON()
}
