package comm

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
)

//IsIPV4 判断一个IP地址是否为IPV4,不能带掩码
func IsIPV4(ip string) bool {
	if net.ParseIP(ip).To4() == nil {
		return false
	}
	return true
}

// IPInSegmentOne 判断指定IP或IP段 是否在某个IP段范围内，是：返回true，否：返回false （"192.168.10.5", "192.168.10.5/32"）
//注意：使用前要判断iprange格式的合法性如; 192.168.10.5/32
func IPInSegmentOne(ip, iprange string) bool {
	ipb := ip2binary(ip)
	ipr := strings.Split(iprange, "/")
	masklen, err := strconv.ParseUint(ipr[1], 10, 32)
	if err != nil {
		fmt.Println(err)
		return false
	}
	iprb := ip2binary(ipr[0])
	return strings.EqualFold(ipb[0:masklen], iprb[0:masklen])
}

// IPInSegmentTwo 判断指定IP是否在指定两个IP地址区间内，是：返回true，否：返回false （"192.168.10.1","192.168.9.5", "192.168.10.5"）
func IPInSegmentTwo(ip, ipStart, ipEnd string) bool {
	trial := net.ParseIP(ip)
	start := net.ParseIP(ipStart)
	end := net.ParseIP(ipEnd)
	if trial.To4() == nil {
		fmt.Printf("%v is not an IPv4 address\n", trial)
		return false
	}
	if bytes.Compare(trial, start) >= 0 && bytes.Compare(trial, end) <= 0 {
		fmt.Printf("%v is between %v and %v\n", trial, start, end)
		return true
	}
	fmt.Printf("%v is NOT between %v and %v\n", trial, start, end)
	return false
}

//将IP地址转化为二进制String
func ip2binary(ip string) string {
	str := strings.Split(ip, ".")
	var ipstr string
	for _, s := range str {
		i, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		ipstr = ipstr + fmt.Sprintf("%08b", i)
	}
	return ipstr
}
