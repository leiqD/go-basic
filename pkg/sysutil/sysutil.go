package sysutil

import "net"

// GetLocalIP 获取本机IP地址，多个IP则用分号分隔
func GetLocalIP() (ipv4s []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipv4s = append(ipv4s, ipnet.IP.String())
			}
		}
	}
	return ipv4s
}
