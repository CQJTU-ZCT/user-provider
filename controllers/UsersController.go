/**
 * @Desc
 * @author zjhfyq
 * @data 2018/4/13 17:36.
 */
package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"user-provider/common"
	"user-provider/models"
	"user-provider/service"
	"strings"
)

type UsersController struct {
	beego.Controller
}

/**
 * 根据身份证号码获取用户信息 精确查询
 */
func (this *UsersController) GetUsersByIdCard() {
	message := models.Message{}
	message.Code = 200
	if this.Ctx.Request.Method == http.MethodGet {
		userInterface :=this.Ctx.Input.GetData("user")
		if user , ok := userInterface.(models.Users);ok {
			var idCard =""
			if user.RoleId >= 2 {
				//只有相关权限才能查看别人的信息
				idCard = this.GetString("idCard")
			}else {
				idCard = user.IdCard
			}
			if common.IsIdCard(idCard) {
				usersResp := service.GetUsersByIdCard(idCard)
				message = message.Put("user", usersResp)
				message.Info = "查询用户信息成功"
			} else {
				message.Info = "身份证号码（idCard）不正确"
			}
		}else {
			message.Code = 401
			message.Info = "校验权限出错，请尝试重新登录"
		}
	} else {
		message.Code = 405
		message.Info = "Method not be allowed"
	}

	this.Ctx.ResponseWriter.WriteHeader(message.Code)
	this.Ctx.WriteString(message.ToJson())
}

/**
 * 根据身份证号码 判断用户是否存在
 */
func (this *UsersController) IsExit() {
	message := models.Message{}
	message.Code = 200
	if this.Ctx.Request.Method == http.MethodGet {
		idCard := this.GetString("idCard")
		if common.IsIdCard(idCard) {
			result := service.IsExits(idCard)
			message = message.Put("exits", result)
			message.Info = "成功"
		} else {
			message.Info = "身份证号码（idCard）不正确"
		}
	} else {
		message.Code = 405
		message.Info = "Method not be allowed"
	}
	this.Ctx.ResponseWriter.WriteHeader(message.Code)
	this.Ctx.WriteString(message.ToJson())
}

func (this *UsersController) ChangePassword()  {
	message := models.Message{Code:200 }
	if this.Ctx.Request.Method == http.MethodPut {
		newPassword := this.GetString("newPassword")
		code := this.GetString("code")
		if len(strings.Trim(newPassword, " ")) <8 {
			message.Info = "密码长度需要大于等于8位"
		}else {
			if  len(strings.Trim(code, " ")) < 4{
				message.Info = "验证码错误"
			}else {
				userInterface := this.Ctx.Input.GetData("user")
				if user , ok  := userInterface.(models.Users);ok {
					if common.RedisGet(user.IdCard+".code") == code {
						message.Info = "验证码错误"
					}else {
						if service.UpdatePassword(newPassword ,user.IdCard){
							message.Info = "修改密码成功"
						}else {
							message.Info = "修改密码失败，请尝试重新修改"
						}
					}
				}
			}
		}
	}else {
		message.Code = 405
		message.Info = "Method not be allowed"
	}
	this.Ctx.ResponseWriter.WriteHeader(message.Code)
	this.Ctx.WriteString(message.ToJson())
}