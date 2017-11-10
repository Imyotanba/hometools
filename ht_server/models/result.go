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