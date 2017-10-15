package dal

import(
	_"hometools/ht_server/entity"
)
type UserDal struct{
	baseDal
	tabname string
}
func NewUserDal() *UserDal{
	u:=new(UserDal)
	u.tabname="Sys_user"
	return u
}