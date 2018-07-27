package main

type Page struct {
	Text      string     `json:"text"`
	Instate   int        `json:"instate"`
	Character string     `json:"character"`
	Decisions []Decision `json:"decisions"`
	IsFinal   bool       `json:"isFinal"`
}

type Decision struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Pages []Page
