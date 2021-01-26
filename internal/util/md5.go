package util

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	n, err := m.Write([]byte(value))
	if err != nil {
		logging.Error(err)
	}
	if n == 0 {
		logging.Error("no any bytes written")
	}

	return hex.EncodeToString(m.Sum(nil))
}
