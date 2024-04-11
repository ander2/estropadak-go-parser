package estropadakParser

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func arc_parse_title(t *html.Tokenizer) string {
	var title string
	var next_token html.TokenType
	for t.Err() != io.EOF && title == "" {
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

func arc_parse_date_location(t *html.Tokenizer) (string, string) {
	var aux, date, location string
	var on_record = false
	var on_tag = false
	var col = 0
	var next_token html.TokenType
	for t.Err() != io.EOF && (date == "" || location == "") {
		tag, has_attrs := t.TagName()
		if next_token == html.StartTagToken && string(tag) == "div" && has_attrs {
			if attr_has_value(*t, "class", "detalles") {
				on_record = true
			}
		}
		if next_token == html.StartTagToken && string(tag) == "li" && on_record {
			col += 1
		}
		if next_token == html.SelfClosingTagToken && on_record && string(tag) == "br" {
			on_tag = true
		}
		if next_token == html.EndTagToken && on_record && string(tag) == "div" {
			on_record = false
		}
		if next_token == html.EndTagToken && on_record && string(tag) == "li" {
			on_tag = false
			aux = ""
			if col == 2 {
				date = strings.TrimSpace(date)
				if len(date) == 16 {
					date = "0" + date
				}
				date = strings.Replace(date, "20", "", 1)
				date = strings.Replace(date, "Ago", "Aug", 1)
				date += " CEST"
				date_aux, err := time.Parse(time.RFC822, date)
				if err != nil {
					fmt.Println(err)
				}
				date = date_aux.Format(time.RFC3339)
			}
		}
		if next_token == html.TextToken {
			if on_tag && col < 3 {
				aux = aux + strings.TrimSpace(string(t.Text()))
				date += aux + " "
			}
			if col == 4 {
				aux = aux + strings.TrimSpace(string(t.Text()))
				location = aux
			}
		}
		next_token = t.Next()
	}
	return date, location
}

func arc_parse_heats(t *html.Tokenizer) []Result {
	var col_counter int
	var heat_counter int
	var aux_text string
	var result Result
	var results []Result
	var result_parsed, on_record bool
	var next_token html.TokenType
	var section_end = false

	on_table := false

	for section_end == false {
		tag, has_attrs := t.TagName()
		if t.Token().Type == html.StartTagToken && string(tag) == "table" && has_attrs {
			if attr_has_value(*t, "class", "clasificacion tanda") {
				on_table = true
			}
		}

		if on_table {
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
					col_counter += 1
				}

				if next_token == html.EndTagToken && string(tag) == "th" && on_record {
					aux_lane := strings.TrimRight(strings.TrimSpace(aux_text), "º")
					lane, err := strconv.Atoi(aux_lane)
					if err == nil {
						result.Lane = lane
						result.Heat = heat_counter
					} else {
						fmt.Println(err)
						os.Exit(1)
					}
					aux_text = ""
				}

				if next_token == html.EndTagToken && string(tag) == "td" {
					if col_counter == 1 {
						team_name := result.TeamName + aux_text
						result.TeamName = strings.TrimSpace(team_name)
					}

					if col_counter > 1 && col_counter < 5 {
						result.Ziabogak = append(result.Ziabogak, strings.TrimSpace(aux_text))
					}

					if col_counter == 5 {
						result.Time = strings.TrimSpace(aux_text)
					}

					aux_text = ""
				}

				if next_token == html.TextToken && on_record {
					aux_text += string(t.Text())
				}

				if next_token == html.EndTagToken && string(tag) == "tr" && on_record {
					results = append(results, result)
				}

				if next_token == html.EndTagToken && string(tag) == "tbody" {
					on_record = false
					on_table = false
					result_parsed = true
					// Sort in heat and calculate heat position
					pos := 1
					sort.Sort(ByTime(results))
					for i, result := range results {
						if result.Heat == heat_counter {
							results[i].HeatPosition = pos
							pos += 1
						}
					}
				}
				next_token = t.Next()
			}
			result_parsed = false
		}

		if t.Token().Type == html.StartTagToken && string(tag) == "h2" && has_attrs {
			if attr_has_value(*t, "class", "volver-arriba") {
				section_end = true
			}
		}
		next_token = t.Next()
	}

	return results
}

