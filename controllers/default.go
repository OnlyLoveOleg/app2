package controllers

import (
	"fmt"
	"github.com/TalLannder/app2/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	sess := c.GetSession("session")
	fmt.Println(sess.(models.User).Name)
	c.Data["Session"] = sess
	c.TplName = "index.tpl"
}
