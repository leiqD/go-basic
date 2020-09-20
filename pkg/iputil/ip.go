package iputil

import (
	"net"
	"strconv"
	"strings"
)

var (
	innerIpA = inetAton("10.255.255.255")
	innerIpB = inetAton("172.16.255.255")
	innerIpC = inetAton("192.168.255.255")
	innerIpD = inetAton("127.255.255.255")
)

func checkIp(ipStr string) bool {
	address := net.ParseIP(ipStr)
	if address == nil {
		//fmt.Println("ip地址格式不正确")
		return false
	} else {
		//fmt.Println("正确的ip地址", address.String())
		return true
	}
}

// ip to int64
func inetAton(ipStr string) int64 {
	bits := strings.Split(ipStr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

//判断内网IP地址
func IsInnerIp(ipStr string) bool {
	if !checkIp(ipStr) {
		return false
	}

	inputIpNum := inetAton(ipStr)
	return inputIpNum>>24 == innerIpA>>24 || inputIpNum>>20 == innerIpB>>20 ||
		inputIpNum>>16 == innerIpC>>16 || inputIpNum>>24 == innerIpD>>24
}
