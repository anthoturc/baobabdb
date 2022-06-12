package model

import "net"

const GoogDns = "8.8.8.8:80"

func OutboundIp() (net.IP, error) {
	conn, err := net.Dial("udp", GoogDns)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP, nil
}
