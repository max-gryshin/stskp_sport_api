package logging

import (
	"fmt"
	"time"

	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/setting"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.App.RuntimeRootPath, setting.AppSetting.App.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.App.LogSaveName,
		time.Now().Format(setting.AppSetting.App.TimeFormat),
		setting.AppSetting.App.LogFileExt,
	)
}
