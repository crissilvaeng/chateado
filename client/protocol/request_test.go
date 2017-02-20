package protocol

import "testing"

func TestRequestLogin(t *testing.T) {
	// Arrange
	expected := `{"cmd":"login","id":"joaquim","msgNr":43}`

	req := Request{
		Command:   "login",
		Sender:    "joaquim",
		Reference: 43,
	}

	// Act
	actual, err := req.Build()

	// Assert
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if expected != actual {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestRequestLogoff(t *testing.T) {
	// Arrange
	expected := `{"cmd":"logoff","id":"joaquim","msgNr":43}`

	req := Request{
		Command:   "logoff",
		Sender:    "joaquim",
		Reference: 43,
	}

	// Act
	actual, err := req.Build()

	// Assert
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if expected != actual {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestRequestEnviar(t *testing.T) {
	// Arrange
	expected := `{"cmd":"enviar","id":"Joaquim","msgNr":26,"dst":"Maria","data":"Oi!"}`

	dest := "Maria"
	data := "Oi!"

	req := Request{
		Command:   "enviar",
		Sender:    "Joaquim",
		Reference: 26,
		Recipient: &dest,
		Content:   &data,
	}

	// Act
	actual, err := req.Build()

	// Assert
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if expected != actual {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestRequestReceber(t *testing.T) {
	// Arrange
	expected := `{"cmd":"receber","id":"Joaquim","msgNr":26}`

	req := Request{
		Command:   "receber",
		Sender:    "Joaquim",
		Reference: 26,
	}

	// Act
	actual, err := req.Build()

	// Assert
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if expected != actual {
		t.Error("Expected ", expected, " got ", actual)
	}
}
