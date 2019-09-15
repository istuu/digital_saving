package models

import (
	"digital_saving/models/db"
	"digital_saving/structs"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func newCustomerCollection() *db.Collection {
	return db.NewCollectionSession("customers")
}

// GetCustomers : Get list all data customers
func GetCustomers() ([]structs.Customer, error) {
	var (
	   err error
	   customers []structs.Customer
    )
    // Get customer collection connection 
	c := newCustomerCollection()
	defer c.Close()
    // get posts
    err = c.Session.Find(nil).Sort("-published_at").All(&customers)
    if err != nil {
	   return customers,err
	}
   return customers, err
}

// CreateCustomer : Create customer
func CreateCustomer(customer structs.Customer) (structs.Customer, error) {
	var (
	   err error
	)
    // Get post collection connection 
	c := newCustomerCollection()
	defer c.Close()
	
    // set default mongodb ID  and created date
    customer.ID = bson.NewObjectId()
    customer.CreatedAt = time.Now()
	// Insert post to mongodb
	err = c.Session.Insert(&customer)
	if err != nil {
	   return customer, err
	}
	 return customer, err
}

// FindCustomer : Find customer by id
func FindCustomer(id bson.ObjectId) (structs.Customer, error) {
	var (
	   err error
	   customer structs.Customer
   )
   // Get customer collection connection 
	c := newCustomerCollection()
	defer c.Close()
   // get customer
   err = c.Session.FindId(id).One(&customer)
   if err != nil {
	   return customer,err
	}
   return customer, err
}