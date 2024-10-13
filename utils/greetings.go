package utils

import "fmt"

// Greet returns a greeting message
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// this function is not exported (private to the package)
func internalHelper() {
	// ...
}
