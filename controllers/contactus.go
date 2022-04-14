package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ContactusController struct {
	beego.Controller
}

func (this *ContactusController) Get() {
	this.TplName = "contactusTemplate.html"
}

func (this *ContactusController) Post() {
	this.TplName = "thankyouTemplate.html"
	this.Data["name"] = this.GetString("name")
	this.Data["email"] = this.GetString("email")
}
