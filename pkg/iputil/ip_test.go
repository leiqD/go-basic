package iputil

import "testing"

func TestIsInnerIp(t *testing.T) {
	//回环地址
	if !IsInnerIp("127.0.0.1") {
		t.Fatal("127.0.0.1")
	}

	//A类地址
	if !IsInnerIp("10.0.0.0") {
		t.Fatal("10.0.0.0")
	}
	if !IsInnerIp("10.100.100.100") {
		t.Fatal("10.100.100.100")
	}
	if !IsInnerIp("10.255.255.255") {
		t.Fatal("10.255.255.255")
	}
	if IsInnerIp("9.255.255.255") {
		t.Fatal("9.255.255.255")
	}
	if IsInnerIp("11.0.0.0") {
		t.Fatal("11.0.0.0")
	}

	//B类地址
	if !IsInnerIp("172.16.0.0") {
		t.Fatal("172.16.0.0")
	}
	if !IsInnerIp("172.20.100.100") {
		t.Fatal("172.20.100.100")
	}
	if !IsInnerIp("172.31.255.255") {
		t.Fatal("172.31.255.255")
	}
	if IsInnerIp("172.15.255.255") {
		t.Fatal("172.15.255.255")
	}
	if IsInnerIp("172.32.0.0") {
		t.Fatal("172.32.0.0")
	}

	//C类地址
	if !IsInnerIp("192.168.0.0") {
		t.Fatal("192.168.0.0")
	}
	if !IsInnerIp("192.168.100.100") {
		t.Fatal("192.168.100.100")
	}
	if !IsInnerIp("192.168.255.255") {
		t.Fatal("192.168.255.255")
	}
	if IsInnerIp("192.167.255.255") {
		t.Fatal("192.167.255.255")
	}
	if IsInnerIp("192.169.0.0") {
		t.Fatal("192.169.0.0")
	}
}
