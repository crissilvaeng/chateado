package core

import (
	"testing"

	"github.com/crissilvaeng/chateado/client/protocol"
)

func TestCollectablePushPassOk(t *testing.T) {
	// Arrange
	queue := NewQueue()
	req := protocol.Request{}
	expected := 1

	// Act
	queue.Push(req)
	actual := queue.Size()

	// Assert
	if expected != actual {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestCollectablePopPassOk(t *testing.T) {
	// Arrange
	queue := NewQueue()
	req := protocol.Request{}
	expected := 0

	// Act
	queue.Push(req)
	queue.Pop()
	actual := queue.Size()

	// Assert
	if expected != actual {
		t.Error("Expected ", expected, " got ", actual)
	}
}
