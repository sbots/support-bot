package models

type Invoice struct {
	// Title product name
	Title string `json:"title"`
	// Description product description
	Description string `json:"description"`
	// StartParameter unique bot deep-linking parameter that can be used to generate this invoice
	StartParameter string `json:"start_parameter"`
	// Currency three-letter ISO 4217 currency code
	// (see https://core.telegram.org/bots/payments#supported-currencies)
	Currency string `json:"currency"`
	// TotalAmount total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json),
	// it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

type SuccessfulPayment struct {
	// Currency three-letter ISO 4217 currency code
	// (see https://core.telegram.org/bots/payments#supported-currencies)
	Currency string `json:"currency"`
	// TotalAmount total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json,
	// (https://core.telegram.org/bots/payments/currencies.json)
	// it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// InvoicePayload bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// ShippingOptionID identifier of the shipping option chosen by the user
	//
	// optional
	ShippingOptionID string `json:"shipping_option_id,omitempty"`
	// OrderInfo order info provided by the user
	//
	// optional
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	// TelegramPaymentChargeID telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	// ProviderPaymentChargeID provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

// OrderInfo represents information about an order.
type OrderInfo struct {
	// Name user name
	//
	// optional
	Name string `json:"name,omitempty"`
	// PhoneNumber user's phone number
	//
	// optional
	PhoneNumber string `json:"phone_number,omitempty"`
	// Email user email
	//
	// optional
	Email string `json:"email,omitempty"`
	// ShippingAddress user shipping address
	//
	// optional
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type ShippingAddress struct {
	// CountryCode ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	// State if applicable
	State string `json:"state"`
	// City city
	City string `json:"city"`
	// StreetLine1 first line for the address
	StreetLine1 string `json:"street_line1"`
	// StreetLine2 second line for the address
	StreetLine2 string `json:"street_line2"`
	// PostCode address post code
	PostCode string `json:"post_code"`
}
