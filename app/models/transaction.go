package models

type Transaction struct {
	Model
	Date          string `json:"date"`
	Event         string `json:"event"`
	Details       string `json:"detail"`
	BillingOn     string `json:"billing_on"`
	ShopName      string `json:"shop_name"`
	ShopCountry   string `json:"shop_country"`
	ShopEmail     string `json:"shop_email"`
	ShopifyDomain string `json:"shopify_domain"`

	Recorded bool `gorm:"-"`
}
