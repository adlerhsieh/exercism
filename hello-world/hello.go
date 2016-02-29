package hello

import "fmt"

const testVersion = 2

// Bring in a name as the argument to say hello to this person.
//
// ```
// HelloWorld()        // => "Hello, World!"
// HelloWorld("Baker") // => "Hello, Baker!"
// ```
func HelloWorld(name string) (sentence string) {
	if name == "" {
		sentence = "Hello, World!"
	} else {
		sentence = fmt.Sprintf("Hello, %s!", name)
	}
	return
}
