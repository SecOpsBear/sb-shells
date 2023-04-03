package validators

import (
	"net"
	"strconv"
)

// returns true if the ip address is a valid format ddd.ddd.ddd.ddd
func IsIPaddressValid(ip string) bool {
	if parsedIP := net.ParseIP(ip); parsedIP == nil {
		return false
	}
	return true
}

// returns true if the port number is between 1 and 65535
func IsIPaddressPortValid(port string) bool {

	ppt, err := strconv.Atoi(port)
	if err != nil {

		return false
	}
	if (ppt >= 1) && (ppt <= 65535) {
		return true
	}
	return false

}
