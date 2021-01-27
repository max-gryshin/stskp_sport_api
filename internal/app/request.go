package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
}
