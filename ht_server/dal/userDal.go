package dal

import(
	"hometools/ht_server/entity"
)
type User entity.Sys_user
func(this *User) GetAll() []entity.Sys_user{
	list:=make([]entity.Sys_user,0)
	DbContext.QueryTable(this).All(list)
	return list
}
