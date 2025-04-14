package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

type Italian struct {}

type Portuguese struct {}

func (v Italian) LanguageName() string {
	return "Italian"
}

func (v Portuguese) LanguageName() string {
	return "Portuguese"
}

func (v Italian) Greet(visitorName string) string {
	return "Ciao " + visitorName + "!"
}

func (v Portuguese) Greet(visitorName string) string {
	return "Olá " + visitorName + "!"
}

func SayHello(visitorName string, visitor Greeter) string {
	return "I can speak " + visitor.LanguageName() + ": " + visitor.Greet(visitorName)
}