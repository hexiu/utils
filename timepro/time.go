package timepro

import "time"
import "standAlone/utils/log"

// StringToTime 序列化字符串时间
func StringToTime(t string) time.Time {
	withNanos := "2006-01-02 15:04:36"
	tm, err := time.Parse(withNanos[:len(t)], t)
	if err != nil {
		log.CheckErr(err, log.LogHigh)
		return time.Now()
	}
	return tm
}
