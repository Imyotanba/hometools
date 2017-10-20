package entity
import(
	"time"
)

type Sys_family struct{
	Id int64 `orm:"pk"`
	Name string
	Photo string
	Description string
	Label string
	Createby string
	Createtime time.Time
}
func (this *Sys_family) TableName() string{
	return "sys_family"
}