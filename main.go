package main

import (
	"fmt"
	"os"
	"io"
	"sort"
	"text/tabwriter"
)

import "estropadak.eus/estropadak-parser/parsers"

func main() {
	var estropada estropadakParser.Estropada

	estropada = estropadakParser.Estropada{}
	reader := io.Reader(os.Stdin)
	_, err := estropadakParser.Act_parse_estropadak_doc(&estropada, reader)
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
