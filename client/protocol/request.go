package protocol

import "encoding/json"

// Requestable define métodos para mensagens de requisição.
type Requestable interface {
	Build() (string, error)
}

// Request encapsula as propriedades necessarias para uma requisição.
type Request struct {
	Command   string  `json:"cmd"`
	Sender    string  `json:"id"`
	Reference int     `json:"msgNr"`
	Recipient *string `json:"dst,omitempty"`
	Content   *string `json:"data,omitempty"`
}

// Build transforma a estrutura de requsição em um JSON.
func (r *Request) Build() (string, error) {
	req, err := json.Marshal(r)

	if err != nil {
		return "", err
	}

	return string(req), nil
}
