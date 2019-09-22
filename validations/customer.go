package validations

import (
	"digital_saving/models"
	"digital_saving/structs"
	"github.com/go-playground/locales/en"
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/universal-translator"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"strconv"
	
	"github.com/astaxie/beego"
)

var ErrorString []string

func CustomerValidator(Customer structs.Customer) []string{
	v := validator.New()
	ErrorString = nil

	translator := en.New()
	uni := ut.New(translator, translator)

	// this is usually known or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, found := uni.GetTranslator("en")

	if !found {
		beego.Warning("translator not found")
	}

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		beego.Warning(err)
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	err := v.Struct(Customer)
	if(err != nil){
		for _, e := range err.(validator.ValidationErrors) {
			ErrorString = append(ErrorString, e.Translate(trans))
		}
	}
	i := strconv.Itoa(Customer.CitizenId)
	cus, err := models.FindCustomer(i)

	beego.Warning(cus)
	if(err == nil){
		ErrorString = append(ErrorString, "Citizen ID is already exist.")
	}

	return ErrorString

}