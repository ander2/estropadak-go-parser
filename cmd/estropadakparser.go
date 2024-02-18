package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"estropadak.eus/estropadak-parser/pkg/parsers"
	"estropadak.eus/estropadak-parser/pkg/formatters"
)

func main() {
	var estropada estropadakParser.Estropada
	var err error

	typePtr := flag.String("t", "ACT", "Parser type: ACT or ARC")
	flag.Parse()

	estropada = estropadakParser.Estropada{}
	reader := io.Reader(os.Stdin)
	if *typePtr == "ACT" {
		_, err = estropadakParser.Act_parse_estropadak_doc(&estropada, reader)
	} else if *typePtr == "ARC" {
		_, err = estropadakParser.Arc_parse_estropadak_doc(&estropada, reader)
	} else {
		fmt.Fprintln(os.Stderr, "Invalid parser type: select ACT or ARC")
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse doc %s", err)
		os.Exit(1)
	}
	sort.Sort(estropadakParser.ByPosition(estropada.Results))
	formatters.Format_result_text(estropada)
}
