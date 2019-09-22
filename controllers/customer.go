package controllers

import (
	"digital_saving/models"
	"digital_saving/structs"
	"digital_saving/validations"
	"encoding/json"

	"github.com/astaxie/beego"
)

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
func (c *CustomerController) Post() {
	var (
		cus structs.Customer
		JSON structs.ReturnJson
	)

	json.Unmarshal(c.Ctx.Input.RequestBody, &cus)
	err := validations.CustomerValidator(cus)
	// returns nil or ValidationErrors ( []FieldError )
	// err := validate.Struct(cus)
	JSON.Success = false	
	JSON.Data    = nil	

	if err != nil {
		JSON.Error   = err
	}else{
		customer, err  := models.CreateCustomer(cus)
		if err != nil {
			JSON.Message   = err.Error()
		}else{
			JSON.Success = true
			JSON.Data    = &customer
		}
	}

	c.Data["json"] = JSON
	c.ServeJSON()
}

// @Title Get
// @Description find object by citizenId
// @Param	citizenId		path 	string	true		"the citizenId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :citizenId is empty
// @router /:citizenId [get]
func (c *CustomerController) Get() {
	var JSON structs.ReturnJson

	citizenId := c.Ctx.Input.Param(":citizenId")
	JSON.Success = false	
	JSON.Data    = nil	
	if citizenId != "" {
		customer, err := models.FindCustomer(citizenId)
		if err != nil {
			JSON.Message   = err.Error()
		}else{
			JSON.Success = true
			JSON.Data    = &customer
		}
	}
	c.Data["json"] = JSON
	c.ServeJSON()
}

// @Title GetAccount
// @Description find object by accountNumber
// @Param	accountNumber		path 	string	true		"the accountNumber you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :accountNumber is empty
// @router /account/:accountNumber [get]
func (c *CustomerController) GetAccount() {
	var JSON structs.AccountJson

	accountNumber := c.Ctx.Input.Param(":accountNumber")
	JSON.Success = false	
	JSON.Data    = nil	
	if accountNumber != "" {
		account, err := models.FindAccount(accountNumber)
		if err != nil {
			JSON.Message   = err.Error()
		}else{
			JSON.Success = true
			JSON.Data    = &account
		}
	}
	c.Data["json"] = JSON
	c.ServeJSON()
}