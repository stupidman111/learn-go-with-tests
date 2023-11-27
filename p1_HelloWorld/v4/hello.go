package main

const ENGLISH_HELLO_PREFIX = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return ENGLISH_HELLO_PREFIX + name
}

