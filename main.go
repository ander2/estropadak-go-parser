package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"

	"estropadak.eus/estropadak-parser/parsers"
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
	format_result(estropada)
}

func format_result(estropada estropadakParser.Estropada) {
	sort.Sort(estropadakParser.ByPosition(estropada.Results))
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "\t%s\n", estropada.Name)
	for _, res := range estropada.Results {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", res.Position, res.TeamName, res.Ziabogak, res.Time)
	}
	w.Flush()
}
