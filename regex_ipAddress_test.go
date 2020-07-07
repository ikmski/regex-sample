package main

import (
	"fmt"
	"testing"
)

func TestIPAddress(t *testing.T) {

	r := newIPAddress()

	{
		str := "0.0.0.0"
		match := r.match(str)
		fmt.Println(match)
		if len(match) != 1 {
			t.Errorf("got %v\nwant %v", len(match), 1)
		}
		if match[0] != str {
			t.Errorf("got %v\nwant %v", match[0], str)
		}
	}
	{
		str := "255.255.255.255"
		match := r.match(str)
		fmt.Println(match)
		if len(match) != 1 {
			t.Errorf("got %v\nwant %v", len(match), 1)
		}
		if match[0] != str {
			t.Errorf("got %v\nwant %v", match[0], str)
		}
	}
	{
		str := "192.168.1.10"
		match := r.match(str)
		fmt.Println(match)
		if len(match) != 1 {
			t.Errorf("got %v\nwant %v", len(match), 1)
		}
		if match[0] != str {
			t.Errorf("got %v\nwant %v", match[0], str)
		}
	}
	{
		str := "192.168.1.0/24"
		match := r.match(str)
		fmt.Println(match)
		if len(match) != 0 {
			t.Errorf("got %v\nwant %v", len(match), 1)
		}
	}
}
