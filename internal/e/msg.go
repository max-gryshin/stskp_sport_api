package e

var MsgFlags = map[int]string{
	SUCCESS:                           "ok",
	ERROR:                             "fail",
	INVALID_PARAMS:                    "invalid params",
	ERROR_AUTH_CHECK_TOKEN_FAIL:       "Token Authentication failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:    "Token Timed out",
	ERROR_AUTH_TOKEN:                  "Token authentication failed",
	ERROR_AUTH:                        "Token error",
	ERROR_AUTH_CHECK_CREDENTIALS_FAIL: "Credentials authentication failed",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:      "Failed to save picture",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:     "Check image failed",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT:   "Check picture error, picture format or size problem",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
