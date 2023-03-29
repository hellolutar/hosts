package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type Hosts struct {
	Host string
	Addr string
}

func Query(filename string) []Hosts {
	log.Println("start...")
	start := time.Now()
	var hosts []Hosts
	//获取地址
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Println(err)
		return nil
	}

	for len(line) > 0 {
		ips, err := net.LookupHost(string(line))

		if err != nil {
			log.Println(err)
			return nil
		}

		for _, ip := range ips {
			hosts = append(hosts, Hosts{
				Host: ip,
				Addr: string(line),
			})
		}

		line, _, err = reader.ReadLine()
		if err != nil && err != io.EOF {
			log.Println(err)
			return nil
		}
	}
	log.Println("query cost time:", time.Now().UnixMicro()-start.UnixMicro())
	return hosts
}

func Write(filename string, hosts []Hosts) bool {
	start := time.Now()
	if len(hosts) == 0 {
		return false
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return false
	}
	defer f.Close()

	for _, host := range hosts {
		line := fmt.Sprintf("%v  %v\n", host.Host, host.Addr)
		log.Println(line)
		f.WriteString(line)
	}
	f.WriteString(fmt.Sprintf("Update time: %v\n", time.Now().Format("2006-01-02 15:04:05")))
	f.WriteString(fmt.Sprintf("cost time:%v \n", time.Now().UnixMicro()-start.UnixMicro()))

	log.Println("query cost time:", time.Now().UnixMicro()-start.UnixMicro())
	return true
}
