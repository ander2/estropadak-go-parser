package formatters

import (
	"encoding/json"
	"fmt"

	"estropadak.eus/estropadak-parser/pkg/parsers"
)

func Format_result_json(estropada estropadakParser.Estropada) {
	data, _ := json.MarshalIndent(estropada, "", "  ")
	fmt.Printf("%s\n", data)
}