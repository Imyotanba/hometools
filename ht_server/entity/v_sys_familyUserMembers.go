package entity
import(
	"time"
)
type V_Sys_FamilyUserMembers struct{
	Id int 
	UseriD int
	FamilyiD int
	Userstatus int
	Openid string
	Username string
	Userphoto string
	Userdescription string
	Familyname string
	Familyphoto string
	Familydescription string
	Familylabel string
	FamilycreateBy string
	Familycreatetime time.Time
}
func (this *V_Sys_FamilyUserMembers) TableName() string{
	return "V_Sys_FamilyUserMembers"
}