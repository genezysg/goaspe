package routers

import (
	"aspe/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.QuestionController{},"get:ListQuestions;post:NewQuestion")
		beego.Router("/:id:int",&controllers.QuestionController{},"get:FindById")
}
