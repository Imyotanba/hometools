package models

type Result struct{
	Code int
	Msg string
	Data interface{}
}
func(this *Result) Success(){
	this.Code=1
	this.Msg="success"
}
func(this *Result) Success_Data(data *interface{}){
	this.Code=1
	this.Msg="success"
	this.Data=data
}
func(this *Result) Failed(message string){
	this.Code=0
	this.Msg=message
}
func(this *Result) Error(err string){
	this.Code=-1
	this.Msg=err
}