package dal

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"hometools/ht_server/common"
	"hometools/ht_server/entity"
)

var (
	DbContext orm.Ormer //orm对象
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", common.Config.String("mysqlDatabase"))
	//注册实体表
	registerModel()

	DbContext = orm.NewOrm()
	DbContext.Using("default")
}
//注册实体表
func registerModel() {
	orm.RegisterModel(
		new(entity.Sys_user),
		new(entity.Sys_family),
		new(entity.Sys_tool),
		new(entity.Sys_familymember),
		new(entity.Sys_familytoolbox),
		new(entity.V_Sys_FamilyUserMembers),
	)
}

//查询基类
type baseDal struct {
	//实体对应表名，子类必须重写
	tabname string
}

func (this *baseDal) GetAll(container interface{}) {
	_, err := DbContext.QueryTable(this.tabname).All(container)
	if err != nil {
		this.errorHandler(err)
	}
}
func (this *baseDal) Get(md interface{}) error {
	return DbContext.Read(md)
}

func (this *baseDal) Add(md interface{}) (int64, error) {
	return DbContext.Insert(md)
}
func (this *baseDal) Update(md interface{}) error {
	_, err := DbContext.Update(md)
	return err
}
func (this *baseDal) Delete(md interface{}) error {
	_, err := DbContext.Delete(md)
	return err
}

//错误统一处理
func (this *baseDal) errorHandler(err error) {
	panic(err)
}
