/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/13 13:31.
 */
package controllers

import (
	"github.com/astaxie/beego"
	"user-provider/models"
	"net/http"
	"strings"
	"user-provider/common"
	"user-provider/service"
)

type RegisterController struct {
	beego.Controller
}




/**
 * 用户注册控制器
 */
func (this *RegisterController)Register()  {
	message := models.Message{}
	message.Code = 200
	if this.Ctx.Request.Method != http.MethodPost {
		message.Code = 405
		message.Info = "Method not be allowed"
	}else {
		idCard := this.GetString("idCard")
		if common.IsIdCard(idCard) {
			realName := this.GetString("realName")
			if len(strings.Trim(realName, " ")) > 0 {
				phone := this.GetString("phone")
				if common.IsPhone(phone) {
					mail := this.GetString("mail")
					if common.IsMail(mail) {
						password := this.GetString("password")
						if len(strings.Trim(password, " ")) >=8 {
							if !service.IsExits(idCard){
								users := models.Users{}
								users.RealName = realName
								users.Mail = mail
								users.Phone = phone
								users.RoleId = 1
								users.AccountStatusId = 1
								users.IdCard = idCard
								users.Password= password
								if service.AddUsers(users){
									message.Info = "注册成功"
								}else {
									message.Info = "注册失败，请检查数据完整性"
								}
							}else {
								message.Info = "该身份证号码已经注册，请勿重复注册"
							}
						}else {
							message.Info = "密码长度大于8位"
						}
					}else {
						message.Info = "邮箱格式不正确"
					}
				}else {
					message.Info = "手机号码格式不正确"
				}
			}else {
				message.Info = "真实姓名不能为空"
			}
		}else {
			message.Info = "身份证号码不合法"
		}
	}
	this.Ctx.ResponseWriter.WriteHeader(message.Code)
	this.Ctx.WriteString(message.ToJson())
}
