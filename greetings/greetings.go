package greetings

import "fmt"

func HelloGreeting(name string) string {
	return fmt.Sprintf("Hello, %s.", name)
}
