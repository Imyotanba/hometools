package entity

type Sys_familytoolbox struct{
	Id int
	Familyid int
	Toolid int
	//启用状态 1 or 0
	State rune
}
func (this *Sys_familytoolbox) TableName() string{
	return "sys_familytoolbox"
}