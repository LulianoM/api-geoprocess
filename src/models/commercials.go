package models

type Commercials struct {
	Metadata

	// Service   *Service
	// ServiceID uuid.UUID `json:"service_id"`

	Name         string `json:"name"`
	Class        string `json:"class"`
	Photo        string
	Coordinates  string `json:"coordinates"`
	PaymentForms string `json:"payment_forms"`

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
