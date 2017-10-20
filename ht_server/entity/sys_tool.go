package entity
import(
	"time"
)
type Sys_tool struct{
	Id int64 `orm:"pk"`
	Name string
	Photo string
	Href string
	Description string
	Createby string
	Createtime time.Time
}
func (this *Sys_tool) TableName() string{
	return "sys_tool"
}