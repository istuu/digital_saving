package structs

type ReturnJson struct{
	Success bool `json:"success" bson:"success"`
	Message string `json:"message,omitempty" bson:"message"`
	Error []string `json:"error,omitempty" bson:"error"`
	Data *Customer `json:"data,omitempty" bson:"data"`
}