package main


import (
	_ "user-provider/routers"
	_ "go-eureka"
	"github.com/astaxie/beego"
)

func main() {

	beego.Run()
}
