package protocol

import "testing"

func TestResponseLogin(t *testing.T) {
	// Arrange
	data := []Content{}
	data = append(data, Content{Sender: "maria", Message: "oi!"})
	data = append(data, Content{Sender: "maria", Message: "kd vc?"})

	expected := Response{
		Identificator: "0",
		Reference:     54,
		Data:          &data,
	}

	raw := `{
                "id": "0",
                "msgNr": 54,
                "data": [
                    {"src":"maria","data":"oi!"},
                    {"src":"maria","data":"kd vc?"}
                ]
            }`

	actual := Response{}

	// Act
	err := actual.Build(raw)

	// Assert
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if expected.Identificator != actual.Identificator {
		t.Error("Expected ", expected.Identificator, " got ", actual.Identificator)
	}

	if expected.Reference != actual.Reference {
		t.Error("Expected ", expected.Reference, " got ", actual.Reference)
	}

	if (*expected.Data)[0].Sender != (*actual.Data)[0].Sender {
		t.Error("Expected ", (*expected.Data)[0].Sender, " got ", (*actual.Data)[0].Sender)
	}

	if (*expected.Data)[0].Message != (*actual.Data)[0].Message {
		t.Error("Expected ", (*expected.Data)[0].Message, " got ", (*actual.Data)[0].Message)
	}

	if (*expected.Data)[1].Sender != (*actual.Data)[1].Sender {
		t.Error("Expected ", (*expected.Data)[1].Sender, " got ", (*actual.Data)[1].Sender)
	}

	if (*expected.Data)[1].Message != (*actual.Data)[1].Message {
		t.Error("Expected ", (*expected.Data)[1].Message, " got ", (*actual.Data)[1].Message)
	}
}

func TestResponseEnviar(t *testing.T) {
	// Arrange
	expected := Response{
		Identificator: "0",
		Reference:     27,
	}

	raw := `{
                "id": "0",
                "msgNr": 27
            }`

	actual := Response{}

	// Act
	err := actual.Build(raw)

	// Assert
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if expected.Identificator != actual.Identificator {
		t.Error("Expected ", expected.Identificator, " got ", actual.Identificator)
	}

	if expected.Reference != actual.Reference {
		t.Error("Expected ", expected.Reference, " got ", actual.Reference)
	}
}

func TestResponseReceber(t *testing.T) {
	// Arrange
	data := []Content{}
	data = append(data, Content{Sender: "maria", Message: "oi!"})
	data = append(data, Content{Sender: "maria", Message: "kd vc?"})

	expected := Response{
		Identificator: "0",
		Reference:     27,
		Data:          &data,
	}

	raw := `{
                "id": "0",
                "msgNr": 27,
                "data": [
                    {"src":"maria","data":"oi!"},
                    {"src":"maria","data":"kd vc?"}
                ]
            }`

	actual := Response{}

	// Act
	err := actual.Build(raw)

	// Assert
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if expected.Identificator != actual.Identificator {
		t.Error("Expected ", expected.Identificator, " got ", actual.Identificator)
	}

	if expected.Reference != actual.Reference {
		t.Error("Expected ", expected.Reference, " got ", actual.Reference)
	}

	if (*expected.Data)[0].Sender != (*actual.Data)[0].Sender {
		t.Error("Expected ", (*expected.Data)[0].Sender, " got ", (*actual.Data)[0].Sender)
	}

	if (*expected.Data)[0].Message != (*actual.Data)[0].Message {
		t.Error("Expected ", (*expected.Data)[0].Message, " got ", (*actual.Data)[0].Message)
	}

	if (*expected.Data)[1].Sender != (*actual.Data)[1].Sender {
		t.Error("Expected ", (*expected.Data)[1].Sender, " got ", (*actual.Data)[1].Sender)
	}

	if (*expected.Data)[1].Message != (*actual.Data)[1].Message {
		t.Error("Expected ", (*expected.Data)[1].Message, " got ", (*actual.Data)[1].Message)
	}
}
