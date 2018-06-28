package main

type Page struct {
	Text      string `json:"text"`
	InStates  int    `json:"intStates"`
	OutStates int    `json:"outStates"`
}

type Pages []Page
