package service

import (
	"fmt"

	"github.com/crissilvaeng/chateado/client/core"
	"github.com/crissilvaeng/chateado/client/protocol"
)

type worker struct {
	queue     core.Collectable
	conn      core.Communicable
	next      int
	available bool
}

// Workerable define interface para interação com o Worker.
type Workerable interface {
	Perform(request protocol.Request)
	Run()
	Stop()
}

func (w *worker) Perform(request protocol.Request) {
	w.queue.Push(request)
}

func (w *worker) Run() {
	for w.available {
		if w.queue.Size() > 0 {
			req := w.queue.Pop()
			req.Reference = w.next

			raw, err := req.Build()
			if err != nil {
				fmt.Println(err)
			}

			err = w.conn.Write(raw)
			if err != nil {
				fmt.Println(err)
			}

			data, err := w.conn.Read()
			if err != nil {
				fmt.Println(err)
			}

			resp := protocol.Response{}
			err = resp.Build(string(data))
			if err != nil {
				fmt.Println(err)
			}

			w.next = resp.Reference

			if resp.Data != nil {
				contents := (*resp.Data)
				for _, content := range contents {
					fmt.Println(">", content.Sender, ":", content.Message)
				}
			}
		}
	}
}

func (w *worker) Stop() {
	w.available = false
}

// NewWorker cria um novo Worker baseado em Collectable e Communicable e
// retorna interface Workerable para sua manipulação.
func NewWorker(queue core.Collectable, conn core.Communicable) Workerable {
	return &worker{queue: queue, conn: conn, next: 0, available: true}
}
