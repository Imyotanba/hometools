package common
import(
	"github.com/astaxie/beego/config"
	"fmt"
)
var(
	Config config.Configer
)
func init(){
	c,err:=config.NewConfig("ini","conf\\app.conf")
	if err!=nil{
		panic(fmt.Errorf("init configer failed,%s",err))
	}else{
		Config=c
	}
}