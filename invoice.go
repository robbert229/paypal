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
