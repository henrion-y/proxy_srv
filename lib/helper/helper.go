package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Date(format string, timestamp ...int64) string {
	var ts = time.Now().Unix()
	if len(timestamp) > 0 {
		ts = timestamp[0]
	}
	var t = time.Unix(ts, 0)
	Y := strconv.Itoa(t.Year())
	m := fmt.Sprintf("%02d", t.Month())
	d := fmt.Sprintf("%02d", t.Day())
	H := fmt.Sprintf("%02d", t.Hour())
	i := fmt.Sprintf("%02d", t.Minute())
	s := fmt.Sprintf("%02d", t.Second())

	format = strings.Replace(format, "Y", Y, -1)
	format = strings.Replace(format, "m", m, -1)
	format = strings.Replace(format, "d", d, -1)
	format = strings.Replace(format, "H", H, -1)
	format = strings.Replace(format, "i", i, -1)
	format = strings.Replace(format, "s", s, -1)
	if format == "1-01-01 08:00:00" {
		return ""
	}
	return format
}

func Success() (int, interface{}) {
	return http.StatusOK, gin.H{
		"status": gin.H{
			"code":             0,
			"message":          "success",
			"time":             Date("Y-m-d H:i:s", time.Now().Unix()),
			"accessTokenState": "keep",
		},
		"data": gin.H{},
	}
}

func SuccessWithData(data interface{}) (int, interface{}) {
	return http.StatusOK, gin.H{
		"status": gin.H{
			"code":             0,
			"message":          "success",
			"time":             Date("Y-m-d H:i:s", time.Now().Unix()),
			"accessTokenState": "keep",
		},
		"data": data,
	}
}

func SuccessWithDataList(datalist interface{}) (int, interface{}) {
	return http.StatusOK, gin.H{
		"status": gin.H{
			"code":             0,
			"message":          "success",
			"time":             Date("Y-m-d H:i:s", time.Now().Unix()),
			"accessTokenState": "keep",
		},
		"data": gin.H{
			"dataList": datalist,
		},
	}
}

func Fail(message ...string) (int, interface{}) {
	var msg string

	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = "服务内部错误"
	}

	return http.StatusInternalServerError, gin.H{
		"status": gin.H{
			"code":             -1,
			"message":          msg,
			"time":             Date("Y-m-d H:i:s", time.Now().Unix()),
			"accessTokenState": "keep",
		},
		"data": gin.H{},
	}
}
