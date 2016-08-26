package paypal

import "fmt"

type (
	CreateInvoiceResp struct {
		ID                         string             `json:"id,omitempty"`
		Number                     string             `json:"number,omitempty"`
		TemplateID                 string             `json:"template_id,omitempty"`
		URI                        string             `json:"uri,omitempty"`
		Status                     *InvoiceStatus     `json:"status,omitempty"`
		MerchantInfo               *MerchantInfo      `json:"merchant_info,omitempty"`
		BillingInfo                []*BillingInfo     `json:"billing_info,omitempty"`
		CCInfo                     []*Participant     `json:"cc_info,omitempty"`
		ShippingInfo               *ShippingInfo      `json:"shipping_info,omitempty"`
		Items                      []*InvoiceItem     `json:"items,omitempty"`
		InvoiceDate                string             `json:"invoice_date,omitempty"`
		PaymentTerm                *PaymentTerm       `json:"payment_term,omitempty"`
		Reference                  string             `json:"reference,omitempty"`
		Discount                   *Cost              `json:"discount,omitempty"`
		ShippingCost               *ShippingCost      `json:"shipping_cost,omitempty"`
		Custom                     *CustomAmount      `json:"custom,omitempty"`
		AllowPartialPayment        bool               `json:"allow_partial_payment,omitempty"`
		MinimumAmountDue           *Currency          `json:"minimum_amount_due,omitempty"`
		TaxCalculatedAfterDiscount bool               `json:"tax_calculated_after_discount,omitempty"`
		TaxInclusive               bool               `json:"tax_inclusive,omitempty"`
		Terms                      string             `json:"terms,omitempty"`
		Note                       string             `json:"note,omitempty"`
		MerchantMemo               string             `json:"merchant_memo,omitempty"`
		LogoURL                    string             `json:"logo_url,omitempty"`
		TotalAmount                *Currency          `json:"total_amount,omitempty"`
		Payments                   []*PaymentDetail   `json:"payments,omitempty"`
		Refunds                    []*RefundDetail    `json:"refunds,omitempty"`
		Metadata                   *Metadata          `json:"metadata,omitempty"`
		PaidAmount                 *PaymentSummary    `json:"paid_amount,omitempty"`
		RefundedAmount             *PaymentSummary    `json:"refunded_amount,omitempty"`
		Attachments                []*FileAttachments `json:"attachments,omitempty"`
		Links                      []*Links           `json:"links,omitempty"`
	}

	ListInvoicesRequest struct {
		Page               int  `json:"page"`
		PageSize           int  `json:"page_size"`
		TotalCountRequired bool `json:"total_count_required"`
	}

	QueryInvoicesResp struct {
		TotalCount int        `json:"total_count"`
		Invoices   []*Invoice `json:"invoices"`
		Linkds     []*Links   `json:"links"`
	}
)

// CreateInvoice creates a payment in Paypal
func (c *Client) CreateInvoice(i Invoice) (*CreateInvoiceResp, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/invoicing/invoices", c.APIBase), i)
	if err != nil {
		return nil, err
	}

	v := &CreateInvoiceResp{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// SendInvoice creates a payment in Paypal
//
// Filter Options
// 		notify_merchant - boolean
// 		Indicates whether to send the merchant an invoice update notification. Default is true.
func (c *Client) SendInvoice(invoiceID string, filter map[string]string) error {
	i := struct {
		Items []*InvoiceItem `json:"items"`
	}{}
	req, err := NewRequest("POST", fmt.Sprintf("%s/invoicing/invoices/%s/send", c.APIBase, invoiceID), i)
	if err != nil {
		return err
	}

	if filter != nil {
		q := req.URL.Query()
		for k, v := range filter {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	return c.SendWithAuth(req, nil)
}

// UpdateInvoice updates an invoice, by ID.
//
// Filter Options
// 		notify_merchant - boolean
// 		Indicates whether to send the invoice update notification to the merchant.
func (c *Client) UpdateInvoice(invoiceID string, filter map[string]string, i *Invoice) (*Invoice, error) {
	req, err := NewRequest("PUT", fmt.Sprintf("%s/invoicing/invoices/%s", c.APIBase, invoiceID), i)
	if err != nil {
		return nil, err
	}

	if filter != nil {
		q := req.URL.Query()
		for k, v := range filter {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	var invoice Invoice
	return &invoice, c.SendWithAuth(req, &invoice)
}

// GetInvoice Shows details for an invoice, by ID.
func (c *Client) GetInvoice(invoiceID string) (*Invoice, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/invoicing/invoices/%s", c.APIBase, invoiceID), nil)
	if err != nil {
		return nil, err
	}

	var invoice Invoice
	return &invoice, c.SendWithAuth(req, &invoice)
}

// ListInvoices lists invoices that belong to the merchant who makes the call.
func (c *Client) ListInvoices() ([]*Invoice, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/invoicing/invoices", c.APIBase), nil)
	if err != nil {
		return nil, err
	}

	var query QueryInvoicesResp
	return query.Invoices, c.SendWithAuth(req, &query)
}

// GenerateNextInvoiceNumber returns the number for the next invoice.
func (c *Client) GenerateNextInvoiceNumber() (string, error) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/invoicing/invoices/next-invoice-number", c.APIBase), nil)
	if err != nil {
		return "", err
	}

	result := struct {
		Number string `json:"number"`
	}{}
	return result.Number, c.SendWithAuth(req, &result)
}
