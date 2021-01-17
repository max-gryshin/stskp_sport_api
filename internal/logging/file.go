package logging

import (
	"fmt"
	"time"

	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/setting"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	// FIXME: Potential BUG: Use filepath.Join() or path.Join() to build valid path
	return fmt.Sprintf("%s%s", setting.AppSetting.App.RuntimeRootPath, setting.AppSetting.App.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	// FIXME: Potential BUG: Use filepath.Join() or path.Join() to build valid path
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.App.LogSaveName,
		// TODO: This line has no since. this method will be called only on startup. And file name will be static until server restarts.
		// Use go-cronowriter to fix this issue.
		time.Now().Format(setting.AppSetting.App.TimeFormat),
		setting.AppSetting.App.LogFileExt,
	)
}
