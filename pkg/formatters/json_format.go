package formatters

import (
	"encoding/json"
	"fmt"
	"os"

	"estropadak.eus/estropadak-parser/pkg/parsers"
)

func Format_result_json(estropada estropadakParser.Estropada) {
	data, err := json.MarshalIndent(estropada, "", "  ")
	if err != nil {
		fmt.Printf("Cannot marshal into json format: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", data)
}