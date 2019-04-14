package controllers

import (
	"app/models"
	"app/models/checksum"
	"app/models/untar"
	"app/webJsonServe"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"path"
	"strings"
)

type UpdateController struct {
	beego.Controller
}

func (u *UpdateController)Get(){
	u.TplName = "upload.html"
}

//@Title Get update package
//@Description
//@Param fileName body file true Uploaded Update Package
//@Success 200 OK
//@Failure 500 Failed
//@router /package [post]
func (u *UpdateController) AcquireUpdate(){

	f, h, err := u.GetFile("fileName")
	if err != nil {
		webJsonServe.ServeFailed(&u.Controller, 500, err.Error())
		return
	}
	fileName := h.Filename
	log.Println("File name:", fileName)

	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}
	fmt.Println(fileName)

	err = f.Close()
	if err != nil {
		webJsonServe.ServeFailed(&u.Controller, 500, err.Error())
		return
	}

	err = u.SaveToFile("fileName", path.Join(models.UpdatePath,fileName))
	if err != nil {
		webJsonServe.ServeFailed(&u.Controller, 500, err.Error())
		return
	}

	u.TplName = "upload.html"
}

//@Title Start update
//@Description
//@Success 200 OK
//@Failure 500 Failed
//@router /start [post]
func (u *UpdateController)Start(){
	//1 un-tar update package
	err := untar.UnTarUpdate()
	if err != nil {
		log.Println(err.Error())
		webJsonServe.ServeFailed(&u.Controller, 500, err.Error())
		return
	}
	err = checksum.CheckMD5s()
	if err != nil {
		log.Println(err.Error())
		webJsonServe.ServeFailed(&u.Controller, 500, err.Error())
		return
	}
	webJsonServe.ServeSuccess(&u.Controller, "success", "")
}