package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["app/controllers:TaskController"] = append(beego.GlobalControllerRouter["app/controllers:TaskController"],
        beego.ControllerComments{
            Method: "RunTask",
            Router: `/run`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["app/controllers:UpdateController"] = append(beego.GlobalControllerRouter["app/controllers:UpdateController"],
        beego.ControllerComments{
            Method: "AcquireUpdate",
            Router: `/package`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["app/controllers:UpdateController"] = append(beego.GlobalControllerRouter["app/controllers:UpdateController"],
        beego.ControllerComments{
            Method: "Start",
            Router: `/start`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
