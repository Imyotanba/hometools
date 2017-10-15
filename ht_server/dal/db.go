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

//查询基类
type baseDal struct{
	tabname string
}
func(this *baseDal) GetAll(container interface{}) {
	_,err:=DbContext.QueryTable(this.tabname).All(container)
	if err!=nil{
		this.ErrorHandler(err)
	}
}
func(this *baseDal) Get(md interface{}){
	if err:=DbContext.Read(md);err!=nil{
		this.ErrorHandler(err)
	}
}

func(this *baseDal) Add(md interface{}) (int64,error){
	return DbContext.Insert(md)
}
func(this *baseDal) Update(md interface{}) error{
	_,err:=DbContext.Update(md)
	return err
}
func(this *baseDal) Delete(md interface{}) error{
	_,err:=DbContext.Delete(md)
	return err
}

//错误统一处理
func(this *baseDal) ErrorHandler(err error){
	panic(err)
}