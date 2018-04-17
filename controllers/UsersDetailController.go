/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/14 14:57.
 */
package controllers

import (
	"github.com/astaxie/beego"
	"user-provider/models"
	"net/http"
	"user-provider/common"
	"user-provider/service"
)

type UsersDetailController struct {
	beego.Controller
}


func (this *UsersDetailController)GetUsersDetailByIdCard()  {
	message := models.Message{}
	message.Code = 200
	if this.Ctx.Request.Method == http.MethodGet {
		userInterface:= this.Ctx.Input.GetData("user")
		if user , ok := userInterface.(models.Users);ok {
			var idCard string
			if user.RoleId >=2 {
				idCard = this.GetString("idCard")
			}else {
				idCard = user.IdCard
			}
			if common.IsIdCard(idCard) {
				data := service.GetUsersDetailByIdCard(idCard)
				message = message.Put("usersDetail",data)
			}else {
				message.Info = "身份证号码格式不正确"
			}
		}else {
			message.Code = 401
			message.Info = "校验权限出错，请尝试重新登录"
		}
	}else {
		message.Info = "Method not be allowed"
		message.Code = 405
	}
	this.Ctx.ResponseWriter.WriteHeader(message.Code)
	this.Ctx.WriteString(message.ToJson())
}
