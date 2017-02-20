package protocol

import "encoding/json"

// Responsable definie metos para mensags de resposta.
type Responsable interface {
	Build(message string) error
}

// Content na response traz o destinatario e a mensagem.
type Content struct {
	Sender  string `json:"src"`
	Message string `json:"data"`
}

// Response encapsula propriedades da resposta.
type Response struct {
	Identificator string     `json:"id"`
	Reference     int        `json:"msgNr"`
	Data          *[]Content `json:"data,omitempty"`
}

// Build converte o JSON da resposta em uma estrutura.
func (r *Response) Build(raw string) error {
	err := json.Unmarshal([]byte(raw), &r)

	if err != nil {
		return err
	}

	return nil
}
