package paypal

import "testing"

func TestInvoice(t *testing.T) {
	withContext(func(c *Client) {
		invoice := Invoice{
			MerchantInfo: &MerchantInfo{
				Email:        "johnrowleyster-facilitator@gmail.com",
				FirstName:    "Dennis",
				LastName:     "Doctor",
				BusinessName: "Medical Professionals, LLC",
				Phone: &Phone{
					CountryCode:    "001",
					NationalNumber: "503214176",
				},
				Address: &InvoiceAddress{
					Line1:       "1234 Main St.",
					City:        "Portland",
					State:       "OR",
					PostalCode:  "97217",
					CountryCode: "US",
				},
			},
			BillingInfo: []*BillingInfo{
				&BillingInfo{
					Email: "johnrowleyster-buyer@gmail.com",
				},
			},
			Items: []*InvoiceItem{
				&InvoiceItem{
					Name:          "Sutures",
					Quantity:      100,
					UnitOfMeasure: QUANTITY,
					UnitPrice: &Currency{
						Currency: "USD",
						Value:    "5",
					},
				},
			},
			Note: "Medical Invoice 16 Jul, 2013 PST",
			PaymentTerm: &PaymentTerm{
				TermType: NET_45,
			},
			ShippingInfo: &ShippingInfo{
				FirstName:    "Sally",
				LastName:     "Patient",
				BusinessName: "Not Applicable",
				Address: &InvoiceAddress{
					Line1:       "1234 Broad St.",
					City:        "Portland",
					State:       "OR",
					PostalCode:  "97216",
					CountryCode: "US",
				},
			},
		}

		_, err := c.CreateInvoice(invoice)

		if err != nil {
			t.Fatal(err)
		}
	})
}
