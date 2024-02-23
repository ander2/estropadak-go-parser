package estropadakParser

import (
	"os"
	"sort"
	"testing"

	"golang.org/x/net/html"
)

func TestParseTitle(t *testing.T) {

	var title string
	doc, err := os.Open("./html/zarautz_act_2023_1.html")
	if err != nil {
		t.Error("Cannot open file", err)
	}
	expected := "XLVI. Zarauzko Ikurriña (J1) (19-08-2023)"
	tokenizer := html.NewTokenizer(doc)

	title = parse_title(tokenizer)
	if title != expected {
		t.Errorf("%s != %s\n", title, expected)
	}

}

func TestParseLocation(t *testing.T) {
	doc, err := os.Open("./html/zarautz_act_2023_1.html")
	if err != nil {
		t.Error("Cannot open file", err)
	}
	expected := "Zarautz Gipuzkoa"
	tokenizer := html.NewTokenizer(doc)

	location := parse_location(tokenizer)
	if location != expected {
		t.Errorf("%s != %s\n", location, expected)
	}

}

func TestParseHeats(t *testing.T) {

	doc, err := os.Open("./html/zarautz_act_2023_1.html")
	if err != nil {
		t.Error("Cannot open file", err)
	}
	tokenizer := html.NewTokenizer(doc)

	expected := []Result{
		{TeamName: "AMENABAR DONOSTIARRA UR KIROLAK", HeatPosition: 2, Ziabogak: []string{"04:53", "10:03", "15:27"}},
		{TeamName: "BERMEO URDAIBAI", HeatPosition: 1, Time: "20:12,14", Ziabogak: []string{"04:55", "09:57", "15:16"}},
		{TeamName: "CR CABO DA CRUZ", HeatPosition: 3, Time: "21:07,90"},
		{TeamName: "GETARIA", HeatPosition: 2, Time: "20:49,76"},
		{TeamName: "HONDARRIBIA", HeatPosition: 3, Time: "20:32,92"},
		{TeamName: "ITSASOKO AMA SANTURTZI", HeatPosition: 1, Time: "21:16,36"},
		{TeamName: "KAIKU BEREZ GALANTA", HeatPosition: 4, Time: "21:31,20"},
		{TeamName: "LEKITTARRA ELECNOR", HeatPosition: 4, Time: "21:11,08"},
		{TeamName: "ONDARROA CIKAUTXO", HeatPosition: 3, Time: "21:29,16"},
		{TeamName: "ORIO ORIALKI", HeatPosition: 1, Time: "20:20,30"},
		{TeamName: "SAMERTOLAMEU FANDICOSTA", HeatPosition: 2, Time: "21:22,44"},
		{TeamName: "ZIERBENA BAHIAS DE BIZKAIA", HeatPosition: 4, Time: "20:38,46"},
	}

	results := parse_heats(tokenizer)
	if len(results) != 12 {
		t.Errorf("%d != %d\n", len(results), 12)
	}

	sort.Sort(ByName(results))
	for i, result := range results {
		if expected[i].HeatPosition != result.HeatPosition {
			t.Errorf("Heat position %d is not %d\n", expected[i].HeatPosition, result.HeatPosition)
		}
		if i < 2 {
			for j := 0; j < 3; j += 1 {
				if expected[i].Ziabogak[j] != result.Ziabogak[j] {
					t.Errorf("Ziaboga %s is not %s\n", expected[i].Ziabogak[j], result.Ziabogak[j])
				}
			}
		}
	}
}

func TestParseHeatsEuskotren(t *testing.T) {

	doc, err := os.Open("./html/fabrika_euskotren_2023.html")
	if err != nil {
		t.Error("Cannot open file", err)
	}
	tokenizer := html.NewTokenizer(doc)
	expected := []Result{
		{TeamName: "CR CABO DA CRUZ", HeatPosition: 3, Time: "11:30,80", Ziabogak: []string{"05:55"}},
		{TeamName: "DONOSTIA ARRAUN LAGUNAK", HeatPosition: 2, Time: "10:48,98", Ziabogak: []string{"05:35"}},
		{TeamName: "HIBAIKA JAMONES ANCIN", HeatPosition: 2, Time: "11:25,96"},
		{TeamName: "HONDARRIBIA BERTAKO IGOGAILUAK", HeatPosition: 1, Time: "11:12,86"},
		{TeamName: "NORTINDAL DONOSTIARRA UR KIROLAK", HeatPosition: 1, Time: "10:44,90", Ziabogak: []string{"05:37"}},
		{TeamName: "ORIO ORIALKI", HeatPosition: 3, Time: "10:54,10"},
		{TeamName: "SD TIRÁN PEREIRA", HeatPosition: 4, Time: "11:50,24"},
		{TeamName: "TOLOSALDEA ARRAUN KLUBA", HeatPosition: 4, Time: "11:12,46"},
	}

	results := parse_heats(tokenizer)
	if len(results) != 8 {
		t.Errorf("%d != %d\n", len(results), 8)
	}

	sort.Sort(ByName(results))
	for i, result := range results {
		if expected[i].HeatPosition != result.HeatPosition {
			t.Errorf("Heat position %d is not %d\n", expected[i].HeatPosition, result.HeatPosition)
		}
		if i < 2 {
			if expected[i].Ziabogak[0] != result.Ziabogak[0] {
				t.Errorf("Ziaboga %s is not %s\n", expected[i].Ziabogak[0], result.Ziabogak[0])
			}
		}
	}
}
