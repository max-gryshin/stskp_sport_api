package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error, returnErrors bool) string {
	var result string
	for _, err := range errors {
		if returnErrors {
			result += err.Key + " " + err.Message + " "
		}
		logging.Error(err.Key, err.Message)
	}

	return result
}
