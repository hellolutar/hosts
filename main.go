package main

import (
	"log"
	"net"
)

func main() {
	updateHosts()
}

func updateHosts() {
	Write("hosts_github", Query("github.txt"))
	//Write("hosts_google", Query("google.txt"))
}

func lookupCname() {
	cname, err := net.LookupCNAME("www.a.shifen.com")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("lookupCname:", cname)
}
func lookupAddr() {
	names, err := net.LookupAddr("140.82.113.4")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("lookupAddr:", names)
}

func lookupHost() {
	addrs, err := net.LookupHost("github.com")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("lookupHost:", addrs)
}
