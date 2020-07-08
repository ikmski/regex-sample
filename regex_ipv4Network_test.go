package main

import (
	"fmt"
	"testing"
)

func TestIPv4Network(t *testing.T) {

	r := newIPv4Network()

	{
		str := "0.0.0.0/0"
		match := r.match(str)
		fmt.Println(match)
		if len(match) != 1 {
			t.Errorf("got %v\nwant %v", len(match), 1)
		} else if match[0] != str {
			t.Errorf("got %v\nwant %v", match[0], str)
		}
	}
	{
		str := "255.255.255.255/32"
		match := r.match(str)
		fmt.Println(match)
		if len(match) != 1 {
			t.Errorf("got %v\nwant %v", len(match), 1)
		} else if match[0] != str {
			t.Errorf("got %v\nwant %v", match[0], str)
		}
	}
	{
		str := "192.168.1.10/24"
		match := r.match(str)
		fmt.Println(match)
		if len(match) != 1 {
			t.Errorf("got %v\nwant %v", len(match), 1)
		} else if match[0] != str {
			t.Errorf("got %v\nwant %v", match[0], str)
		}
	}
	{
		str := "192.168.1.0"
		match := r.match(str)
		fmt.Println(match)
		if len(match) != 0 {
			t.Errorf("got %v\nwant %v", len(match), 1)
		}
	}
}
