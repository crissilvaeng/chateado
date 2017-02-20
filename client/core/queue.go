package core

import (
	"sync"

	"github.com/crissilvaeng/chateado/client/protocol"
)

// Collectable é uma interface publica para manipulação da fila.
type Collectable interface {
	Push(message protocol.Request)
	Pop() protocol.Request
	Size() int
}

type queue struct {
	messages []protocol.Request
	mutex    *sync.Mutex
}

func (q *queue) Push(message protocol.Request) {
	q.mutex.Lock()
	q.messages = append(q.messages, message)
	q.mutex.Unlock()
}

func (q *queue) Pop() protocol.Request {
	q.mutex.Lock()
	message := q.messages[len(q.messages)-1]
	q.messages = q.messages[:len(q.messages)-1]
	q.mutex.Unlock()
	return message
}

func (q *queue) Size() int {
	return len(q.messages)
}

// NewQueue devolve um interface unica para manipulação da fila privada.
func NewQueue() Collectable {
	return &queue{mutex: &sync.Mutex{}}
}
