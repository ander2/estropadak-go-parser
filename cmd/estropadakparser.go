package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"

	"estropadak.eus/estropadak-parser/pkg/parsers"
	"estropadak.eus/estropadak-parser/pkg/formatters"
)

func main() {
	var estropada estropadakParser.Estropada
	var err error
	var reader io.Reader

	typePtr := flag.String("t", "ACT", "Parser type: ACT or ARC")
	urlPtr := flag.String("u", "", "Content URL: http://www.liga-arc.com/es/regata/489/xvii.-hondarribiko-arrantzaleen-kofradiko-bandera")
	formatPtr := flag.String("f", "", "Output format: text, yaml or json")
	flag.Parse()

	estropada = estropadakParser.Estropada{}
	if *urlPtr == "" {
		reader = io.Reader(os.Stdin)
	} else {
		resp, err := http.Get(*urlPtr)
		if err != nil {
			fmt.Printf("Cannot fetch content on %s: %s\n", *urlPtr, err)
			os.Exit(1)
		}
		if resp.StatusCode != 200 {
			fmt.Printf("Cannot fetch content on %s: %s\n", *urlPtr, resp.Status)
			os.Exit(1)
		}
		defer resp.Body.Close()
		reader = resp.Body
	}
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
	if *formatPtr == "json" {
		formatters.Format_result_json(estropada)
	} else if *formatPtr == "yaml" {
		formatters.Format_result_yaml(estropada)
	} else {
		formatters.Format_result_text(estropada)
	}
}
