package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"aspe/models"
)

type QuestionController struct {
	beego.Controller
}

func (this *QuestionController) ListQuestions() {
	res := struct{Questions []*models.Question}{ models.QuestionRep.GetAll()}
	this.Data["json"] = res
	this.ServeJson()

}

func (this *QuestionController) FindById() {
	id:=this.Ctx.Input.Param(":id")
	intid,_ :=strconv.ParseInt(id,10,64)
	q, found:=models.QuestionRep.FindById(intid)
	if !found{
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	this.Data["json"] = q
	this.ServeJson()	
	beego.Info("Found",found)
}


