package validation

import (
	"github.com/astaxie/beego/validation"
)

func validationAny(json interface{}) bool {
	valid := validation.Validation{}
	check, err := valid.Valid(json)

	if err != nil {
		return false
	}

	return check
}
