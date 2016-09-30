package routers

import (
	"github.com/TalLannder/app2/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
