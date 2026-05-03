package main

import (
	"fmt"
	"strings"
)

type Languaje struct {
	lang     string
	greeting string
	world    string
}

var greetingLangs = [...]Languaje{
	{"english", "Hello ", "World!"},
	{"spanish", "Hola ", "Mundo!"},
	{"french", "Bonjour ", "Monde!"},
	{"portuguese", "Olá ", "Mundo!"},
}

func SayHello(name string, lang string) string {
	lang = strings.ToLower(lang)
	defaultLang := greetingLangs[0]

	for index, val := range greetingLangs {
		if val.lang == lang {
			defaultLang = greetingLangs[index]
		}
	}

	if name == "" {
		name = defaultLang.world
	}

	return defaultLang.greeting + name
}

func main() {
	fmt.Println(SayHello("", "english"))
}
