package formatters

import (
	"fmt"

	"github.com/go-yaml/yaml"
	"estropadak.eus/estropadak-parser/pkg/parsers"
)

func Format_result_yaml(estropada estropadakParser.Estropada) {
	data, err := yaml.Marshal(estropada)
	if err != nil {
		fmt.Printf("Cannot marshal into yaml format: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", data)
}