package utils

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			// This will return the first non-loopback IP address (e.g., the IP assigned to your hotspot)
			return ipNet.IP.String(), nil
		}
	}
	return "", fmt.Errorf("no active IP address found")
}


func CurrentTimeUnixString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
