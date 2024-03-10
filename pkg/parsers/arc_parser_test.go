package estropadakParser

import (
	"os"
	"sort"
	"testing"

	"golang.org/x/net/html"
)

func TestParseARCTitle(t *testing.T) {

	var title string
	doc, err := os.Open("./test_data/hondarribia_arc1_2023.html")
	if err != nil {
		t.Error("Cannot open file", err)
	}
	expected := "XVII. HONDARRIBIKO ARRANTZALEEN KOFRADIKO BANDERA"
	tokenizer := html.NewTokenizer(doc)

	title = arc_parse_title(tokenizer)
	if title != expected {
		t.Errorf("%s != %s\n", title, expected)
	}

}

func TestParseARCDateLocation(t *testing.T) {
	doc, err := os.Open("./test_data/hondarribia_arc1_2023.html")
	if err != nil {
		t.Error("Cannot open file", err)
	}
	expected_date := "6 Ago 2023"
	expected_location := "Hondarribia (Gipuzkoa)"
	tokenizer := html.NewTokenizer(doc)
	date, location := arc_parse_date_location(tokenizer)
	if date != expected_date {
		t.Errorf("%s != %s", date, expected_date)
	}
	if location != expected_location {
		t.Errorf("%s != %s", location, expected_location)
	}
}

