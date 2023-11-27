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
	prefix := switchGreetingLanguage(language)

	return prefix + name
}

func switchGreetingLanguage(language string) (prefix string) {
	switch language {
	case FRENSH:
		prefix = FRENCH_HELLO_PREFIX
	case SPANISH:
		prefix = SPANISH_HELLO_PREFIX
	default:
		prefix = ENGLISH_HELLO_PREFIX
	}
	return
}

