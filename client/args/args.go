package args

import "os"
import "fmt"

var arguments []string

func init() {
	arguments = os.Args

	if len(arguments) < 3 {
		fmt.Println("usage: go run main.go joaquim 127.0.0.1:3000")
		os.Exit(1)
	}
}

// Username devolve o nome do usuario passado como command line argument.
func Username() string {
	return arguments[1]
}

// Address devolve o endereço do servidor passado como command line argument.
func Address() string {
	return arguments[2]
}
