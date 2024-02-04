package estropadakParser

import (
	"os"
	"sort"
	"testing"

	"golang.org/x/net/html"
)

func TestParseARCTitle(t *testing.T) {

	var title string
	doc, err := os.Open("./html/hondarribia_arc1_2023.html")
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

func TestParseARCHeats(t *testing.T) {

	doc, err := os.Open("./html/hondarribia_arc1_2023.html")
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

// func TestParseHeatsEuskotren(t *testing.T) {

// 	doc, err := os.Open("./html/fabrika_euskotren_2023.html")
// 	if err != nil {
//             t.Error("Cannot open file", err)
// 	}
// 	tokenizer := html.NewTokenizer(doc)
//     expected := []Result{
//         {TeamName: "CR CABO DA CRUZ", HeatPosition:3, Time: "11:30,80", Ziabogak: []string{"05:55"}},
//         {TeamName: "DONOSTIA ARRAUN LAGUNAK", HeatPosition:2, Time:"10:48,98", Ziabogak: []string{"05:35"}},
//         {TeamName: "HIBAIKA JAMONES ANCIN", HeatPosition:2, Time:"11:25,96"},
//         {TeamName: "HONDARRIBIA BERTAKO IGOGAILUAK", HeatPosition:1, Time:"11:12,86"},
//         {TeamName: "NORTINDAL DONOSTIARRA UR KIROLAK", HeatPosition:1, Time:"10:44,90", Ziabogak: []string{"05:37"}},
//         {TeamName: "ORIO ORIALKI", HeatPosition:3, Time:"10:54,10"},
//         {TeamName: "SD TIRÁN PEREIRA", HeatPosition:4, Time:"11:50,24"},
//         {TeamName: "TOLOSALDEA ARRAUN KLUBA", HeatPosition:4, Time:"11:12,46"},
//     }

//     results := parse_heats(tokenizer)
//     if len(results) != 8 {
//         t.Errorf("%d != %d\n", len(results), 8)
//     }

//     sort.Sort(ByName(results))
//     for i, result := range(results) {
//         if expected[i].HeatPosition != result.HeatPosition {
//             t.Errorf("Heat position %d is not %d\n", expected[i].HeatPosition, result.HeatPosition)
//         }
//         if i < 2 {
//             if expected[i].Ziabogak[0] != result.Ziabogak[0] {
//                 t.Errorf("Ziaboga %s is not %s\n", expected[i].Ziabogak[0], result.Ziabogak[0])
//             }
//         }
//     }
// }
