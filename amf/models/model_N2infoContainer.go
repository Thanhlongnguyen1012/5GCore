package models

type N2infoContainer struct {
	N2InformationClass string  `json:"n2InformationClass"`
	SmInfo             *Sminfo `json:"smInfo"`
}
