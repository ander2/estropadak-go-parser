package formatters

import (
	"fmt"
	"os"
	"text/tabwriter"

	estropadakParser "estropadak.eus/estropadak-parser/pkg/parsers"
)

func Format_result_text(estropada estropadakParser.Estropada) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "\t%s\n", estropada.Name)
	for _, res := range estropada.Results {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", res.Position, res.TeamName, res.Ziabogak, res.Time)
	}
	w.Flush()
}

func Format_calendar_text(estropadak []estropadakParser.Estropada) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for i, estropada := range estropadak {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i, estropada.Date, estropada.Location, estropada.Name)
	}
	w.Flush()
}
