package routers

import (
	"textIntelligent/controllers"
	"github.com/astaxie/beego"
)

func init() {
 //   beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/intelligent/text/getLanguageDetection",&controllers.TextIntelligentController{})
}
