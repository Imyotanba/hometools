package entity
import(
	"time"
)

type Sys_user struct{
	Id int64 `orm:"pk"`
	Openid string
	Name string
	Photo string
	Description string
	Createby string
	Createtime time.Time
}
func (this *Sys_user) TableName() string{
	return "sys_user"
}