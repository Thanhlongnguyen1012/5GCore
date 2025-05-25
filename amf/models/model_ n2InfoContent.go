package models

type N2infoContent struct {
	NgapIeType string    `json:"ngapIeType"`
	NgapData   *Ngapdata `json:"ngapData"`
}
