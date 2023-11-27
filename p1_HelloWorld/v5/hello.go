package main

const SPANISH = "Spanish"
const SPANISH_HELLO_PREFIX = "Hola, "

const FRENSH = "French"
const FRENCH_HELLO_PREFIX = "Bonjour, "

const ENGLISH_HELLO_PREFIX = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	//default language is English
	prefix := ENGLISH_HELLO_PREFIX

	switch language {
	case SPANISH:
		prefix = SPANISH_HELLO_PREFIX
	case FRENSH:
		prefix = FRENCH_HELLO_PREFIX
	}

	return prefix + name
}

