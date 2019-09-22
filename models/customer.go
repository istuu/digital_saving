package models

import (
	"digital_saving/models/db"
	"digital_saving/structs"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strconv"
	"math/rand"
	"strings"
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
    err = c.Session.Find(nil).Sort("-first_name").All(&customers)
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
    customer.ID          = bson.NewObjectId()
	customer.CreatedAt   = time.Now()
	customer.ProductCode = "IDBSA001"
	customer.AccountNumber = RandomNumber()
	// Insert post to mongodb
	err = c.Session.Insert(&customer)
	if err != nil {
	   return customer, err
	}
	 return customer, err
}

// FindCustomer : Find customer by citizen_id
func FindCustomer(citizen_id string) (structs.Customer, error) {
	var (
	   err error
	   customer structs.Customer
	)
	
    // Get customer collection connection 
	c := newCustomerCollection()
 	defer c.Close()

	i, _ := strconv.Atoi(citizen_id)
    // get customer
    err = c.Session.Find(bson.M{"citizen_id": i}).One(&customer)
    if err != nil {
	   return customer,err
	}

   return customer, err
}

func FindAccount(AccountNumber string) (structs.AccountInfo, error) {
	var (
	   err error
	   customer structs.Customer
	   account structs.AccountInfo
	)
	
    // Get customer collection connection 
	c := newCustomerCollection()
 	defer c.Close()

	i, _ := strconv.Atoi(AccountNumber)
    // get customer
    err = c.Session.Find(bson.M{"account_number": i}).One(&customer)
    if err != nil {
	   return account,err
	}

	FullName := []string{customer.FirstName, customer.MiddleName, customer.LastName} 

	account.AccountNumber	= customer.AccountNumber 
	account.Name  			= strings.Join(FullName, " ") 
	account.ProductName		= "Digital Basics Savings"
	account.Interest 		= "3.00%" 
	account.AccountStatus	= "ACTIVE" 
	account.Branch  		= "DIGITAL BRANCH"
	account.Currency  		= "IDR" 
	account.OpeningDate     = customer.CreatedAt.Format("2 January 2006")

   return account, err
}

func RandomNumber() int {
	low := 10000000000
	hi  := 99999999999
    return low + rand.Intn(hi-low)
}