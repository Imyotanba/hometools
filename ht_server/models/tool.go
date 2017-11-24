package models
import(
	"hometools/ht_server/entity"
	"hometools/ht_server/dal"
)
type Tool struct{
	entity.Sys_tool
}

func(this *Tool) GetAll() *[]entity.Sys_tool{
	var list []entity.Sys_tool
	dal.NewToolDal().GetAll(&list)
	return &list
}
func(this *Tool) Get(id int64) *entity.Sys_tool{
	m:=entity.Sys_tool{
		Id:id,
	}
	dal.NewToolDal().Get(&m)
	return &m
}
//获取家庭工具列表
func(this *Tool) GetTools(familyID int64) *[]entity.Sys_tool{
	return dal.NewToolDal().GetFamilyTools(familyID)
}
//添加一个工具到家庭工具箱
func(this *Tool) AddTool(toolID,familyID int64) error{
	mod:=entity.Sys_familytoolbox{
		Familyid:familyID,
		Toolid:toolID,
		State:1,
	}
	_,err:= dal.NewToolDal().Add(&mod)
	return err
}

