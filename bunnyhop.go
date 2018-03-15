package bunnyhop

import (
	"net"
	"time"

	"github.com/mdlayher/arp"
)

// IPInUse defines a generic function to check to see if an IP is in use on a given network.
func IPInUse(iface, ip string, timeout time.Duration) (bool, error) {
	ifi, err := net.InterfaceByName(iface)
	if err != nil {
		return false, err
	}

	c, err := arp.Dial(ifi)
	if err != nil {
		return false, err
	}

	defer c.Close()

	if err := c.SetDeadline(time.Now().Add(timeout)); err != nil {
		return false, err
	}

	ipAddr := net.ParseIP(ip).To4()
	_, err = c.Resolve(ipAddr)

	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
