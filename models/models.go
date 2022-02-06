package models

type InputText struct {
	Text string `json:"text"`
}

type OutputFrequency struct {
	Word string `json:"word"`
	Freq int    `json:"freq"`
}
