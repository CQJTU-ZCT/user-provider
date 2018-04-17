/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/16 16:16.
 */
package common

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"strings"
	"user-provider/models"
	"log"
	"encoding/json"
	"net/http"
	"encoding/base64"
	"io/ioutil"
)

var AuthFilter = func(ctx *context.Context) {
	next := false
	validateUrl := beego.AppConfig.String("validateUrl")
	token := ctx.Input.Query("token")
	message := models.Message{Code:500,Info:"验证服务器忙，验证身份失败，请稍后再试..."}
	if !strings.HasPrefix(validateUrl,"https://") && !strings.HasPrefix(validateUrl,"http://"){
		validateUrl = "http://"+validateUrl
	}
	if len(token) <=0 {
		token = ctx.Input.Header("token")
	}
	if len(token) <= 0 {
		message.Code=200
		message.Info = "token不能为空"
	}else {
		client := &http.Client{}
		encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
		log.Println("验证服务"," GET ",validateUrl +"?token="+token)
		newRequest , err :=http.NewRequest("GET",validateUrl+"?token="+token,nil)
		if err != nil {
			log.Println("验证服务出错",err)
		}else {
			newRequest.Header.Set("Authorization","Basic "+ base64.NewEncoding(encodeStd).EncodeToString([]byte("test:test")))
			resp , err := client.Do(newRequest)
			if err != nil {
				log.Println(err)
			}else {
				body ,err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Println(err)
				}else {
					var users models.Users
					var message models.Message
					err := json.Unmarshal(body ,&message)
					if err != nil {
						log.Println(err)
					}else {
						if  userMap , ok := message.Map["user"].(map[string]interface {});ok{
							if strings.Trim(userMap["idCard"].(string), " ") != "" {
								next = true
								users.IdCard = userMap["idCard"].(string)
								users.RoleId = int(userMap["roleId"].(float64))
								users.AccountStatusId = int(userMap["accountStatusId"].(float64))
								users.RealName = userMap["realname"].(string)
								users.Mail = userMap["mail"].(string)
								users.Phone = userMap["phone"].(string)
								ctx.Input.SetData("user",users)
							}
						}
					}
				}
			}
		}
	}

	if next == false {
		ctx.ResponseWriter.WriteHeader(message.Code)
		ctx.WriteString(message.ToJson())
		return
	}
}