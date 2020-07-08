package main

type regex interface {
	match(str string) []string
}
