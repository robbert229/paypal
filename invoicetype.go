package paypal

type NotificationChannel string

const (
	SMS   NotificationChannel = "SMS"
	EMAIL NotificationChannel = "EMAIL"
)

type Language string

const (
	DA_DK Language = "da_DK"
	DE_DE Language = "de_DE"
	EN_AU Language = "en_AU"
	EN_GB Language = "en_GB"
	EN_US Language = "en_US"
	ES_ES Language = "es_ES"
	ES_XC Language = "es_XC"
	FR_CA Language = "fr_CA"
	FR_FR Language = "fr_FR"
	FR_XC Language = "fr_XC"
	HE_IL Language = "he_IL"
	ID_ID Language = "id_ID"
	IT_IT Language = "it_IT"
	JA_JP Language = "ja_JP"
	NL_NL Language = "nl_NL"
	NO_NO Language = "no_NO"
	PL_PL Language = "pl_PL"
	PT_BR Language = "pt_BR"
	PT_PT Language = "pt_PT"
	RU_RU Language = "ru_RU"
	SV_SE Language = "sv_SE"
	TH_TH Language = "th_TH"
	ZH_CN Language = "zh_CN"
	ZH_HK Language = "zh_HK"
	ZH_TW Language = "zh_TW"
	ZH_XC Language = "zh_XC"
)

type TermType string

const (
	DUE_ON_RECEIPT        TermType = "DUE_ON_RECEIPT"
	DUE_ON_DATE_SPECIFIED TermType = "DUE_ON_DATE_SPECIFIED"
	NET_10                TermType = "NET_10"
	NET_15                TermType = "NET_15"
	NET_30                TermType = "NET_30"
	NET_45                TermType = "NET_45"
	NET_60                TermType = "NET_60"
	NET_90                TermType = "NET_90"
	NO_DUE_DATE           TermType = "NO_DUE_DATE"
)

// InvoiceStatus represents the status of an invoice.
type InvoiceStatus string

const (
	DRAFT              InvoiceStatus = "DRAFT"
	SENT               InvoiceStatus = "SENT"
	PAID               InvoiceStatus = "PAID"
	MARKED_AS_PAID     InvoiceStatus = "MARKED_AS_PAID"
	CANCELLED          InvoiceStatus = "CANCELLED"
	REFUNDED           InvoiceStatus = "REFUNDED"
	PARTIALLY_REFUNDED InvoiceStatus = "PARTIALLY_REFUNDED"
	MARKED_AS_REFUNDED InvoiceStatus = "MARKED_AS_REFUNDED"
	UNPAID             InvoiceStatus = "UNPAID"
	PAYMENT_PENDING    InvoiceStatus = "PAYMENT_PENDING"
)

type PaymentType string

const (
	PAYPAL   PaymentType = "PAYPAL"
	EXTERNAL PaymentType = "EXTERNAL"
)

type TransactionType string

const (
	SALE          TransactionType = "SALE"
	AUTHORIZATION TransactionType = "AUTHORIZATION"
	CAPTURE       TransactionType = "CAPTURE"
)

type ExternalTransactionMethod string

const (
	METHOD_BANK_TRANSFER ExternalTransactionMethod = "BANK_TRANSFER"
	METHOD_CASH          ExternalTransactionMethod = "CASH"
	METHOD_CHECK         ExternalTransactionMethod = "CHECK"
	METHOD_CREDIT_CARD   ExternalTransactionMethod = "CREDIT_CARD"
	METHOD_DEBIT_CARD    ExternalTransactionMethod = "DEBIT_CARD"
	METHOD_PAYPAL        ExternalTransactionMethod = "PAYPAL"
	METHOD_WIRE_TRANSFER ExternalTransactionMethod = "WIRE_TRANSFER"
	METHOD_OTHER         ExternalTransactionMethod = "OTHER"
)

type UnitOfMeasure string

