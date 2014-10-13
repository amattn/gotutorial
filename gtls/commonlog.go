package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// http://en.wikipedia.org/wiki/Common_Log_Format
// example 127.0.0.1 user-identifier frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326
func CommonLogFormat(req *http.Request, statusCode, contentLength int) string {

	user := "-"
	if req.URL.User != nil {
		user = req.URL.User.Username()
	}
	common_log_format_parts := []string{
		req.RemoteAddr,
		"-",
		user,
		time.Now().Format("[02/Jan/2006:15:04:05 -0700]"),
		`"` + req.Method,
		req.URL.RequestURI(),
		req.Proto + `"`,
		strconv.FormatInt(int64(statusCode), 10),
		strconv.FormatInt(int64(contentLength), 10),
		"\n",
	}
	return fmt.Sprint(strings.Join(common_log_format_parts, " "))
}
