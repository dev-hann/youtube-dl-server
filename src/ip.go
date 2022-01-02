package src

import (
	"log"
	"net"
	"strings"
)

func MyIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	local := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(local.String(), ":")[0]
	return ip
}
