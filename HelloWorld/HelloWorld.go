package main

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "

func helloWorld(prefix string, name string) string {

	if prefix == "Spanish" {
		return spanishPrefix + name
	}
	if prefix == "English" {return englishPrefix + name}
	return "no existent prefix"
}
