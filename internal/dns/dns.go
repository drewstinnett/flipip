package dns

import (
	"errors"
	"net"
	"strings"
)

// Figure out if this is v4 or v6, and run appropriate reverser
func Reverse(target string) (string, error) {
	ipv, err := checkIPAddressType(target)
	if err != nil {
		return "", err
	}

	var f func(string) (string, error)

	if ipv == "6" {
		f = reverse6
	} else if ipv == "4" {
		f = reverse4
	} else {
		return "", errors.New("Unknown IP version")
	}
	res, err := f(target)
	if err != nil {
		return "", err
	}
	return res, nil
}

func reverse4(target string) (string, error) {
	s := strings.Split(target, ".")
	var sb strings.Builder
	for i := range s {
		n := s[len(s)-1-i]
		sb.WriteString(n)
		sb.WriteString(".")
	}
	return sb.String() + "in-addr.arpa.", nil
}

const hexDigit = "0123456789abcdef"

// Mostly from the core library
func reverse6(target string) (string, error) {
	ip := net.ParseIP(target)
	buf := make([]byte, 0, len(ip)*4+len("ip6.arpa."))
	// Add it, in reverse, to the buffer
	for i := len(ip) - 1; i >= 0; i-- {
		v := ip[i]
		buf = append(buf, hexDigit[v&0xF],
			'.',
			hexDigit[v>>4],
			'.')
	}
	buf = append(buf, "ip6.arpa."...)
	ret := string(buf)
	// Errors?
	if ret == "ip6.arpa." {
		return "", errors.New("IPV6 Reverse failed")
	}
	return string(buf), nil
}

func checkIPAddressType(ip string) (string, error) {
	if net.ParseIP(ip) == nil {
		return "", errors.New("Invalid IP")
	}
	var v string
	for i := 0; i < len(ip); i++ {
		switch ip[i] {
		case '.':
			v = "4"
		case ':':
			v = "6"
		}
	}
	return v, nil
}
