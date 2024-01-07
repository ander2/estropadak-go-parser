package main

import (
	"fmt"
	"os"
	"io"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type Result struct {
	TeamName string
	Position int
	Lane int
	Time string
}

type Estropada struct {
	Name string
	Date string
	Location string
	Results []Result
}

func main() {
	var estropada Estropada

	estropada = Estropada{}
	reader := io.Reader(os.Stdin)
	result, err := parse_estropadak_doc(&estropada, reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse doc %s", err)
		os.Exit(1)
	}
	fmt.Printf("Got %s %s after parse\n", result, estropada.Name)
	for _, res := range estropada.Results {
		fmt.Printf("- [%d] %s - %s \n", res.Position, res.TeamName, res.Time)
	}
}


func parse_title(t *html.Tokenizer) string {
	title := ""
	tag, _ := t.TagName()
	if t.Token().Type == html.StartTagToken && string(tag) == "h3" {
		next_token := t.Next()
		for next_token != html.EndTagToken {
			if next_token == html.TextToken {
				title += t.Token().Data
			}
			next_token = t.Next()
		}
	}
	return title
}


func parse_results(t *html.Tokenizer) []Result {
	var col_counter int
	// var pos_text string
	var aux_text string
	// var aux_name []byte
	var result Result
	var results []Result
	var result_parsed, on_record bool

	tag, has_attrs := t.TagName()
	if t.Token().Type == html.StartTagToken && string(tag) == "table" && has_attrs {
		for attr, val, more_attrs := t.TagAttr(); more_attrs == true; attr, val, more_attrs = t.TagAttr() {
			fmt.Printf("Found attr -%s- with val %s\n", attr, val)
			if string(attr) == "summary" && string(val) == "Clasificación por regata" {
				fmt.Println("Found table with summary")
				next_token := t.Next()
				for result_parsed == false {
					tag, has_attrs = t.TagName()

					if next_token == html.StartTagToken && string(tag) == "tbody" {
						on_record = true
					}

					if next_token == html.StartTagToken && string(tag) == "tr" && on_record {
						col_counter = 0
						result = Result{}
					}

					if next_token == html.StartTagToken && string(tag) == "td" {
						col_counter +=1
					}

					if next_token == html.EndTagToken && string(tag) == "td" {
						if col_counter == 1 {
							aux_position := strings.TrimSpace(aux_text)
							pos, err := strconv.Atoi(aux_position)
							if err == nil {
								result.Position = pos
								fmt.Printf("Recorded %s %d %d\n", aux_position, pos, result.Position)
							} else {
								fmt.Println(err)
								os.Exit(1)
							}
						}
						if col_counter == 2 {
							team_name := result.TeamName + aux_text
							result.TeamName = strings.TrimSpace(team_name)
						}

						if col_counter == 4 {
							result.Time = strings.TrimSpace(aux_text)
						}
						aux_text = ""
					}

					if next_token == html.TextToken && on_record{
						aux_text += string(t.Text())
						// aux_position += aux_text

					}

					// if next_token == html.TextToken && col_counter == 2 && on_record{
					// 	aux_text = t.Text()
					// 	aux_name = aux_name
					// }

					if next_token == html.EndTagToken && string(tag) == "tr" && on_record{
						fmt.Printf("Appending %s\n", result.TeamName)
						results = append(results, result)
					}

					if next_token == html.EndTagToken && string(tag) == "tbody" {
						result_parsed = true
					}
					next_token = t.Next()
				}
			}
		}
	}
	return results
}

func parse_estropadak_doc(estropada *Estropada, doc io.Reader) (string, error) {

	var title string;
	tokenizer := html.NewTokenizer(doc)
	for tokenizer.Err() != io.EOF {
		if title == "" {
			title = parse_title(tokenizer)
			if title != "" {
				fmt.Println(title)
				estropada.Name = title
			}
		}
		if len(estropada.Results) == 0 {
			results := parse_results(tokenizer)
			if len(results) > 0 {
				estropada.Results = results
			}
		}
		tokenizer.Next()
	}
	return "ok", nil
}