package main

import (
	_ "golang-catapi/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

