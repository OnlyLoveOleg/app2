package main

import (
	"encoding/gob"
	"github.com/TalLannder/app2/models"
	_ "github.com/TalLannder/app2/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func init() {
	gob.RegisterName("github.com/TalLannder/app1/models.User", models.User{})
}

func main() {
	beego.Run()
}
