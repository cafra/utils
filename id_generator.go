package utils

import (
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/zheng-ji/goSnowFlake"
)

var (
	idGenerator, _ = goSnowFlake.NewIdWorker(Ip2Int(GetInternalIp()))
)

func NextId() int64 {
	ts, err := idGenerator.NextId()
	if err != nil {
		log.Printf("IDGenerator.NextId err=%v", err)
	}
	return ts
}

func GetInternalIp() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("GetInternalIp err", err)
		return
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return
}

// 思路: 4段数字,最大值255,4部分数据和<1024
func Ip2Int(ip string) (workerid int64) {
	for _, s := range strings.Split(ip, ".") {
		sub, _ := strconv.ParseInt(s, 10, 64)
		workerid += sub
	}
	return
}
