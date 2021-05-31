package helper

import "strings"

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"` //interface berarti tipe data dia dinamis, kadang bisa string, bisa berupa array dan ga harus semuanya string
	Data    interface{} `json:"data"`
}

//EmptyObj is used for when data doesn't want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil, //dikasih null karena dia buildresponse bukan response
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic success response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n") //karena error nya dalam satu json akan ada banyak dan dipisahkan oleh \n maka di split dengan splitter \n
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}