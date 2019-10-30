package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	ipHosts := flag.String("h", "", "print all IPs of a host")
	ipFlag := flag.String("ip", "", "print all hosts of an IP")
	interfaceFlag := flag.Bool("i", false, "print all interfaces and IPs")
	flag.Parse()

	if *interfaceFlag {
		printInterfaces()
	}

	if *ipFlag != "" {
		printCIDR(ipFlag)
		printHosts(ipFlag)
	}

	if *ipHosts != "" {
		printIps(ipHosts)
	}
}

func printCIDR(ipFlag *string) {
	address := net.ParseIP(*ipFlag)
	fmt.Println(address)
	mask := net.CIDRMask(24, 32)
	network := address.Mask(mask)
	fmt.Println(mask, network)

}

func printIps(ipHosts *string) {
	ips, _ := net.LookupHost(*ipHosts)
	fmt.Println(ips)
}

func printHosts(ipFlag *string) {
	hosts, _ := net.LookupAddr(*ipFlag)
	fmt.Println(hosts)
}

func printInterfaces() {
	interfaces, _ := net.Interfaces()

	for _, i := range interfaces {
		iName, _ := net.InterfaceByName(i.Name)
		addresses, _ := iName.Addrs()
		fmt.Println(iName.Name, addresses)
	}
}
