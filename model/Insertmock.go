package model

import (
	"encoding/json"
	"log"
	"os"
)

func CreateMock() []Produto {
	var produtos []Produto

	file, _ := os.ReadFile("model/MOCK_DATA.json")

	err := json.Unmarshal(file, &produtos)

	if err != nil {
		log.Println("Something happened", err)
	}

	return produtos
}
