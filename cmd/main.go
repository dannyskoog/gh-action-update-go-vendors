package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/subosito/gotenv"
)

func main() {
	// Do stuff
	id := uuid.New()
	fmt.Printf("Random UUID: %s\n", id)

	// Do more stuff
	_ = gotenv.Load()
	fmt.Printf("Foo env variable: %s\n", os.Getenv("FOO"))
}
