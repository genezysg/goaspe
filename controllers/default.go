package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"aspe/models"
)

type QuestionController struct {
	beego.Controller
}

func (this *QuestionController) ListQuestions() {
	res := struct{Questions []*models.Question}{ models.Questionrepository.GetAll()}
	this.Data["json"] = res
	this.ServeJson()

}

func (this *QuestionController) FindById() {
	id:=this.Ctx.Input.Param(":id")
	intid,_ :=strconv.ParseInt(id,10,64)
	q, found:=models.Questionrepository.FindById(intid)
	if !found{
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	this.Data["json"] = q
	this.ServeJson()	
	beego.Info("Found",found)
}


func (this *QuestionController) NewQuestion(){
	var q models.Question
	
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &q); err != nil {
		fmt.Println(this.Ctx.Input.RequestBody)
		this.Ctx.Output.SetStatus(400)
		//models.Questionrepository.Save(&q)
		beego.Info(err.Error())
		return
	}
	models.Questionrepository.Save(&q)
	return
	
	
	
}


