package main

import (
	"bufio"
	"os"

	"github.com/crissilvaeng/chateado/client/args"
	"github.com/crissilvaeng/chateado/client/core"
	"github.com/crissilvaeng/chateado/client/service"
)

func main() {
	service.Sender = args.Username()

	comm, err := core.NewConnection("tcp", args.Address())
	if err != nil {
		os.Exit(2)
	}

	queue := core.NewQueue()
	worker := service.NewWorker(queue, comm)

	defer worker.Perform(service.Logoff())
	defer worker.Stop()

	go worker.Run()

	worker.Perform(service.Login())

	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		worker.Perform(service.Sends("maria", message))
	}
}
