package estropadakParser


type Result struct {
	TeamName string
	Position int
	Points int
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
