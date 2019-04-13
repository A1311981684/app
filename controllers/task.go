package controllers

import (
	"app/webJsonServe"
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}


//@Title Run Task
//@Description
//@Param id query int true id of task
//@Success 200 OK
//@Failure 500 Error
//@router /run [put]
func (t *TaskController) RunTask() {
	id, err := t.GetInt("id")
	if err != nil {
		webJsonServe.ServeFailed(&t.Controller, 500, err.Error())
		return
	}
	webJsonServe.ServeSuccess(&t.Controller, "success", id)
}
