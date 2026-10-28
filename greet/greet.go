// Package greet builds greeting messages.
package greet

import "fmt"

// Hello returns a greeting for name, defaulting to "World" when name is empty.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
