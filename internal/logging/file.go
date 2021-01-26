package logging

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"

	"path/filepath"
	"strings"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return filepath.Join(setting.AppSetting.App.RuntimeRootPath, setting.AppSetting.App.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	// TODO: This line has no since. this method will be called only on startup. And file name will be static until server restarts.
	// Use go-cronowriter to fix this issue.
	return strings.Join([]string{setting.AppSetting.App.LogSaveName, setting.AppSetting.App.LogFileExt}, ".")
}
