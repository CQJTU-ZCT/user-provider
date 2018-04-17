/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/17 11:38.
 */
package controllers

import (
	"github.com/astaxie/beego"
	"user-provider/models"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController)Error404()  {
	message :=models.Message{Code:404 , Info:"No handler found for "+this.Ctx.Request.Method +"  "+this.Ctx.Request.RequestURI}
	this.Ctx.ResponseWriter.WriteHeader(message.Code)
	this.Ctx.WriteString(message.ToJson())
}


func (this *ErrorController)Error500()  {
	message :=models.Message{Code:500 , Info:"服务器忙，请稍后再试..."}
	this.Ctx.ResponseWriter.WriteHeader(message.Code)
	this.Ctx.WriteString(message.ToJson())
}

