package core

import "net"

// Communicable prove uma interface para operar leitura e escrita.
type Communicable interface {
	Write(message string) error
	Read() (string, error)
}

type connection struct {
	conn net.Conn
}

func (c *connection) Write(message string) error {
	raw := []byte(message)

	_, err := c.conn.Write(raw)
	if err != nil {
		return err
	}

	return nil
}

func (c *connection) Read() (string, error) {
	raw := make([]byte, 1024)

	length, err := c.conn.Read(raw)
	if err != nil {
		return "", err
	}

	data := raw[:length]

	return string(data), nil
}

// NewConnection estabelece uma conexão baseado no protocolo e endereço.
func NewConnection(protocol, address string) (Communicable, error) {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		return nil, err
	}

	comm := &connection{}
	comm.conn = conn

	return comm, nil
}
