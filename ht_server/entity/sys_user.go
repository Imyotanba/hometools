package entity
import(
	"time"
)

type Sys_user struct{
	ID int64
	OpenID string
	Name string
	Photo []byte
	Description string
	CreateBy string
	CreateTime time.Time
}