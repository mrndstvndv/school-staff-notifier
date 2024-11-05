package utils

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var debugLogs = true

func GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

func CurrentTimeUnixString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

/*
Credits to: https://github.com/Kethsar/ytarchive/blob/dev/util.go
*/

// TODO: add an error log, then we remove all fatal logs in the application
func LogDebug(format string, args ...interface{}) {
	if debugLogs {
		msg := format
		if len(args) > 0 {
			msg = fmt.Sprintf(format, args...)
		}
		log.Printf("DEBUG: \033[36m%s\033[0m\033[K\n", msg)
	}
}
