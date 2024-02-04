package estropadakParser

import (
	// "fmt"
	// "os"
	"io"
	// "strconv"
	// "strings"

	"golang.org/x/net/html"
)

func arc_parse_title(t *html.Tokenizer) string {
	var title string
	var next_token html.TokenType
	for t.Err() != io.EOF && title == ""{
		tag, _ := t.TagName()
		if t.Token().Type == html.StartTagToken && string(tag) == "h2" {
			for next_token != html.EndTagToken {
				if next_token == html.TextToken {
					title += t.Token().Data
				}
				next_token = t.Next()
			}
		}
		next_token = t.Next()
    }
	return title
}


func Arc_parse_estropadak_doc(estropada *Estropada, doc io.Reader) (string, error) {

	var title string
	// var results []Result
	tokenizer := html.NewTokenizer(doc)
	for tokenizer.Err() != io.EOF {
		if title == "" {
			title = arc_parse_title(tokenizer)
			if title != "" {
				estropada.Name = title
			}
		}
		// if len(estropada.Results) == 0 {
		// 	results = parse_heats(tokenizer)
		// 	if len(results) > 0 {
		// 		estropada.Results = results
		// 	}
		// }
		// general_results := parse_results(tokenizer)
		// for _, res := range general_results  {
		// 	for i, part_res := range estropada.Results  {
		// 		if part_res.TeamName == res.TeamName {
		// 			estropada.Results[i].Position = res.Position
		// 			estropada.Results[i].Time = res.Time
		// 		}
		// 	}
		// }
		tokenizer.Next()
	}
	return "ok", nil
}