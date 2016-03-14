package routers

import (
	"github.com/yuankui/glogger/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
