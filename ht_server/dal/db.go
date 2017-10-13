package dal
import(
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"hometools/ht_server/common"
)
var(
	DbContext orm.Ormer   //orm对象
)
func init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",common.Config.String("mysqlDatabase"))
	DbContext=orm.NewOrm()
	DbContext.Using("mysql")
}

func RegisterModel(models ...*interface{}){
	orm.RegisterModel(models)
}