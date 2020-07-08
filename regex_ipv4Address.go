package main

import (
	"regexp"
)

type ipv4Address struct {
	r *regexp.Regexp
}

func newIPv4Address() regex {
	obj := &ipv4Address{
		r: regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`),
	}
	return obj
}

func (obj *ipv4Address) match(str string) []string {
	return obj.r.FindStringSubmatch(str)
}
