package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["digital_saving/controllers:CustomerController"] = append(beego.GlobalControllerRouter["digital_saving/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["digital_saving/controllers:CustomerController"] = append(beego.GlobalControllerRouter["digital_saving/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["digital_saving/controllers:CustomerController"] = append(beego.GlobalControllerRouter["digital_saving/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:citizenId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["digital_saving/controllers:CustomerController"] = append(beego.GlobalControllerRouter["digital_saving/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "GetAccount",
            Router: `/account/:accountNumber`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
