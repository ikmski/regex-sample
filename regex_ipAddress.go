package main

import (
	"regexp"
)

type ipAddress struct {
	regex *regexp.Regexp
}

func newIPAddress() *ipAddress {
	r := &ipAddress{
		regex: regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`),
	}
	return r
}

func (r *ipAddress) match(str string) []string {
	return r.regex.FindStringSubmatch(str)
}
