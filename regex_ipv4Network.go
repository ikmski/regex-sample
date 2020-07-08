package main

import (
	"regexp"
)

type ipv4Network struct {
	r *regexp.Regexp
}

func newIPv4Network() regex {
	obj := &ipv4Network{
		r: regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\/\d{1,2}$`),
	}
	return obj
}

func (obj *ipv4Network) match(str string) []string {
	return obj.r.FindStringSubmatch(str)
}
