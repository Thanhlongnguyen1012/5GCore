package models

type ProblemDetails struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Cause  string `json:"cause"`
}
