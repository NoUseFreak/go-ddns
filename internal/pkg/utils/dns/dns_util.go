package dns

import (
	"io/ioutil"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// DnsUtil is a collection of DNS related functions
type DnsUtil struct {
}

// GetMyIP finds the current external ip.
func (du DnsUtil) GetMyIP() (string, error) {
	url := "https://ifconfig.me/ip"

	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// NeedsUpdate checks if the domain matched the ip
func (du DnsUtil) NeedsUpdate(ip string, domain string) (bool, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.WithField("message", err.Error()).Warnf("Domain %s could not be checked\n", domain)
		return true, err
	}
	for _, curIP := range ips {
		if curIP.String() == ip {
			return false, nil
		}
	}

	return true, nil
}
