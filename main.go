package main

import "github.com/graniticio/granitic/v2"
import "v2-demo-app/bindings"

func main() {
	granitic.StartGranitic(bindings.Components())
}
