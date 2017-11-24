package dal
import(
	"hometools/ht_server/entity"
)
type ToolDal struct{
	baseDal
}

func NewToolDal() *ToolDal{
	m:=new(ToolDal)
	m.tabname=new(entity.Sys_tool).TableName()
	return m
}
func(this *ToolDal) GetFamilyTools(familyID int64) *[]entity.Sys_tool{
	sql:="select t1.* from sys_tool t1,sys_familytoolbox t2 where t1.ID=t2.ToolID and t2.FamilyID=?"
	list:=make([]entity.Sys_tool,0)
	if err:= DbContext.Raw(sql,familyID).QueryRow(&list);err!=nil{
		this.errorHandler(err)
	}
	return &list
}