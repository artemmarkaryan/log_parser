package main

type config struct {
	Name      string     `json:"name"`
	Delimiter string     `json:"delimiter"`
	Templates []template `json:"templates"`
}

type template struct {
	Name      string  `json:"name"`
	Delimiter string  `json:"delimiter"`
	Groups    []group `json:"groups"`
}

type group struct {
	Name   string  `json:"name"`
	Fields []field `json:"fields"`
}

type field struct {
	Name    string `json:"name"`
	Pattern string `json:"pattern"`
}

type parsedOutput struct {
	Index    int
	Template string
	Fields   map[string]string
}
