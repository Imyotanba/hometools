package entity
import(
	"time"
)
type V_Sys_FamilyUserMembers struct{
	Id int 
	Userid int
	Familyid int
	Userstatus int
	Openid string
	Username string
	Userphoto string
	Userdescription string
	Familyname string
	Familyphoto string
	Familydescription string
	Familylabel string
	Familycreateby string
	Familycreatetime time.Time
}
func (this *V_Sys_FamilyUserMembers) TableName() string{
	return "V_Sys_FamilyUserMembers"
}