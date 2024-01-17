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
	HeatPosition int
	Heat int
	Lane int
	Time string
	Ziabogak []string
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
		fmt.Printf("- [%d] %s %s - %s \n", res.Position, res.TeamName, res.Ziabogak, res.Time)
	}
}


func parse_title(t *html.Tokenizer) string {
	var title string
	var next_token html.TokenType
	for t.Err() != io.EOF && title == ""{
		tag, _ := t.TagName()
		if t.Token().Type == html.StartTagToken && string(tag) == "h3" {
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

func parse_heats(t *html.Tokenizer) []Result {
	var col_counter int
	var heat_counter int
	// var pos_text string
	var aux_text string
	// var aux_name []byte
	var result Result
	var results []Result
	var result_parsed, on_record bool
	var next_token html.TokenType

	for heat_counter < 3 {
		tag, has_attrs := t.TagName()
		if t.Token().Type == html.StartTagToken && string(tag) == "table" && has_attrs {
			for attr, val, more_attrs := t.TagAttr(); more_attrs == true; attr, val, more_attrs = t.TagAttr() {
				if string(attr) == "summary" && string(val) == "Resultados por regata" {
					heat_counter += 1
					next_token = t.Next()
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
								aux_lane := strings.TrimSpace(aux_text)
								lane, err := strconv.Atoi(aux_lane)
								if err == nil {
									result.Lane = lane
								} else {
									fmt.Println(err)
									os.Exit(1)
								}
							}
							if col_counter == 2 {
								team_name := result.TeamName + aux_text
								result.TeamName = strings.TrimSpace(team_name)
							}

							if col_counter > 2 && col_counter < 6 {
								result.Ziabogak = append(result.Ziabogak, strings.TrimSpace(aux_text))
							}

							if col_counter == 7 {
								aux_pos := strings.TrimSpace(aux_text)
								pos, err := strconv.Atoi(aux_pos)
								if err == nil {
									result.HeatPosition = pos
								} else {
									fmt.Println(err)
									os.Exit(1)
								}
							}
							aux_text = ""
						}

						if next_token == html.TextToken && on_record{
							aux_text += string(t.Text())
						}

						if next_token == html.EndTagToken && string(tag) == "tr" && on_record{
							results = append(results, result)
						}

						if next_token == html.EndTagToken && string(tag) == "tbody" {
							on_record = false
							result_parsed = true
						}
						next_token = t.Next()
					}
					result_parsed = false
				}
			}
		}
		next_token = t.Next()
	}
	return results
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
						on_record = false
					}
					next_token = t.Next()
				}
			}
		}
	}
	return results
}

func parse_estropadak_doc(estropada *Estropada, doc io.Reader) (string, error) {

	var title string
	var results []Result
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
			results = parse_heats(tokenizer)
			if len(results) > 0 {
				estropada.Results = results
			}
		}
		// if len(estropada.Results) == 0 {
		general_results := parse_results(tokenizer)
		for _, res := range general_results  {
			for i, part_res := range estropada.Results  {
				fmt.Println(">>>>>>>", res.TeamName, part_res.TeamName)
				if part_res.TeamName == res.TeamName {
					estropada.Results[i].Position = res.Position
					estropada.Results[i].Time = res.Time
				}
			}
		}
		// 	if len(results) > 0 {
		// 		estropada.Results = results
		// 	}
		// }
		tokenizer.Next()
	}
	return "ok", nil
}
