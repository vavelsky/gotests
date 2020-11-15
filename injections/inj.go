package injections

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Greet(os.Stdout, "Elodie")
}

// Greet welcomes
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
