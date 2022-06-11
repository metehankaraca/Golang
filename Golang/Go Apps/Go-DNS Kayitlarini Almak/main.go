package main

import (
	"fmt"
	"net"
)

func main() {
	var siteurl string
	fmt.Print("Web sitesi adresinizi giriniz (Orn: google.com): ")
	fmt.Scan(&siteurl)

	nameserverlist := getNameServerList(siteurl)
	if nameserverlist == nil {
		fmt.Println("Name Server Listesi cekilirken hata olustu!")
	} else {
		fmt.Println("Name Server Listesi yazdiriliyor...")
		for i := 0; i < len(nameserverlist); i++ {
			fmt.Println(nameserverlist[i])
		}
	}

	iplist := getIpList(siteurl)
	if iplist == nil {
		fmt.Println("IP Listesi cekilirken hata olustu!")
	} else {
		fmt.Println("IP Listesi yazdiriliyor...")
		for i := 0; i < len(iplist); i++ {
			fmt.Println(iplist[i])
		}
	}

	txtList := getTxtList(siteurl)
	if iplist == nil {
		fmt.Println("TXT Listesi cekilirken hata olustu!")
	} else {
		fmt.Println("TXT Listesi yazdiriliyor...")
		for i := 0; i < len(txtList); i++ {
			fmt.Println(txtList[i])
		}
	}

	fmt.Println("CNAME Degeri yazdiriliyor...")
	fmt.Println(getCname(siteurl))
}

func getNameServerList(link string) []string {
	var servers []string
	nameserver, _ := net.LookupNS(link)
	for _, ns := range nameserver {
		servers = append(servers, ns.Host)
	}
	return servers
}

func getIpList(link string) []string {
	var ips []string

	iplist, _ := net.LookupIP(link)
	for _, ip := range iplist {
		ips = append(ips, ip.String())
	}
	return ips
}

func getCname(link string) string {
	cname, _ := net.LookupCNAME(link)
	return cname
}

func getTxtList(link string) []string {
	var txts []string
	txtList, _ := net.LookupTXT(link)
	for _, txt := range txtList {
		txts = append(txts, txt)
	}
	return txts
}
