package core_error

import core_constants "github.com/BM-laoli/go-gin-example/src/core/constants"

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := core_constants.MsgFlags[code]
	if ok {
		return msg
	}

	return core_constants.MsgFlags[core_constants.ERROR]
}
