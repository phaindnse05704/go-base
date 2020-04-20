package models

type Log struct {
	Model
	IP            string `json:"date"`
	Event         string `json:"event"`
	Details       string `json:"detail"`
	ShopEmail     string `json:"shop_email"`
	ShopifyDomain string `json:"shopify_domain"`
	ShopifyData   string `json:"shopify_data"`
}
