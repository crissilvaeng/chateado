package service

import "github.com/crissilvaeng/chateado/client/protocol"

// Sender é o username do usuário logado no cliente de chat.
var Sender string

func build(command string) protocol.Request {
	return protocol.Request{
		Command: command,
		Sender:  Sender,
	}
}

// Login constroi uma requisição de login.
func Login() protocol.Request {
	return build("login")
}

// Logoff constroi uma requisição de logoff.
func Logoff() protocol.Request {
	return build("logoff")
}

// Sends constroi uma requisição de enviar.
func Sends(recipient, content string) protocol.Request {
	req := build("enviar")
	req.Recipient = &recipient
	req.Content = &content
	return req
}

// Receives constroi uma requisição de receber.
func Receives() protocol.Request {
	return build("receber")
}
