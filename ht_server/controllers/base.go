package controllers
import(
	"github.com/astaxie/beego"
)


func Return_Json(context *beego.Controller,data interface{}){
	context.Data["json"]=data
	context.ServeJSON()
}