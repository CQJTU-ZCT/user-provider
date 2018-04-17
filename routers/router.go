package routers

import (
	"github.com/astaxie/beego"
	"user-provider/controllers"
	"user-provider/common"
)

func init() {

	beego.Router("/users/register",&controllers.RegisterController{},"*:Register")
	beego.Router("/exits/",&controllers.UsersController{} ,"*:IsExit")
	beego.Router("/exits/:IdCard",&controllers.UsersController{} ,"*:IsExit")


	//查看用户基础信息  idCard 可以使用url参数，也可以path参数   普通用户的idCard不生效
	beego.Router("/users/",&controllers.UsersController{},"*:GetUsersByIdCard")
	beego.Router("/users/:idCard",&controllers.UsersController{},"*:GetUsersByIdCard")

	//查看用户详细信息  idCard 可以使用url参数，也可以path参数   普通用户的idCard不生效
	beego.Router("/users-detail/",&controllers.UsersDetailController{},"*:GetUsersDetailByIdCard")
	beego.Router("/users-detail/:idCard",&controllers.UsersDetailController{},"*:GetUsersDetailByIdCard")

	//修改密码
	beego.Router("/user/password/",&controllers.UsersController{},"*:ChangePassword")




	//添加过滤器
	beego.InsertFilter("/users/",beego.BeforeRouter,common.AuthFilter)
	beego.InsertFilter("/users/:idCard",beego.BeforeRouter,common.AuthFilter)

	beego.InsertFilter("/users-detail/",beego.BeforeRouter,common.AuthFilter)
	beego.InsertFilter("/users-detail/:idCard",beego.BeforeRouter,common.AuthFilter)

	beego.InsertFilter("/user/password/",beego.BeforeRouter,common.AuthFilter)


	//注册错误路由
	beego.ErrorController(&controllers.ErrorController{})
}
