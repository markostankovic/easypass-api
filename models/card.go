package models

type Card struct {
	ID		string   `json:"id,omitempty"`
	code	string   `json:"code,omitempty"`
	active	string   `json:"active,omitempty"`
	dateActivated	string   `json:"dateActivated,omitempty"`
	endDate	string   `json:"endDate,omitempty"`
	startDate	string   `json:"startDate,omitempty"`
	serialNumber	string   `json:"serialNumber,omitempty"`
}
