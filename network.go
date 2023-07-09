package mob

import (
	"fmt"
	"net"
	"sort"
	"time"

	"go.olapie.com/mob/nomobile"
	"go.olapie.com/utils"
)

const (
	NoNetwork = 0
	Cellular  = 1
	WIFI      = 2
)

const (
	Idle       = 0
	Connecting = 1
	Connected  = 2
)

func GetOutboundIP() string {
	return utils.GetOutboundIPString()
}

const (
	DNSGoogle1    = "8.8.8.8"
	DNSGoogle2    = "8.8.8.8"
	DNSCloudflare = "1.1.1.1"
	DNS114A       = "114.114.114.114"
	DNS114B       = "114.114.115.115"
	DNSAlibaba1   = "223.5.5.5"
	DNSAlibaba2   = "223.6.6.6"
	DNSBaidu      = "180.76.76.76"
)

var cnDNSList = []string{DNS114A, DNSAlibaba1, DNSBaidu, DNS114B, DNSAlibaba2}
var otherDNSList = []string{DNSGoogle1, DNSCloudflare, DNSGoogle2}

func IsNetworkReachable() bool {
	return checkNetwork(cnDNSList...) || checkNetwork(otherDNSList...)
}

func checkNetwork(ips ...string) bool {
	for _, ip := range ips {
		conn, err := net.DialTimeout("tcp", ip+":53", time.Second*2)
		if err == nil {
			conn.Close()
			return true
		}
	}
	return false
}

type IFace struct {
	Name         string
	MTU          int
	Flags        string
	HardwareAddr string
	Addrs        *StringList
}

type IFaceList struct {
	nomobile.List[*IFace]
}

func (l *IFaceList) Less(i, j int) bool {
	return l.Get(i).Name < l.Get(j).Name
}

func (l *IFaceList) Swap(i, j int) {
	a := l.List.Elements()
	a[i], a[j] = a[j], a[i]
}

func (l *IFaceList) Sort() {
	sort.Sort(l)
}

func ListIFaces() *IFaceList {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("No network interface")
		return nil
	}

	res := new(IFaceList)
	for _, i := range ifaces {
		iface := &IFace{
			Name:         i.Name,
			MTU:          i.MTU,
			Flags:        i.Flags.String(),
			HardwareAddr: i.HardwareAddr.String(),
			Addrs:        NewStringList(),
		}

		if addrs, err := i.Addrs(); err == nil {
			for _, addr := range addrs {
				iface.Addrs.Add(addr.String())
			}
		}

		if addrs, err := i.MulticastAddrs(); err == nil {
			for _, addr := range addrs {
				iface.Addrs.Add(addr.String())
			}
		}

		res.Add(iface)
	}

	return res
}