func arc_parse_results(t *html.Tokenizer) []Result {
	var col_counter int
	var aux_text string
	var next_token html.TokenType
	var result Result
	var results []Result
	var result_parsed, on_record bool
	section_end := false

	for section_end == false {
		tag, has_attrs := t.TagName()
		if t.Token().Type == html.StartTagToken && string(tag) == "table" && has_attrs {
			if attr_has_value(*t, "class", "tab-item clasificacion clasificacion-regata grid_12 alpha") {
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

					if next_token == html.StartTagToken && string(tag) == "th" {
						aux_text = ""
					}

					if next_token == html.StartTagToken && string(tag) == "td" {
						col_counter += 1
						aux_text = ""

					}

					if next_token == html.EndTagToken && string(tag) == "th" && on_record {
						position, err := strconv.Atoi(strings.TrimRight(aux_text, "º"))
						if err == nil {
							result.Position = position
						} else {
							result.Position = 12
						}
						aux_text = ""
					}

					if next_token == html.EndTagToken && string(tag) == "td" {
						if col_counter == 1 {
							team_name := result.TeamName + aux_text
							result.TeamName = strings.TrimSpace(team_name)
						}
						if col_counter == 2 {
							time := result.Time + aux_text
							result.Time = strings.TrimSpace(time)
						}
						if col_counter == 3 {
							points, err := strconv.Atoi(aux_text)
							if err != nil {
								result.Points = points
							}
						}

					}

					if next_token == html.TextToken && on_record {
						aux_text += string(t.Text())

					}

					if next_token == html.EndTagToken && string(tag) == "tr" && on_record {
						results = append(results, result)
					}

					if next_token == html.EndTagToken && string(tag) == "tbody" {
						result_parsed = true
						on_record = false
						section_end = true
					}
					next_token = t.Next()
				}
			}
		}
		next_token = t.Next()
	}
	return results
}

func Arc_parse_estropadak_doc(estropada *Estropada, doc io.Reader) (string, error) {

	var title string
	var results []Result
	var heat_results []Result
	tokenizer := html.NewTokenizer(doc)
	for tokenizer.Err() != io.EOF {
		if title == "" {
			title = arc_parse_title(tokenizer)
			if title != "" {
				estropada.Name = title
			}
		}
		if estropada.Date == "" {
			date, location := arc_parse_date_location(tokenizer)
			if date != "" {
				estropada.Date = date
			}
			if location != "" {
				estropada.Location = location
			}
		}
		if len(estropada.Results) == 0 {
			results = arc_parse_results(tokenizer)
			if len(results) > 0 {
				estropada.Results = results
			}
		}
		if len(heat_results) == 0 {
			heat_results = arc_parse_heats(tokenizer)
			for _, res := range heat_results {
				for i, part_res := range estropada.Results {
					if part_res.TeamName == res.TeamName {
						estropada.Results[i].Ziabogak = res.Ziabogak
					}
				}
			}
		}
		tokenizer.Next()
	}
	return "ok", nil
}

func Arc_parse_calendar(doc io.Reader) ([]Estropada, error) {
	var estropadak []Estropada
	var estropada Estropada
	var next_token html.TokenType
	t := html.NewTokenizer(doc)
	var calendar_parsed, on_record bool
	var col_counter int
	var aux_text, league string

	for t.Err() != io.EOF && !calendar_parsed {
		tag, has_attrs := t.TagName()
		if t.Token().Type == html.StartTagToken && string(tag) == "tr" && has_attrs {
			if attr_has_value(*t, "class", "tab-item") {
				on_record = true
				estropada = Estropada{}
			}
			if attr_has_value(*t, "class", "g1") {
				league = "ARC1"
			}
			if attr_has_value(*t, "class", "g2") {
				league = "ARC2"
			}
		}

		if next_token == html.StartTagToken && string(tag) == "td" && on_record {
			col_counter += 1
		}

		if next_token == html.EndTagToken && string(tag) == "td" && on_record {
			if col_counter == 1 {
				date_aux := strings.TrimSpace(aux_text)
				date_parts := strings.Split(date_aux, " ")
				date_parts[1] = strings.Replace(date_parts[1], "Junio", "Jun", 1)
				date_parts[1] = strings.Replace(date_parts[1], "Julio", "Jul", 1)
				date_parts[1] = strings.Replace(date_parts[1], "Agosto", "Aug", 1)

				year := time.Now().Year()
				date_aux = fmt.Sprintf("%s %s %d 00:00 CEST", date_parts[0], date_parts[1], year%100)
				if len(date_aux) == 19 {
					date_aux = "0" + date_aux
				}
				date, err := time.Parse(time.RFC822, date_aux)
				if err != nil {
					fmt.Println(err)
				}
				estropada.Date = date.Format(time.RFC3339)
			}

			if col_counter == 2 {
				estropada.Name = strings.TrimSpace(aux_text)
			}

			aux_text = ""
		}

		if next_token == html.EndTagToken && string(tag) == "tr" && on_record {
			estropada.League = league
			estropadak = append(estropadak, estropada)
			col_counter = 0
			aux_text = ""
			on_record = false
		}

		if next_token == html.EndTagToken && string(tag) == "tbody" && on_record {
			calendar_parsed = true
		}

		if next_token == html.TextToken && on_record {
			aux_text += string(t.Raw())
		}
		next_token = t.Next()
	}
	return estropadak, nil
}