func TestParseARCHeats(t *testing.T) {

	doc, err := os.Open("./test_data/hondarribia_arc1_2023.html")
	if err != nil {
		t.Error("Cannot open file", err)
	}
	tokenizer := html.NewTokenizer(doc)

	expected := []Result{
		{Heat: 3, TeamName: "Arkote A.  T.", HeatPosition: 2, Time: "20:59,16", Ziabogak: []string{"5:17", "10:20", "16:06"}},
		{Heat: 1, TeamName: "Busturialdea", HeatPosition: 2, Time: "21:44,97", Ziabogak: []string{"5:29", "10:39", "16:35"}},
		{Heat: 2, TeamName: "Camargo", HeatPosition: 2, Time: "21:15,56"},
		{Heat: 1, TeamName: "Castro Canteras de Santullan", HeatPosition: 4, Time: "22:09,71"},
		{Heat: 1, TeamName: "Deusto-bilbao", HeatPosition: 3, Time: "22:05,25"},
		{Heat: 1, TeamName: "Hondarribia", HeatPosition: 1, Time: "21:37,17"},
		{Heat: 3, TeamName: "Lapurdi Antton Bilbao", HeatPosition: 1, Time: "20:47,50"},
		{Heat: 2, TeamName: "Pedreña", HeatPosition: 3, Time: "21:18,59"},
		{Heat: 3, TeamName: "San Juan CMO Valves", HeatPosition: 4, Time: "21:19,65"},
		{Heat: 3, TeamName: "Sanpedrotarra A.e.", HeatPosition: 3, Time: "21:00,32"},
		{Heat: 2, TeamName: "Zarautz Gesalaga Okelan", HeatPosition: 1, Time: "21:10,00"},
		{Heat: 2, TeamName: "Zumaiako Telmo Deun A.k.e..", HeatPosition: 4, Time: "21:22,77"},
	}

	results := arc_parse_heats(tokenizer)
	if len(results) != 12 {
		t.Errorf("%d != %d\n", len(results), 12)
	}

	sort.Sort(ByName(results))
	for i, result := range results {
		if expected[i].TeamName != result.TeamName {
			t.Errorf("Team %s is not %s\n", expected[i].TeamName, result.TeamName)
		}

		if expected[i].Heat != result.Heat {
			t.Errorf("Team %s heat %d is not %d\n", result.TeamName, expected[i].Heat, result.Heat)
		}

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

func TestParseARCResult(t *testing.T) {

	doc, err := os.Open("./test_data/hondarribia_arc1_2023.html")
	if err != nil {
		t.Error("Cannot open file", err)
	}
	tokenizer := html.NewTokenizer(doc)

	expected := []Result{
		{Position: 2, Heat: 3, TeamName: "Arkote A.  T.", HeatPosition: 2, Time: "20:59,16", Ziabogak: []string{"5:17", "10:20", "16:06"}},
		{Position: 10, Heat: 1, TeamName: "Busturialdea", HeatPosition: 2, Time: "21:44,97", Ziabogak: []string{"5:29", "10:39", "16:35"}},
		{Position: 5, Heat: 2, TeamName: "Camargo", HeatPosition: 2, Time: "21:15,56"},
		{Position: 12, Heat: 1, TeamName: "Castro Canteras de Santullan", HeatPosition: 4, Time: "22:09,71"},
		{Position: 11, Heat: 1, TeamName: "Deusto-bilbao", HeatPosition: 3, Time: "22:05,25"},
		{Position: 9, Heat: 1, TeamName: "Hondarribia", HeatPosition: 1, Time: "21:37,17"},
		{Position: 1, Heat: 3, TeamName: "Lapurdi Antton Bilbao", HeatPosition: 1, Time: "20:47,50"},
		{Position: 6, Heat: 2, TeamName: "Pedreña", HeatPosition: 3, Time: "21:18,59"},
		{Position: 7, Heat: 3, TeamName: "San Juan CMO Valves", HeatPosition: 4, Time: "21:19,65"},
		{Position: 3, Heat: 3, TeamName: "Sanpedrotarra A.e.", HeatPosition: 3, Time: "21:00,32"},
		{Position: 4, Heat: 2, TeamName: "Zarautz Gesalaga Okelan", HeatPosition: 1, Time: "21:10,00"},
		{Position: 8, Heat: 2, TeamName: "Zumaiako Telmo Deun A.k.e..", HeatPosition: 4, Time: "21:22,77"},
	}

	results := arc_parse_results(tokenizer)
	if len(results) != 12 {
		t.Errorf("%d != %d\n", len(results), 12)
	}

	sort.Sort(ByName(results))
	for i, result := range results {
		if expected[i].Position != result.Position {
			t.Errorf("Team %s position %d is not %d\n", result.TeamName, expected[i].Position, result.Position)
		}
	}
}

func TestArcCalendar(t *testing.T) {
	COUNT := 27
	doc, err := os.Open("./test_data/arc_2023.html")
	if err != nil {
		t.Error("Cannot open file", err)
	}

	expected_names := [...]string{
		"XXXI. Donibane Ziburukonestropada- ETXEBAT sari nagusia",
		"XIII. KEPA DEUN ARRANTZALEEN KOFRADIA IKURRIÑA",
		"XXXV. ELANTXOBEKO TRAINERU ESTROPADA",
		"Mutrikuko VII. Yurrita group bandera",
		"Mundakako Onura homes XX. Ikurriña",
		"II. Bandera FEGEMU",
		"II: CMO Valves Bandera",
		"XIII. Bandera Kirol Txartela",
		"Zumaiako XXXVIII. Ikurriña",
		"ORIO KANPINA V. BANDERA",
		"XVIII. Bandera Ayto. Marina de Cudeyo . Gran premio Dynasol",
		"XLI. Bandera Noble Villa de Portugalete",
		"Plentziako XXXIV. Ikurriña",
		"XXIX Bandera Real Astillero de Guarnizo",
		"XXXI. Bandera ayto. Camargo. I. Memorial Fernando López Lejardi",
		"GOIMEK SARI NAGUSIA",
		"HONDARRIBIA ARRAUN ELKARTEA BANDERA",
		"XVII. HONDARRIBIKO ARRANTZALEEN KOFRADIKO BANDERA",
		"XXXVIII. Pasaiako Ikurriña",
		"Errenteriako Hiria X. Ikurriña",
		"LI. Bandera Ciudad de Castro Urdiales",
		"Zarauzko arraun elkartearen XIV. Ikurriña",
		"Ikurriña Bilbao",
		"XLV. Ikurriña Villa Bilbao",
		"XXIX. Bandera Ría del Asón",
		"Colindres. ARC Playoff I / XXIX. Bandera Ria del Asón",
		"Lutxana. ARC Playoff II",
	}

	estropadak, _ := Arc_parse_calendar(doc)
	if len(estropadak) != COUNT {
		t.Errorf("Expected %d estropada, got %d", COUNT, len(estropadak))
	}

	for i, name := range expected_names {
		if estropadak[i].Name != name {
			t.Errorf("Expected '%s' for estropada, got '%s'", name, estropadak[i].Name)
		}

	}

}
