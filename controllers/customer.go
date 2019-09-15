package controllers

import (
	"digital_saving/models"
	"digital_saving/structs"
	// "encoding/json"

	"github.com/astaxie/beego"
)

type Customer structs.Customer

// Operations about customer
type CustomerController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Customers
// @Success 200 {Customer} structs.Customer
// @Failure 403 :CustomerId is empty
// @router / [get]
func (c *CustomerController) GetAll() {
	customer, _ := models.GetCustomers()
	c.Data["json"] = customer
	c.ServeJSON()
}

// @Title Create
// @Description create customer
// @Param	body		body 	structs.Customer	true		"The customer content"
// @Success 200 {Customer} structs.Customer
// @Failure 403 body is empty
// @router / [post]
// func (c *CustomerController) Post() {
// 	var cus Customer
// 	json.Unmarshal(c.Ctx.Input.RequestBody, &cus)
// 	customer, _ := models.CreateCustomer(cus)

// 	c.Data["json"] = customer
// 	c.ServeJSON()
// }