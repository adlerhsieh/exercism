package main

import (
	godo "gopkg.in/godo.v2"
)

// Tasks setups tasks for Godo
func Tasks(p *godo.Project) {
	p.Task("default", godo.S{"test"}, nil)

	p.Task("test", nil, func(c *godo.Context) {
		if c.FileEvent != nil {
			c.Run("go test " + GoFiles(c.FileEvent.Path))
		}
	}).Src("**/*.go")
}

func main() {
	godo.Godo(Tasks)
}
