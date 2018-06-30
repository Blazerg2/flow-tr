package main

type Page struct {
	Text      string   `json:"text"`
	InState   int      `json:"intStates"`
	Character string   `json:"character"`
	Decisions []string `json:"decisions"`
}

type Pages []Page
