package helper

import validation "github.com/aminesmkhani/go-clean/api/validations"


type BaseHttpResponse struct{
	Result any `json:"result"`
	Success bool `json:"success"`
	ResultCode int `json:"resultCode"`
	ValidationErrors *[]validation.ValidationError `json:"validationErrors"`
	Error 		any `json:"error"`
}