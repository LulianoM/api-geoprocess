package models

type Service struct {
	// Seller   *Seller
	// SellerID uuid.UUID `json:"seller_id"`
	Metadata
	Type TypeOfService
}

type (
	TypeOfService string
)

const (
	FOOD    TypeOfService = "FOOD"
	SERVICE TypeOfService = "SERVICE"
)