const (
	QUANTITY UnitOfMeasure = "QUANTITY"
	HOURS    UnitOfMeasure = "HOURS"
	AMOUNT   UnitOfMeasure = "AMOUNT"
)

type (
	InvoiceAddress struct {
		Line1       string `json:"line1"`
		Line2       string `json:"line2,omitempty"`
		City        string `json:"city"`
		CountryCode string `json:"country_code"`
		PostalCode  string `json:"postal_code,omitempty"`
		State       string `json:"state,omitempty"`
		Phone       *Phone `json:"phone,omitempty"`
	}

	Phone struct {
		CountryCode    string `json:"country_code,omitempty"`
		NationalNumber string `json:"national_number,omitempty"`
	}

	MerchantInfo struct {
		Email               string          `json:"email,omitempty"`
		FirstName           string          `json:"first_name,omitempty"`
		LastName            string          `json:"last_name,omitempty"`
		BusinessName        string          `json:"business_name,omitempty"`
		Phone               *Phone          `json:"phone,omitempty"`
		Fax                 *Phone          `json:"fax,omitempty"`
		Website             string          `json:"website,omitempty"`
		Address             *InvoiceAddress `json:"address,omitempty"`
		TaxID               string          `json:"tax_id,omitempty"`
		AdditionalInfoLabel string          `json:"additional_info_label,omitempty"`
		AdditionalInfo      string          `json:"additional_info,omitempty"`
	}

	BillingInfo struct {
		Email               string               `json:"email,omitempty"`
		FirstName           string               `json:"first_name,omitempty"`
		LastName            string               `json:"last_name,omitempty"`
		BusinessName        string               `json:"business_name,omitempty"`
		Address             *InvoiceAddress      `json:"address,omitempty"`
		Language            *Language            `json:"language,omitempty"`
		AdditionalInfoLabel string               `json:"additional_info_label,omitempty"`
		AdditionalInfo      string               `json:"additional_info,omitempty"`
		NotificationChannel *NotificationChannel `json:"notification_channel,omitempty"`
		Phone               *Phone               `json:"phone,omitempty"`
	}

	Participant struct {
		Email        string          `json:"email,omitempty"`
		FirstName    string          `json:"first_name,omitempty"`
		LastName     string          `json:"last_name,omitempty"`
		BusinessName string          `json:"business_name,omitempty"`
		Phone        *Phone          `json:"phone,omitempty"`
		Fax          *Phone          `json:"fax,omitempty"`
		Website      string          `json:"website,omitempty"`
		Address      *InvoiceAddress `json:"address,omitempty"`
	}

	ShippingInfo struct {
		FirstName    string          `json:"first_name,omitempty"`
		LastName     string          `json:"last_name,omitempty"`
		BusinessName string          `json:"business_name,omitempty"`
		Address      *InvoiceAddress `json:"address,omitempty"`
	}

	PaymentTerm struct {
		DueDate  string   `json:"due_date,omitempty"` // The date when the invoice was enabled. The date format is yyyy-MM-dd z as defined in Internet Date/Time Format. (https://tools.ietf.org/html/rfc3339#section-5.6)
		TermType TermType `json:"term_type,omitempty"`
	}

	Cost struct {
		Percent float64 `json:"percent,omitempty"`
		Amount  *Amount `json:"amount,omitempty"`
	}

	Tax struct {
		ID      string  `json:"id,omitempty"`
		Name    string  `json:"name,omitempty"`
		Percent float64 `json:"percent,omitempty"`
		Amount  *Amount `json:"amount,omitempty"`
	}

	ShippingCost struct {
		Amount *Amount `json:"amount,omitempty"`
		Tax    *Tax    `json:"tax,omitempty"`
	}

	CustomAmount struct {
		Amount *Amount `json:"amount,omitempty"`
		Label  string  `json:"label,omitempty"`
	}

	Currency struct {
		Currency string `json:"currency,omitempty"`
		Value    string `json:"value,omitempty"`
	}

	FileAttachments struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	}

	InvoiceItem struct {
		Name          string        `json:"name,omitempty"`
		Description   string        `json:"description,omitempty"`
		Quantity      float64       `json:"quantity,omitempty"`
		UnitPrice     *Currency     `json:"unit_price,omitempty"`
		Tax           *Tax          `json:"tax,omitempty"`
		Date          string        `json:"date,omitempty"`
		Discount      *Cost         `json:"discount,omitempty"`
		UnitOfMeasure UnitOfMeasure `json:"unit_of_measure,omitempty"`
	}

	Invoice struct {
		Number                     int                `json:"number,omitempty"`
		TemplateID                 string             `json:"template_id,omitempty"`
		MerchantInfo               *MerchantInfo      `json:"merchant_info,omitempty"`
		BillingInfo                []*BillingInfo     `json:"billing_info,omitempty"`
		CCInfo                     []*Participant     `json:"cc_info,omitempty"`
		ShippingInfo               *ShippingInfo      `json:"shipping_info,omitempty"`
		Items                      []*InvoiceItem     `json:"items,omitempty"`
		InvoiceDate                string             `json:"invoice_date,omitempty"` // The date when the invoice was enabled. The date format is yyyy-MM-dd z as defined in Internet Date/Time Format. (https://tools.ietf.org/html/rfc3339#section-5.6)
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
		Attachments                []*FileAttachments `json:"attachments,omitempty"`
	}

	PaymentDetail struct {
		Type            *PaymentType               `json:"type,omitempty"`
		TransactionID   string                     `json:"transaction_id,omitempty"`
		TransactionType *TransactionType           `json:"transaction_type,omitempty"`
		Date            string                     `json:"date,omitempty"` // The date when the invoice was enabled. The date format is yyyy-MM-dd z as defined in Internet Date/Time Format. (https://tools.ietf.org/html/rfc3339#section-5.6)
		Method          *ExternalTransactionMethod `json:"method,omitempty"`
		Note            string                     `json:"note,omitempty"`
		Amount          *Amount                    `json:"amount,omitempty"`
	}

	RefundDetail struct {
		PaymentType *PaymentType `json:"type,omitempty"`
		Date        string       `json:"date,omitempty"` // The date when the invoice was enabled. The date format is yyyy-MM-dd z as defined in Internet Date/Time Format. (https://tools.ietf.org/html/rfc3339#section-5.6)
		Note        string       `json:"note,omitempty"`
		Amount      *Amount      `json:"amount,omitempty"`
	}

	Metadata struct {
		CreatedDate    string `json:"date,omitempty"` // The date when the invoice was enabled. The date format is yyyy-MM-dd z as defined in Internet Date/Time Format. (https://tools.ietf.org/html/rfc3339#section-5.6)
		CreatedBy      string `json:"created_by,omitempty"`
		CancelledDate  string `json:"cancelled_date,omitempty"` // The date when the invoice was enabled. The date format is yyyy-MM-dd z as defined in Internet Date/Time Format. (https://tools.ietf.org/html/rfc3339#section-5.6)
		CancelledBy    string `json:"cancelled_by,omitempty"`
		LastUpdateDate string `json:"last_update_date,omitempty"` // The date when the invoice was enabled. The date format is yyyy-MM-dd z as defined in Internet Date/Time Format. (https://tools.ietf.org/html/rfc3339#section-5.6)
		LastUpdatedBy  string `json:"last_updated_by,omitempty"`
		FirstSentDate  string `json:"first_sent_date,omitempty"`
		LastSentDate   string `json:"last_sent_date,omitempty"`
		LastSentBy     string `json:"last_sent_by,omitempty"`
		PayerViewURL   string `json:"payer_view_url,omitempty"`
	}

	PaymentSummary struct {
		Paypal *Currency `json:"paypal,omitempty"`
		Other  *Currency `json:"other,omitempty"`
	}
)
