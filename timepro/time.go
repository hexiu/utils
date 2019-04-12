package timepro

import (
	"time"

	"github.com/hexiu/utils/log"
)

// StringToTime 序列化字符串时间
func StringToTime(t string) time.Time {
	withNanos := "2006-01-02 15:04:05"
	// var cstZone = time.FixedZone("CST", 8*3600)
	cstZone, err := time.LoadLocation("Local")
	if err != nil {
		return time.Now()
	}
	tm, err := time.ParseInLocation(withNanos[:len(t)], t, cstZone)
	if err != nil {
		log.CheckErr(err, log.LogHigh)
		return time.Now()
	}
	return tm.Local()
}
