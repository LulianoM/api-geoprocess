package models

import "github.com/google/uuid"

type Commercials struct {
	Metadata

	Service   *Service
	ServiceID uuid.UUID `json:"service_id"`

	Name         string `json:"name"`
	Class        string
	Photo        string
	Coordinates  string
	PaymentForms string

	// ContactData   *Contact
	// OperationData *Operation
	// ProductsData  *Products
}

type Contact struct {
	PhoneNumber       string
	Location          string
	LocationReference string
	Adress            string
}

type Operation struct {
	OpeningTime              string
	ClosingTime              string
	OperationDay             string
	OperationDaysDescription string
}

type Products struct {
	Type   string
	Candys string
	Snacks string
	Drinks string
}
