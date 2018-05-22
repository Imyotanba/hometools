package entity

type Sys_familymember struct{
	Id int64
	Familyid int64
	Userid int64
	//身份 1.家主 0.其他
	Userstatus int
}
func (this *Sys_familymember) TableName() string{
	return "sys_familymember"
}