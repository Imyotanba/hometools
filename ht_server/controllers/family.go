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
		Return_Json(&this.Controller,&info)
	}else{
		result:=models.Result{}
		result.Error(fmt.Sprintf("error: %s",err))
		Return_Json(&this.Controller,&result)
	}
}

//创建
func(this *FamilyController) Post(){
	result:=models.Result{}
	fam:= &models.Family{}
	if err:=json.Unmarshal(this.Ctx.Input.RequestBody,fam);err!=nil{
		result.Error(fmt.Sprintf("序列化数据失败:%s",err))
		Return_Json(&this.Controller,&result)
		return
	}
	createBy:=this.GetSession("CurrentUser").(models.User)
	if(&createBy==nil){
		result.Error("登录信息失效")
		Return_Json(&this.Controller,&result)
		return
	}
	if err:=fam.Create(createBy.Username);err!=nil{
		result.Code=0
		result.Msg=fmt.Sprintf("创建失败:%s",err)
	}else{
		result.Success()
	}
}
//修改资料
func(this *FamilyController) Put(){
	result:=models.Result{}
	fam:= &models.Family{}
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
	Return_Json(&this.Controller,&result)
}
//获取所有成员
// @router /:id/members [get]
func(this *FamilyController) GetMembers(){
	id,err:=this.GetInt64(":id")
	if(err==nil){
		members:=new(models.Family).Members(id)
		Return_Json(&this.Controller,&members)
	}
}
//添加成员
// @router /members/:id [post]
func(this *FamilyController) AddMember(){
	result:=models.Result{}
	fid,err:=this.GetInt64(":id")
	fmt.Print(string(this.Ctx.Input.RequestBody))
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
	Return_Json(&this.Controller,&result)
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
	Return_Json(&this.Controller,&result)
}
